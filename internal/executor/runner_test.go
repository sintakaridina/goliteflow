package executor

import (
	"context"
	"testing"
	"time"

	"github.com/sintakaridina/goliteflow/internal/parser"
)

func TestTaskRunner_ExecuteTask(t *testing.T) {
	runner := NewTaskRunner()
	ctx := context.Background()

	tests := []struct {
		name    string
		task    parser.Task
		wantErr bool
	}{
		{
			name: "successful command",
			task: parser.Task{
				ID:      "test1",
				Command: "echo hello",
				Retry:   1,
			},
			wantErr: false,
		},
		{
			name: "failing command",
			task: parser.Task{
				ID:      "test2",
				Command: "exit 1",
				Retry:   1,
			},
			wantErr: true,
		},
		{
			name: "command with retry",
			task: parser.Task{
				ID:      "test3",
				Command: "echo retry test",
				Retry:   3,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := runner.ExecuteTask(ctx, tt.task, "test-workflow")

			if result.TaskID != tt.task.ID {
				t.Errorf("Expected TaskID %s, got %s", tt.task.ID, result.TaskID)
			}

			if result.WorkflowID != "test-workflow" {
				t.Errorf("Expected WorkflowID test-workflow, got %s", result.WorkflowID)
			}

			if result.Success == tt.wantErr {
				t.Errorf("Expected Success %v, got %v", !tt.wantErr, result.Success)
			}

			if result.StartTime.IsZero() || result.EndTime.IsZero() {
				t.Error("Expected StartTime and EndTime to be set")
			}

			if result.Duration <= 0 {
				t.Error("Expected Duration to be positive")
			}
		})
	}
}

func TestTaskRunner_ExecuteWorkflow(t *testing.T) {
	runner := NewTaskRunner()
	ctx := context.Background()

	workflow := &parser.Workflow{
		Name:     "test-workflow",
		Schedule: "0 0 * * *",
		Tasks: []parser.Task{
			{ID: "task1", Command: "echo task1"},
			{ID: "task2", Command: "echo task2", DependsOn: []string{"task1"}},
		},
	}

	execution := runner.ExecuteWorkflow(ctx, workflow)

	if execution.WorkflowID != workflow.Name {
		t.Errorf("Expected WorkflowID %s, got %s", workflow.Name, execution.WorkflowID)
	}

	if execution.Status != "completed" {
		t.Errorf("Expected status completed, got %s", execution.Status)
	}

	if len(execution.TaskResults) != 2 {
		t.Errorf("Expected 2 task results, got %d", len(execution.TaskResults))
	}

	// Check that tasks were executed in dependency order
	if execution.TaskResults[0].TaskID != "task1" {
		t.Errorf("Expected first task to be task1, got %s", execution.TaskResults[0].TaskID)
	}

	if execution.TaskResults[1].TaskID != "task2" {
		t.Errorf("Expected second task to be task2, got %s", execution.TaskResults[1].TaskID)
	}
}

func TestTaskRunner_ExecuteWorkflow_WithFailure(t *testing.T) {
	runner := NewTaskRunner()
	ctx := context.Background()

	workflow := &parser.Workflow{
		Name:     "test-workflow",
		Schedule: "0 0 * * *",
		Tasks: []parser.Task{
			{ID: "task1", Command: "exit 1"}, // This will fail
			{ID: "task2", Command: "echo task2", DependsOn: []string{"task1"}},
		},
	}

	execution := runner.ExecuteWorkflow(ctx, workflow)

	if execution.Status != "failed" {
		t.Errorf("Expected status failed, got %s", execution.Status)
	}

	if len(execution.TaskResults) != 1 {
		t.Errorf("Expected 1 task result (workflow should stop on failure), got %d", len(execution.TaskResults))
	}
}

func TestTaskRunner_CalculateBackoff(t *testing.T) {
	runner := NewTaskRunner()

	tests := []struct {
		attempt int
		min     time.Duration
		max     time.Duration
	}{
		{0, time.Second, 2 * time.Second},
		{1, 2 * time.Second, 3 * time.Second},
		{2, 4 * time.Second, 5 * time.Second},
		{10, 5 * time.Minute, 5 * time.Minute}, // Should be capped at 5 minutes
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			backoff := runner.calculateBackoff(tt.attempt)
			if backoff < tt.min || backoff > tt.max {
				t.Errorf("Attempt %d: expected backoff between %v and %v, got %v",
					tt.attempt, tt.min, tt.max, backoff)
			}
		})
	}
}
