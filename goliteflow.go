package goliteflow

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"github.com/sintakaridina/goliteflow/internal/logger"
	"github.com/sintakaridina/goliteflow/internal/parser"
	"github.com/sintakaridina/goliteflow/internal/reporter"
	"github.com/sintakaridina/goliteflow/internal/scheduler"
)

// GoliteFlow is the main library interface
type GoliteFlow struct {
	scheduler *scheduler.Scheduler
	config    *parser.WorkflowConfig
	logger    *logger.Logger
}

// New creates a new GoliteFlow instance
func New() *GoliteFlow {
	return &GoliteFlow{
		logger: logger.NewLogger(),
	}
}

// LoadConfig loads workflow configuration from a YAML file
func (gf *GoliteFlow) LoadConfig(filename string) error {
	yamlParser := parser.NewYAMLParser()
	config, err := yamlParser.ParseFile(filename)
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	gf.config = config
	gf.logger.Infof("Loaded %d workflows from %s", len(config.Workflows), filename)
	return nil
}

// Start starts the workflow scheduler
func (gf *GoliteFlow) Start() error {
	if gf.config == nil {
		return fmt.Errorf("configuration not loaded, call LoadConfig first")
	}

	gf.scheduler = scheduler.NewScheduler()

	// Add workflows to scheduler
	if err := gf.scheduler.AddWorkflows(gf.config.Workflows); err != nil {
		return fmt.Errorf("failed to add workflows to scheduler: %w", err)
	}

	// Start scheduler
	if err := gf.scheduler.Start(); err != nil {
		return fmt.Errorf("failed to start scheduler: %w", err)
	}

	gf.logger.Info("GoliteFlow scheduler started successfully")
	return nil
}

// Stop stops the workflow scheduler
func (gf *GoliteFlow) Stop() {
	if gf.scheduler != nil {
		gf.scheduler.Stop()
		gf.logger.Info("GoliteFlow scheduler stopped")
	}
}

// Run executes workflows once (non-daemon mode)
func (gf *GoliteFlow) Run() error {
	if gf.config == nil {
		return fmt.Errorf("configuration not loaded, call LoadConfig first")
	}

	gf.logger.Info("Running workflows once...")

	for _, workflow := range gf.config.Workflows {
		gf.logger.Infof("Executing workflow: %s", workflow.Name)

		// Create a temporary scheduler for one-time execution
		tempScheduler := scheduler.NewScheduler()
		if err := tempScheduler.AddWorkflows([]parser.Workflow{workflow}); err != nil {
			gf.logger.Errorf("Failed to add workflow '%s' to scheduler: %v", workflow.Name, err)
			continue
		}

		execution, err := tempScheduler.ExecuteWorkflowNow(workflow.Name)
		if err != nil {
			gf.logger.Errorf("Failed to execute workflow '%s': %v", workflow.Name, err)
			continue
		}

		gf.logger.Infof("Workflow '%s' completed with status: %s", workflow.Name, execution.Status)
	}

	return nil
}

// RunWithContext executes workflows with a context for cancellation
func (gf *GoliteFlow) RunWithContext(ctx context.Context) error {
	if gf.config == nil {
		return fmt.Errorf("configuration not loaded, call LoadConfig first")
	}

	gf.logger.Info("Running workflows with context...")

	for _, workflow := range gf.config.Workflows {
		select {
		case <-ctx.Done():
			gf.logger.Info("Context cancelled, stopping workflow execution")
			return ctx.Err()
		default:
		}

		gf.logger.Infof("Executing workflow: %s", workflow.Name)

		// Create a temporary scheduler for one-time execution
		tempScheduler := scheduler.NewScheduler()
		if err := tempScheduler.AddWorkflows([]parser.Workflow{workflow}); err != nil {
			gf.logger.Errorf("Failed to add workflow '%s' to scheduler: %v", workflow.Name, err)
			continue
		}

		execution, err := tempScheduler.ExecuteWorkflowNow(workflow.Name)
		if err != nil {
			gf.logger.Errorf("Failed to execute workflow '%s': %v", workflow.Name, err)
			continue
		}

		gf.logger.Infof("Workflow '%s' completed with status: %s", workflow.Name, execution.Status)
	}

	return nil
}

// GenerateReport generates an HTML report of execution history
func (gf *GoliteFlow) GenerateReport(outputFile string) error {
	if gf.scheduler == nil {
		return fmt.Errorf("scheduler not started, call Start first")
	}

	htmlReporter, err := reporter.NewHTMLReporter()
	if err != nil {
		return fmt.Errorf("failed to create HTML reporter: %w", err)
	}

	executions := gf.scheduler.GetAllExecutions()
	if err := htmlReporter.GenerateReport(executions, outputFile); err != nil {
		return fmt.Errorf("failed to generate HTML report: %w", err)
	}

	gf.logger.Infof("Report generated: %s", outputFile)
	return nil
}

// GetStats returns scheduler statistics
func (gf *GoliteFlow) GetStats() *scheduler.SchedulerStats {
	if gf.scheduler == nil {
		return nil
	}

	stats := gf.scheduler.GetStats()
	return &stats
}

// GetExecutions returns execution history for a specific workflow
func (gf *GoliteFlow) GetExecutions(workflowName string) []parser.WorkflowExecution {
	if gf.scheduler == nil {
		return nil
	}

	return gf.scheduler.GetExecutions(workflowName)
}

// GetNextRunTimes returns the next scheduled run times for all workflows
func (gf *GoliteFlow) GetNextRunTimes() map[string]time.Time {
	if gf.scheduler == nil {
		return nil
	}

	return gf.scheduler.GetNextRunTimes()
}

// SetLogLevel sets the logging level
func (gf *GoliteFlow) SetLogLevel(level zerolog.Level) {
	gf.logger.SetLevel(level)
}

// GetLogger returns the logger instance
func (gf *GoliteFlow) GetLogger() *logger.Logger {
	return gf.logger
}

// Convenience functions for simple usage

// Run is a convenience function that loads config and runs workflows once
func Run(configFile string) error {
	gf := New()
	if err := gf.LoadConfig(configFile); err != nil {
		return err
	}
	return gf.Run()
}

// RunWithReport is a convenience function that runs workflows and generates a report
func RunWithReport(configFile, reportFile string) error {
	gf := New()
	if err := gf.LoadConfig(configFile); err != nil {
		return err
	}

	if err := gf.Run(); err != nil {
		return err
	}

	// Create a temporary scheduler to capture execution data
	tempScheduler := scheduler.NewScheduler()
	if err := tempScheduler.AddWorkflows(gf.config.Workflows); err != nil {
		return fmt.Errorf("failed to add workflows to scheduler: %w", err)
	}

	// Execute workflows and capture results
	for _, workflow := range gf.config.Workflows {
		_, err := tempScheduler.ExecuteWorkflowNow(workflow.Name)
		if err != nil {
			gf.logger.Errorf("Failed to execute workflow '%s': %v", workflow.Name, err)
		}
	}

	// Generate report
	htmlReporter, err := reporter.NewHTMLReporter()
	if err != nil {
		return fmt.Errorf("failed to create HTML reporter: %w", err)
	}

	executions := tempScheduler.GetAllExecutions()
	if err := htmlReporter.GenerateReport(executions, reportFile); err != nil {
		return fmt.Errorf("failed to generate HTML report: %w", err)
	}

	return nil
}

// ValidateConfig validates a workflow configuration file
func ValidateConfig(configFile string) error {
	yamlParser := parser.NewYAMLParser()
	_, err := yamlParser.ParseFile(configFile)
	return err
}
