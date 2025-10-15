package executor

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/sintakaridina/goliteflow/internal/parser"
)

// TaskRunner handles execution of individual tasks
type TaskRunner struct {
	timeout time.Duration
}

// NewTaskRunner creates a new task runner
func NewTaskRunner() *TaskRunner {
	return &TaskRunner{
		timeout: 30 * time.Minute, // default timeout
	}
}

// SetDefaultTimeout sets the default timeout for tasks
func (tr *TaskRunner) SetDefaultTimeout(timeout time.Duration) {
	tr.timeout = timeout
}

// ExecuteTask executes a single task with retry logic
func (tr *TaskRunner) ExecuteTask(ctx context.Context, task parser.Task, workflowID string) parser.ExecutionResult {
	result := parser.ExecutionResult{
		TaskID:     task.ID,
		WorkflowID: workflowID,
		StartTime:  time.Now(),
	}

	// Parse task timeout
	taskTimeout := tr.timeout
	if task.Timeout != "" {
		if parsedTimeout, err := time.ParseDuration(task.Timeout); err == nil {
			taskTimeout = parsedTimeout
		}
	}

	// Create context with timeout
	taskCtx, cancel := context.WithTimeout(ctx, taskTimeout)
	defer cancel()

	// Execute with retries
	maxRetries := task.Retry
	if maxRetries == 0 {
		maxRetries = 1 // at least one attempt
	}

	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		result.RetryCount = attempt

		// Execute the command
		cmdResult := tr.executeCommand(taskCtx, task.Command)

		// Merge results
		result.ExitCode = cmdResult.ExitCode
		result.Stdout = cmdResult.Stdout
		result.Stderr = cmdResult.Stderr
		result.Error = cmdResult.Error
		result.Success = cmdResult.ExitCode == 0

		// If successful, break out of retry loop
		if result.Success {
			break
		}

		lastErr = fmt.Errorf("command failed with exit code %d: %s", cmdResult.ExitCode, cmdResult.Error)

		// If this is not the last attempt, wait before retrying
		if attempt < maxRetries-1 {
			backoffDuration := tr.calculateBackoff(attempt)
			select {
			case <-taskCtx.Done():
				result.Error = "task cancelled or timed out"
				result.Success = false
				break
			case <-time.After(backoffDuration):
				// Continue to next attempt
			}
		}
	}

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)

	// If all retries failed, set the final error
	if !result.Success && lastErr != nil {
		result.Error = lastErr.Error()
	}

	return result
}

// CommandResult represents the result of a single command execution
type CommandResult struct {
	ExitCode int
	Stdout   string
	Stderr   string
	Error    string
}

// executeCommand executes a single command
func (tr *TaskRunner) executeCommand(ctx context.Context, command string) CommandResult {
	result := CommandResult{}

	// Parse command and arguments
	parts := strings.Fields(command)
	if len(parts) == 0 {
		result.Error = "empty command"
		result.ExitCode = 1
		return result
	}

	cmd := exec.CommandContext(ctx, parts[0], parts[1:]...)

	// Capture stdout and stderr
	stdout, err := cmd.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			result.ExitCode = exitError.ExitCode()
			result.Stderr = string(exitError.Stderr)
		} else {
			result.ExitCode = 1
			result.Error = err.Error()
		}
	} else {
		result.ExitCode = 0
		result.Stdout = string(stdout)
	}

	return result
}

// calculateBackoff calculates the backoff duration for retries
func (tr *TaskRunner) calculateBackoff(attempt int) time.Duration {
	// Exponential backoff: 1s, 2s, 4s, 8s, etc.
	baseDelay := time.Second
	backoff := time.Duration(1<<uint(attempt)) * baseDelay

	// Cap at 5 minutes
	maxBackoff := 5 * time.Minute
	if backoff > maxBackoff {
		backoff = maxBackoff
	}

	return backoff
}

// ExecuteWorkflow executes all tasks in a workflow in dependency order
func (tr *TaskRunner) ExecuteWorkflow(ctx context.Context, workflow *parser.Workflow) parser.WorkflowExecution {
	execution := parser.WorkflowExecution{
		WorkflowID:  workflow.Name,
		StartTime:   time.Now(),
		Status:      "running",
		TaskResults: []parser.ExecutionResult{},
	}

	// Sort tasks by dependencies
	sortedTasks, err := tr.sortTasksByDependencies(workflow)
	if err != nil {
		execution.Status = "failed"
		execution.ErrorMessage = fmt.Sprintf("failed to sort tasks: %v", err)
		execution.EndTime = time.Now()
		execution.Duration = execution.EndTime.Sub(execution.StartTime)
		return execution
	}

	// Track completed tasks
	completedTasks := make(map[string]bool)

	// Execute tasks in order
	for _, task := range sortedTasks {
		// Check if all dependencies are completed
		if !tr.areDependenciesCompleted(task, completedTasks) {
			execution.Status = "failed"
			execution.ErrorMessage = fmt.Sprintf("dependency check failed for task '%s'", task.ID)
			execution.EndTime = time.Now()
			execution.Duration = execution.EndTime.Sub(execution.StartTime)
			return execution
		}

		// Execute the task
		result := tr.ExecuteTask(ctx, task, workflow.Name)
		execution.TaskResults = append(execution.TaskResults, result)
		completedTasks[task.ID] = true

		// If task failed and we should stop on failure, mark workflow as failed
		if !result.Success {
			execution.Status = "failed"
			execution.ErrorMessage = fmt.Sprintf("task '%s' failed: %s", task.ID, result.Error)
			execution.EndTime = time.Now()
			execution.Duration = execution.EndTime.Sub(execution.StartTime)
			return execution
		}
	}

	// All tasks completed successfully
	execution.Status = "completed"
	execution.EndTime = time.Now()
	execution.Duration = execution.EndTime.Sub(execution.StartTime)

	return execution
}

// sortTasksByDependencies sorts tasks by their dependencies using topological sort
func (tr *TaskRunner) sortTasksByDependencies(workflow *parser.Workflow) ([]parser.Task, error) {
	// Create a map of task dependencies
	deps := make(map[string][]string)
	taskMap := make(map[string]parser.Task)

	for _, task := range workflow.Tasks {
		deps[task.ID] = task.DependsOn
		taskMap[task.ID] = task
	}

	// Topological sort
	visited := make(map[string]bool)
	temp := make(map[string]bool)
	result := []parser.Task{}

	var visit func(string) error
	visit = func(taskID string) error {
		if temp[taskID] {
			return fmt.Errorf("circular dependency detected involving task '%s'", taskID)
		}
		if visited[taskID] {
			return nil
		}

		temp[taskID] = true
		for _, dep := range deps[taskID] {
			if err := visit(dep); err != nil {
				return err
			}
		}
		temp[taskID] = false
		visited[taskID] = true

		result = append(result, taskMap[taskID])
		return nil
	}

	for _, task := range workflow.Tasks {
		if !visited[task.ID] {
			if err := visit(task.ID); err != nil {
				return nil, err
			}
		}
	}

	return result, nil
}

// areDependenciesCompleted checks if all dependencies for a task are completed
func (tr *TaskRunner) areDependenciesCompleted(task parser.Task, completedTasks map[string]bool) bool {
	for _, depID := range task.DependsOn {
		if !completedTasks[depID] {
			return false
		}
	}
	return true
}
