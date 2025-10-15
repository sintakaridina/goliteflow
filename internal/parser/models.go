package parser

import "time"

// WorkflowConfig represents the root configuration structure
type WorkflowConfig struct {
	Version   string     `yaml:"version"`
	Workflows []Workflow `yaml:"workflows"`
}

// Workflow represents a single workflow definition
type Workflow struct {
	Name     string `yaml:"name"`
	Schedule string `yaml:"schedule"`
	Tasks    []Task `yaml:"tasks"`
}

// Task represents a single task within a workflow
type Task struct {
	ID        string   `yaml:"id"`
	Command   string   `yaml:"command"`
	Retry     int      `yaml:"retry,omitempty"`
	DependsOn []string `yaml:"depends_on,omitempty"`
	Timeout   string   `yaml:"timeout,omitempty"`
}

// ExecutionResult represents the result of a task execution
type ExecutionResult struct {
	TaskID     string        `json:"task_id"`
	WorkflowID string        `json:"workflow_id"`
	StartTime  time.Time     `json:"start_time"`
	EndTime    time.Time     `json:"end_time"`
	Duration   time.Duration `json:"duration"`
	ExitCode   int           `json:"exit_code"`
	Success    bool          `json:"success"`
	RetryCount int           `json:"retry_count"`
	Stdout     string        `json:"stdout"`
	Stderr     string        `json:"stderr"`
	Error      string        `json:"error,omitempty"`
}

// WorkflowExecution represents the execution state of a workflow
type WorkflowExecution struct {
	WorkflowID   string            `json:"workflow_id"`
	StartTime    time.Time         `json:"start_time"`
	EndTime      time.Time         `json:"end_time"`
	Duration     time.Duration     `json:"duration"`
	Status       string            `json:"status"` // running, completed, failed
	TaskResults  []ExecutionResult `json:"task_results"`
	ErrorMessage string            `json:"error_message,omitempty"`
}

// ExecutionReport represents the complete execution report
type ExecutionReport struct {
	GeneratedAt     time.Time           `json:"generated_at"`
	TotalWorkflows  int                 `json:"total_workflows"`
	SuccessfulRuns  int                 `json:"successful_runs"`
	FailedRuns      int                 `json:"failed_runs"`
	WorkflowResults []WorkflowExecution `json:"workflow_results"`
}
