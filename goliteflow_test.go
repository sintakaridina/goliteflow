package goliteflow

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestGoliteFlow_LoadConfig(t *testing.T) {
	gf := New()

	// Test with valid config
	err := gf.LoadConfig("testdata/simple-workflow.yml")
	if err != nil {
		t.Fatalf("LoadConfig() error = %v", err)
	}

	if gf.config == nil {
		t.Error("Expected config to be loaded")
	}

	if len(gf.config.Workflows) == 0 {
		t.Error("Expected at least one workflow")
	}
}

func TestGoliteFlow_LoadConfig_InvalidFile(t *testing.T) {
	gf := New()

	// Test with invalid config
	err := gf.LoadConfig("testdata/invalid-workflow.yml")
	if err == nil {
		t.Error("Expected error for invalid config file")
	}
}

func TestGoliteFlow_Run(t *testing.T) {
	gf := New()

	err := gf.LoadConfig("testdata/simple-workflow.yml")
	if err != nil {
		t.Fatalf("LoadConfig() error = %v", err)
	}

	err = gf.Run()
	if err != nil {
		t.Fatalf("Run() error = %v", err)
	}
}

func TestGoliteFlow_RunWithContext(t *testing.T) {
	gf := New()

	err := gf.LoadConfig("testdata/simple-workflow.yml")
	if err != nil {
		t.Fatalf("LoadConfig() error = %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = gf.RunWithContext(ctx)
	if err != nil {
		t.Fatalf("RunWithContext() error = %v", err)
	}
}

func TestGoliteFlow_RunWithContext_Cancellation(t *testing.T) {
	gf := New()

	err := gf.LoadConfig("testdata/simple-workflow.yml")
	if err != nil {
		t.Fatalf("LoadConfig() error = %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	// Cancel immediately
	cancel()

	err = gf.RunWithContext(ctx)
	if err != context.Canceled {
		t.Errorf("Expected context.Canceled error, got %v", err)
	}
}

func TestGoliteFlow_StartStop(t *testing.T) {
	gf := New()

	err := gf.LoadConfig("testdata/simple-workflow.yml")
	if err != nil {
		t.Fatalf("LoadConfig() error = %v", err)
	}

	err = gf.Start()
	if err != nil {
		t.Fatalf("Start() error = %v", err)
	}

	// Let it run briefly
	time.Sleep(100 * time.Millisecond)

	gf.Stop()
}

func TestGoliteFlow_GenerateReport(t *testing.T) {
	gf := New()

	err := gf.LoadConfig("testdata/simple-workflow.yml")
	if err != nil {
		t.Fatalf("LoadConfig() error = %v", err)
	}

	err = gf.Start()
	if err != nil {
		t.Fatalf("Start() error = %v", err)
	}
	defer gf.Stop()

	// Execute a workflow to generate some data
	err = gf.Run()
	if err != nil {
		t.Fatalf("Run() error = %v", err)
	}

	// Generate report
	reportFile := "test-report.html"
	err = gf.GenerateReport(reportFile)
	if err != nil {
		t.Fatalf("GenerateReport() error = %v", err)
	}

	// Check if report file was created
	if _, err := os.Stat(reportFile); os.IsNotExist(err) {
		t.Error("Expected report file to be created")
	}

	// Clean up
	os.Remove(reportFile)
}

func TestGoliteFlow_GetStats(t *testing.T) {
	gf := New()

	err := gf.LoadConfig("testdata/simple-workflow.yml")
	if err != nil {
		t.Fatalf("LoadConfig() error = %v", err)
	}

	err = gf.Start()
	if err != nil {
		t.Fatalf("Start() error = %v", err)
	}
	defer gf.Stop()

	// Execute a workflow
	err = gf.Run()
	if err != nil {
		t.Fatalf("Run() error = %v", err)
	}

	stats := gf.GetStats()
	if stats == nil {
		t.Error("Expected stats to be non-nil")
	}

	if stats.TotalWorkflows == 0 {
		t.Error("Expected TotalWorkflows to be greater than 0")
	}
}

func TestGoliteFlow_GetNextRunTimes(t *testing.T) {
	gf := New()

	err := gf.LoadConfig("testdata/simple-workflow.yml")
	if err != nil {
		t.Fatalf("LoadConfig() error = %v", err)
	}

	err = gf.Start()
	if err != nil {
		t.Fatalf("Start() error = %v", err)
	}
	defer gf.Stop()

	nextRuns := gf.GetNextRunTimes()
	if nextRuns == nil {
		t.Error("Expected next run times to be non-nil")
	}

	if len(nextRuns) == 0 {
		t.Error("Expected at least one next run time")
	}
}

func TestRun_ConvenienceFunction(t *testing.T) {
	err := Run("testdata/simple-workflow.yml")
	if err != nil {
		t.Fatalf("Run() error = %v", err)
	}
}

func TestRunWithReport_ConvenienceFunction(t *testing.T) {
	reportFile := "test-convenience-report.html"
	err := RunWithReport("testdata/simple-workflow.yml", reportFile)
	if err != nil {
		t.Fatalf("RunWithReport() error = %v", err)
	}

	// Check if report file was created
	if _, err := os.Stat(reportFile); os.IsNotExist(err) {
		t.Error("Expected report file to be created")
	}

	// Clean up
	os.Remove(reportFile)
}

func TestValidateConfig_ConvenienceFunction(t *testing.T) {
	// Test valid config
	err := ValidateConfig("testdata/simple-workflow.yml")
	if err != nil {
		t.Fatalf("ValidateConfig() error = %v", err)
	}

	// Test invalid config
	err = ValidateConfig("testdata/invalid-workflow.yml")
	if err == nil {
		t.Error("Expected error for invalid config")
	}
}
