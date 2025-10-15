package reporter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"
)

// ReportConfig holds configuration for report management
type ReportConfig struct {
	MaxExecutions    int    `yaml:"max_executions" json:"max_executions"`         // Maximum executions to show in main report
	ArchiveAfterDays int    `yaml:"archive_after_days" json:"archive_after_days"` // Archive reports older than N days
	CleanupAfterDays int    `yaml:"cleanup_after_days" json:"cleanup_after_days"` // Delete reports older than N days
	ReportDir        string `yaml:"report_dir" json:"report_dir"`                 // Directory for reports
	ArchiveDir       string `yaml:"archive_dir" json:"archive_dir"`               // Directory for archived reports
	EnablePagination bool   `yaml:"enable_pagination" json:"enable_pagination"`   // Enable pagination for large reports
	PageSize         int    `yaml:"page_size" json:"page_size"`                   // Executions per page
}

// DefaultReportConfig returns default configuration
func DefaultReportConfig() ReportConfig {
	return ReportConfig{
		MaxExecutions:    50,
		ArchiveAfterDays: 30,
		CleanupAfterDays: 90,
		ReportDir:        "reports",
		ArchiveDir:       "reports/archive",
		EnablePagination: true,
		PageSize:         20,
	}
}

// ReportManager manages report lifecycle and storage
type ReportManager struct {
	config ReportConfig
}

// NewReportManager creates a new report manager
func NewReportManager(config ReportConfig) *ReportManager {
	return &ReportManager{
		config: config,
	}
}

// ExecutionIndex represents metadata about stored executions
type ExecutionIndex struct {
	ExecutionID string    `json:"execution_id"`
	WorkflowID  string    `json:"workflow_id"`
	StartTime   time.Time `json:"start_time"`
	Status      string    `json:"status"`
	FilePath    string    `json:"file_path"`
}

// ReportIndex manages the index of all executions
type ReportIndex struct {
	LastUpdated time.Time        `json:"last_updated"`
	Executions  []ExecutionIndex `json:"executions"`
	Version     string           `json:"version"`
}

// LoadReportIndex loads the execution index from disk
func (rm *ReportManager) LoadReportIndex() (*ReportIndex, error) {
	indexPath := filepath.Join(rm.config.ReportDir, "index.json")

	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		// Create new index if it doesn't exist
		return &ReportIndex{
			LastUpdated: time.Now(),
			Executions:  []ExecutionIndex{},
			Version:     "1.0",
		}, nil
	}

	data, err := os.ReadFile(indexPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read index file: %w", err)
	}

	var index ReportIndex
	if err := json.Unmarshal(data, &index); err != nil {
		return nil, fmt.Errorf("failed to parse index file: %w", err)
	}

	return &index, nil
}

// SaveReportIndex saves the execution index to disk
func (rm *ReportManager) SaveReportIndex(index *ReportIndex) error {
	indexPath := filepath.Join(rm.config.ReportDir, "index.json")

	// Ensure directory exists
	if err := os.MkdirAll(rm.config.ReportDir, 0755); err != nil {
		return fmt.Errorf("failed to create report directory: %w", err)
	}

	index.LastUpdated = time.Now()

	data, err := json.MarshalIndent(index, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal index: %w", err)
	}

	if err := os.WriteFile(indexPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write index file: %w", err)
	}

	return nil
}

// GetRecentExecutions returns the most recent executions for the main report
func (rm *ReportManager) GetRecentExecutions(index *ReportIndex) []ExecutionIndex {
	// Sort executions by start time (newest first)
	sort.Slice(index.Executions, func(i, j int) bool {
		return index.Executions[i].StartTime.After(index.Executions[j].StartTime)
	})

	// Return up to MaxExecutions recent executions
	if len(index.Executions) <= rm.config.MaxExecutions {
		return index.Executions
	}

	return index.Executions[:rm.config.MaxExecutions]
}

// ArchiveOldExecutions moves old executions to archive directory
func (rm *ReportManager) ArchiveOldExecutions(index *ReportIndex) error {
	cutoffDate := time.Now().AddDate(0, 0, -rm.config.ArchiveAfterDays)

	var toArchive []ExecutionIndex
	var remaining []ExecutionIndex

	for _, exec := range index.Executions {
		if exec.StartTime.Before(cutoffDate) {
			toArchive = append(toArchive, exec)
		} else {
			remaining = append(remaining, exec)
		}
	}

	if len(toArchive) == 0 {
		return nil // Nothing to archive
	}

	// Create archive directory
	if err := os.MkdirAll(rm.config.ArchiveDir, 0755); err != nil {
		return fmt.Errorf("failed to create archive directory: %w", err)
	}

	// Group archives by month
	archivesByMonth := make(map[string][]ExecutionIndex)
	for _, exec := range toArchive {
		monthKey := exec.StartTime.Format("2006-01")
		archivesByMonth[monthKey] = append(archivesByMonth[monthKey], exec)
	}

	// Save monthly archives
	for monthKey, executions := range archivesByMonth {
		archivePath := filepath.Join(rm.config.ArchiveDir, fmt.Sprintf("%s.json", monthKey))

		archiveData := struct {
			Month      string           `json:"month"`
			ArchivedAt time.Time        `json:"archived_at"`
			Executions []ExecutionIndex `json:"executions"`
		}{
			Month:      monthKey,
			ArchivedAt: time.Now(),
			Executions: executions,
		}

		data, err := json.MarshalIndent(archiveData, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal archive data for %s: %w", monthKey, err)
		}

		if err := os.WriteFile(archivePath, data, 0644); err != nil {
			return fmt.Errorf("failed to write archive file for %s: %w", monthKey, err)
		}
	}

	// Update index to remove archived executions
	index.Executions = remaining

	return nil
}

// CleanupOldReports removes reports older than the cleanup threshold
func (rm *ReportManager) CleanupOldReports() error {
	cutoffDate := time.Now().AddDate(0, 0, -rm.config.CleanupAfterDays)

	// Clean up archive files
	archiveFiles, err := filepath.Glob(filepath.Join(rm.config.ArchiveDir, "*.json"))
	if err != nil {
		return fmt.Errorf("failed to list archive files: %w", err)
	}

	for _, archiveFile := range archiveFiles {
		info, err := os.Stat(archiveFile)
		if err != nil {
			continue // Skip files we can't stat
		}

		if info.ModTime().Before(cutoffDate) {
			if err := os.Remove(archiveFile); err != nil {
				return fmt.Errorf("failed to remove old archive file %s: %w", archiveFile, err)
			}
		}
	}

	return nil
}

// GetExecutionStats returns statistics about executions
func (rm *ReportManager) GetExecutionStats(index *ReportIndex) ExecutionStats {
	stats := ExecutionStats{
		Total:     len(index.Executions),
		Completed: 0,
		Failed:    0,
		Recent:    0,
	}

	recentCutoff := time.Now().AddDate(0, 0, -7) // Last 7 days

	for _, exec := range index.Executions {
		switch exec.Status {
		case "completed":
			stats.Completed++
		case "failed":
			stats.Failed++
		}

		if exec.StartTime.After(recentCutoff) {
			stats.Recent++
		}
	}

	if stats.Total > 0 {
		stats.SuccessRate = float64(stats.Completed) / float64(stats.Total) * 100
	}

	return stats
}

// ExecutionStats holds statistics about executions
type ExecutionStats struct {
	Total       int     `json:"total"`
	Completed   int     `json:"completed"`
	Failed      int     `json:"failed"`
	Recent      int     `json:"recent"`       // Last 7 days
	SuccessRate float64 `json:"success_rate"` // Percentage
}
