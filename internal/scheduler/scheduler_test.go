package scheduler

import (
	"testing"
	"time"

	"github.com/sintakaridina/goliteflow/internal/parser"
)

func TestScheduler_AddWorkflows(t *testing.T) {
	sched := NewScheduler()

	workflows := []parser.Workflow{
		{
			Name:     "test1",
			Schedule: "0 0 * * *",
			Tasks: []parser.Task{
				{ID: "task1", Command: "echo hello"},
			},
		},
		{
			Name:     "test2",
			Schedule: "0 1 * * *",
			Tasks: []parser.Task{
				{ID: "task2", Command: "echo world"},
			},
		},
	}

	err := sched.AddWorkflows(workflows)
	if err != nil {
		t.Fatalf("AddWorkflows() error = %v", err)
	}

	// Check that workflows were added
	addedWorkflows := sched.GetWorkflows()
	if len(addedWorkflows) != 2 {
		t.Errorf("Expected 2 workflows, got %d", len(addedWorkflows))
	}
}

func TestScheduler_AddWorkflows_InvalidCron(t *testing.T) {
	sched := NewScheduler()

	workflows := []parser.Workflow{
		{
			Name:     "test",
			Schedule: "invalid cron expression",
			Tasks: []parser.Task{
				{ID: "task1", Command: "echo hello"},
			},
		},
	}

	err := sched.AddWorkflows(workflows)
	if err == nil {
		t.Error("Expected error for invalid cron expression, got nil")
	}
}

func TestScheduler_ExecuteWorkflowNow(t *testing.T) {
	sched := NewScheduler()

	workflow := parser.Workflow{
		Name:     "test",
		Schedule: "0 0 * * *",
		Tasks: []parser.Task{
			{ID: "task1", Command: "echo hello"},
		},
	}

	err := sched.AddWorkflows([]parser.Workflow{workflow})
	if err != nil {
		t.Fatalf("AddWorkflows() error = %v", err)
	}

	execution, err := sched.ExecuteWorkflowNow("test")
	if err != nil {
		t.Fatalf("ExecuteWorkflowNow() error = %v", err)
	}

	if execution.WorkflowID != "test" {
		t.Errorf("Expected WorkflowID test, got %s", execution.WorkflowID)
	}

	if execution.Status != "completed" {
		t.Errorf("Expected status completed, got %s", execution.Status)
	}
}

func TestScheduler_ExecuteWorkflowNow_NotFound(t *testing.T) {
	sched := NewScheduler()

	_, err := sched.ExecuteWorkflowNow("nonexistent")
	if err == nil {
		t.Error("Expected error for nonexistent workflow, got nil")
	}
}

func TestScheduler_GetStats(t *testing.T) {
	sched := NewScheduler()

	workflows := []parser.Workflow{
		{
			Name:     "test1",
			Schedule: "0 0 * * *",
			Tasks: []parser.Task{
				{ID: "task1", Command: "echo hello"},
			},
		},
	}

	err := sched.AddWorkflows(workflows)
	if err != nil {
		t.Fatalf("AddWorkflows() error = %v", err)
	}

	// Execute a workflow to generate stats
	_, err = sched.ExecuteWorkflowNow("test1")
	if err != nil {
		t.Fatalf("ExecuteWorkflowNow() error = %v", err)
	}

	stats := sched.GetStats()
	if stats.TotalWorkflows != 1 {
		t.Errorf("Expected TotalWorkflows 1, got %d", stats.TotalWorkflows)
	}

	if stats.TotalExecutions != 1 {
		t.Errorf("Expected TotalExecutions 1, got %d", stats.TotalExecutions)
	}

	if stats.SuccessfulExecutions != 1 {
		t.Errorf("Expected SuccessfulExecutions 1, got %d", stats.SuccessfulExecutions)
	}
}

func TestScheduler_GetNextRunTimes(t *testing.T) {
	sched := NewScheduler()

	workflows := []parser.Workflow{
		{
			Name:     "test1",
			Schedule: "0 0 * * *", // Daily at midnight
			Tasks: []parser.Task{
				{ID: "task1", Command: "echo hello"},
			},
		},
	}

	err := sched.AddWorkflows(workflows)
	if err != nil {
		t.Fatalf("AddWorkflows() error = %v", err)
	}

	nextRuns := sched.GetNextRunTimes()
	if len(nextRuns) != 1 {
		t.Errorf("Expected 1 next run time, got %d", len(nextRuns))
	}

	nextRun, exists := nextRuns["test1"]
	if !exists {
		t.Error("Expected next run time for test1 workflow")
	}

	if nextRun.IsZero() {
		t.Error("Expected non-zero next run time")
	}

	// Next run should be in the future
	if nextRun.Before(time.Now()) {
		t.Error("Expected next run time to be in the future")
	}
}

func TestScheduler_StartStop(t *testing.T) {
	sched := NewScheduler()

	workflows := []parser.Workflow{
		{
			Name:     "test",
			Schedule: "0 0 * * *",
			Tasks: []parser.Task{
				{ID: "task1", Command: "echo hello"},
			},
		},
	}

	err := sched.AddWorkflows(workflows)
	if err != nil {
		t.Fatalf("AddWorkflows() error = %v", err)
	}

	// Start scheduler
	err = sched.Start()
	if err != nil {
		t.Fatalf("Start() error = %v", err)
	}

	// Stop scheduler
	sched.Stop()

	// Scheduler should be stopped (no way to directly test this, but it shouldn't panic)
}
