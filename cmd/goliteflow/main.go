package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/sintakaridina/goliteflow/internal/logger"
	"github.com/sintakaridina/goliteflow/internal/parser"
	"github.com/sintakaridina/goliteflow/internal/reporter"
	"github.com/sintakaridina/goliteflow/internal/scheduler"
	"github.com/spf13/cobra"
)

var (
	configFile string
	outputFile string
	verbose    bool
	daemon     bool
	version    bool
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "goliteflow",
	Short: "A lightweight workflow scheduler and task orchestrator",
	Long: `GoliteFlow is a lightweight workflow scheduler and task orchestrator 
designed for monolithic or small applications. It executes tasks/workflows 
defined in YAML files with retry logic, conditional execution, and monitoring.`,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			PrintVersion()
			return
		}
		cmd.Help()
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run workflows from configuration file",
	Long:  `Execute workflows defined in the YAML configuration file.`,
	RunE:  runWorkflows,
}

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate HTML report from execution data",
	Long:  `Generate an HTML report containing workflow execution history and statistics.`,
	RunE:  generateReport,
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate workflow configuration file",
	Long:  `Validate the syntax and structure of a workflow configuration file.`,
	RunE:  validateConfig,
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "lite-workflows.yml", "Configuration file path")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose logging")
	rootCmd.Flags().BoolVar(&version, "version", false, "Show version information")

	// Run command flags
	runCmd.Flags().BoolVarP(&daemon, "daemon", "d", false, "Run as daemon (continuous execution)")
	runCmd.Flags().StringVarP(&outputFile, "output", "o", "report.html", "Output file for HTML report")

	// Report command flags
	reportCmd.Flags().StringVarP(&outputFile, "output", "o", "report.html", "Output file for HTML report")

	// Add commands
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(reportCmd)
	rootCmd.AddCommand(validateCmd)
}

func runWorkflows(cmd *cobra.Command, args []string) error {
	// Initialize logger
	if verbose {
		logger.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		logger.SetGlobalLevel(zerolog.InfoLevel)
	}

	log := logger.GetGlobalLogger()
	log.Info("Starting GoliteFlow")

	// Parse configuration
	yamlParser := parser.NewYAMLParser()
	config, err := yamlParser.ParseFile(configFile)
	if err != nil {
		return fmt.Errorf("failed to parse configuration: %w", err)
	}

	log.Infof("Loaded %d workflows from %s", len(config.Workflows), configFile)

	// Create scheduler
	sched := scheduler.NewScheduler()

	// Add workflows to scheduler
	if err := sched.AddWorkflows(config.Workflows); err != nil {
		return fmt.Errorf("failed to add workflows to scheduler: %w", err)
	}

	// Start scheduler
	if err := sched.Start(); err != nil {
		return fmt.Errorf("failed to start scheduler: %w", err)
	}
	defer sched.Stop()

	log.Info("Scheduler started successfully")

	// Set up signal handling
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start report generation goroutine
	go func() {
		ticker := time.NewTicker(5 * time.Minute) // Generate report every 5 minutes
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if err := generateHTMLReport(sched, outputFile); err != nil {
					log.Errorf("Failed to generate report: %v", err)
				}
			case execution := <-sched.GetReportChannel():
				log.Infof("Workflow '%s' completed with status: %s", execution.WorkflowID, execution.Status)
			}
		}
	}()

	if daemon {
		log.Info("Running in daemon mode. Press Ctrl+C to stop.")

		// Wait for signal
		<-sigChan
		log.Info("Received shutdown signal, stopping...")
	} else {
		// Run once - execute all workflows immediately
		log.Info("Running workflows once...")

		for _, workflow := range config.Workflows {
			log.Infof("Executing workflow: %s", workflow.Name)
			execution, err := sched.ExecuteWorkflowNow(workflow.Name)
			if err != nil {
				log.Errorf("Failed to execute workflow '%s': %v", workflow.Name, err)
				continue
			}

			log.Infof("Workflow '%s' completed with status: %s", workflow.Name, execution.Status)
		}

		// Generate final report
		if err := generateHTMLReport(sched, outputFile); err != nil {
			log.Errorf("Failed to generate final report: %v", err)
		} else {
			log.Infof("Report generated: %s", outputFile)
		}
	}

	return nil
}

func generateReport(cmd *cobra.Command, args []string) error {
	// Initialize logger
	if verbose {
		logger.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		logger.SetGlobalLevel(zerolog.InfoLevel)
	}

	log := logger.GetGlobalLogger()
	log.Info("Generating HTML report")

	// For now, we'll create a mock scheduler with empty data
	// In a real implementation, this would load execution data from storage
	sched := scheduler.NewScheduler()

	if err := generateHTMLReport(sched, outputFile); err != nil {
		return fmt.Errorf("failed to generate report: %w", err)
	}

	log.Infof("Report generated successfully: %s", outputFile)
	return nil
}

func validateConfig(cmd *cobra.Command, args []string) error {
	// Initialize logger
	if verbose {
		logger.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		logger.SetGlobalLevel(zerolog.InfoLevel)
	}

	log := logger.GetGlobalLogger()
	log.Infof("Validating configuration file: %s", configFile)

	// Parse and validate configuration
	yamlParser := parser.NewYAMLParser()
	config, err := yamlParser.ParseFile(configFile)
	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	log.Infof("Configuration is valid!")
	log.Infof("Found %d workflows:", len(config.Workflows))

	for _, workflow := range config.Workflows {
		log.Infof("  - %s (schedule: %s, tasks: %d)", workflow.Name, workflow.Schedule, len(workflow.Tasks))
	}

	return nil
}

func generateHTMLReport(sched *scheduler.Scheduler, outputFile string) error {
	htmlReporter, err := reporter.NewHTMLReporter()
	if err != nil {
		return fmt.Errorf("failed to create HTML reporter: %w", err)
	}

	executions := sched.GetAllExecutions()
	if err := htmlReporter.GenerateReport(executions, outputFile); err != nil {
		return fmt.Errorf("failed to generate HTML report: %w", err)
	}

	return nil
}
