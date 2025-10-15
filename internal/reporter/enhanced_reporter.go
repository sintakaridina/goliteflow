package reporter

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"time"

	"github.com/sintakaridina/goliteflow/internal/parser"
)

// EnhancedHTMLReporter generates managed HTML reports with pagination and archival
type EnhancedHTMLReporter struct {
	template      *template.Template
	reportManager *ReportManager
	config        ReportConfig
}

// NewEnhancedHTMLReporter creates a new enhanced HTML reporter
func NewEnhancedHTMLReporter(config ReportConfig) (*EnhancedHTMLReporter, error) {
	tmpl, err := template.New("report").Funcs(template.FuncMap{
		"formatDuration": formatDuration,
		"formatTime":     formatTime,
		"statusColor":    statusColor,
		"truncateString": truncateString,
	}).Parse(enhancedHtmlTemplate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML template: %w", err)
	}

	return &EnhancedHTMLReporter{
		template:      tmpl,
		reportManager: NewReportManager(config),
		config:        config,
	}, nil
}

// GenerateManagedReport generates a managed HTML report with rotation and archival
func (ehr *EnhancedHTMLReporter) GenerateManagedReport(executions map[string][]parser.WorkflowExecution, outputPath string) error {
	// Load or create report index
	index, err := ehr.reportManager.LoadReportIndex()
	if err != nil {
		return fmt.Errorf("failed to load report index: %w", err)
	}

	// Add new executions to index
	if err := ehr.addExecutionsToIndex(executions, index); err != nil {
		return fmt.Errorf("failed to add executions to index: %w", err)
	}

	// Archive old executions if needed
	if err := ehr.reportManager.ArchiveOldExecutions(index); err != nil {
		return fmt.Errorf("failed to archive old executions: %w", err)
	}

	// Cleanup very old reports
	if err := ehr.reportManager.CleanupOldReports(); err != nil {
		return fmt.Errorf("failed to cleanup old reports: %w", err)
	}

	// Generate main report with recent executions only
	recentExecutions := ehr.reportManager.GetRecentExecutions(index)
	report := ehr.buildManagedReport(recentExecutions, executions)

	// Generate HTML
	var buf bytes.Buffer
	if err := ehr.template.Execute(&buf, report); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	// Ensure output directory exists
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Write main report
	if err := os.WriteFile(outputPath, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write report file: %w", err)
	}

	// Save updated index
	if err := ehr.reportManager.SaveReportIndex(index); err != nil {
		return fmt.Errorf("failed to save report index: %w", err)
	}

	return nil
}

// addExecutionsToIndex adds new workflow executions to the report index
func (ehr *EnhancedHTMLReporter) addExecutionsToIndex(executions map[string][]parser.WorkflowExecution, index *ReportIndex) error {
	for workflowName, workflowExecutions := range executions {
		for _, execution := range workflowExecutions {
			// Generate unique execution ID
			execID := generateExecutionID(workflowName, execution.StartTime)

			// Check if execution already exists
			exists := false
			for _, existing := range index.Executions {
				if existing.ExecutionID == execID {
					exists = true
					break
				}
			}

			if !exists {
				// Store execution data to file
				execFilePath := filepath.Join(ehr.config.ReportDir, "executions", fmt.Sprintf("%s.json", execID))
				if err := ehr.storeExecutionData(execution, execFilePath); err != nil {
					return fmt.Errorf("failed to store execution data: %w", err)
				}

				// Add to index
				indexEntry := ExecutionIndex{
					ExecutionID: execID,
					WorkflowID:  workflowName,
					StartTime:   execution.StartTime,
					Status:      execution.Status,
					FilePath:    execFilePath,
				}
				index.Executions = append(index.Executions, indexEntry)
			}
		}
	}

	return nil
}

// storeExecutionData stores execution data to individual file
func (ehr *EnhancedHTMLReporter) storeExecutionData(execution parser.WorkflowExecution, filePath string) error {
	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return fmt.Errorf("failed to create execution directory: %w", err)
	}

	data, err := json.MarshalIndent(execution, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal execution data: %w", err)
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write execution file: %w", err)
	}

	return nil
}

// buildManagedReport builds report data for recent executions
func (ehr *EnhancedHTMLReporter) buildManagedReport(recentExecutions []ExecutionIndex, allExecutions map[string][]parser.WorkflowExecution) ManagedReportData {
	stats := ehr.reportManager.GetExecutionStats(&ReportIndex{Executions: recentExecutions})

	report := ManagedReportData{
		GeneratedAt:      time.Now(),
		Stats:            stats,
		RecentExecutions: recentExecutions,
		Config:           ehr.config,
		WorkflowSummary:  make(map[string]WorkflowSummary),
	}

	// Build workflow summaries
	for workflowName, executions := range allExecutions {
		summary := WorkflowSummary{
			Name:         workflowName,
			TotalRuns:    len(executions),
			SuccessCount: 0,
			LastRun:      time.Time{},
		}

		for _, exec := range executions {
			if exec.Status == "completed" {
				summary.SuccessCount++
			}
			if exec.StartTime.After(summary.LastRun) {
				summary.LastRun = exec.StartTime
			}
		}

		if summary.TotalRuns > 0 {
			summary.SuccessRate = float64(summary.SuccessCount) / float64(summary.TotalRuns) * 100
		}

		report.WorkflowSummary[workflowName] = summary
	}

	return report
}

// generateExecutionID generates a unique ID for an execution
func generateExecutionID(workflowName string, startTime time.Time) string {
	data := fmt.Sprintf("%s-%s", workflowName, startTime.Format("2006-01-02T15:04:05.000Z"))
	hash := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", hash)[:16]
}

// ManagedReportData represents the data structure for managed HTML reports
type ManagedReportData struct {
	GeneratedAt      time.Time                  `json:"generated_at"`
	Stats            ExecutionStats             `json:"stats"`
	RecentExecutions []ExecutionIndex           `json:"recent_executions"`
	Config           ReportConfig               `json:"config"`
	WorkflowSummary  map[string]WorkflowSummary `json:"workflow_summary"`
}

// WorkflowSummary represents a summary of workflow executions
type WorkflowSummary struct {
	Name         string    `json:"name"`
	TotalRuns    int       `json:"total_runs"`
	SuccessCount int       `json:"success_count"`
	SuccessRate  float64   `json:"success_rate"`
	LastRun      time.Time `json:"last_run"`
}

// Template helper functions
func formatDuration(d time.Duration) string {
	if d < time.Second {
		return fmt.Sprintf("%.0fms", float64(d.Nanoseconds())/1000000)
	}
	if d < time.Minute {
		return fmt.Sprintf("%.1fs", d.Seconds())
	}
	if d < time.Hour {
		return fmt.Sprintf("%.1fm", d.Minutes())
	}
	return fmt.Sprintf("%.1fh", d.Hours())
}

func formatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func statusColor(status string) string {
	switch status {
	case "completed":
		return "success"
	case "failed":
		return "danger"
	case "running":
		return "warning"
	default:
		return "secondary"
	}
}

func truncateString(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[:length] + "..."
}
