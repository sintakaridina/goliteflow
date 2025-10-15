package scheduler

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/sintakaridina/goliteflow/internal/executor"
	"github.com/sintakaridina/goliteflow/internal/parser"
)

// Scheduler manages workflow execution based on cron schedules
type Scheduler struct {
	cron       *cron.Cron
	runner     *executor.TaskRunner
	workflows  []parser.Workflow
	executions map[string][]parser.WorkflowExecution
	mu         sync.RWMutex
	ctx        context.Context
	cancel     context.CancelFunc
	reportChan chan parser.WorkflowExecution
}

// NewScheduler creates a new scheduler instance
func NewScheduler() *Scheduler {
	ctx, cancel := context.WithCancel(context.Background())

	// Create cron scheduler without seconds precision (standard cron)
	c := cron.New()

	return &Scheduler{
		cron:       c,
		runner:     executor.NewTaskRunner(),
		workflows:  []parser.Workflow{},
		executions: make(map[string][]parser.WorkflowExecution),
		ctx:        ctx,
		cancel:     cancel,
		reportChan: make(chan parser.WorkflowExecution, 100),
	}
}

// AddWorkflows adds workflows to the scheduler
func (s *Scheduler) AddWorkflows(workflows []parser.Workflow) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, workflow := range workflows {
		// Validate cron expression
		if _, err := cron.ParseStandard(workflow.Schedule); err != nil {
			return fmt.Errorf("invalid cron expression for workflow '%s': %w", workflow.Name, err)
		}

		// Add to cron scheduler
		_, err := s.cron.AddFunc(workflow.Schedule, func() {
			s.executeWorkflow(workflow)
		})
		if err != nil {
			return fmt.Errorf("failed to add workflow '%s' to scheduler: %w", workflow.Name, err)
		}
	}

	s.workflows = append(s.workflows, workflows...)
	return nil
}

// Start starts the scheduler
func (s *Scheduler) Start() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.cron == nil {
		return fmt.Errorf("scheduler not initialized")
	}

	s.cron.Start()
	return nil
}

// Stop stops the scheduler
func (s *Scheduler) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.cron != nil {
		s.cron.Stop()
	}
	s.cancel()
	close(s.reportChan)
}

// executeWorkflow executes a single workflow
func (s *Scheduler) executeWorkflow(workflow parser.Workflow) {
	execution := s.runner.ExecuteWorkflow(s.ctx, &workflow)

	// Store execution result
	s.mu.Lock()
	s.executions[workflow.Name] = append(s.executions[workflow.Name], execution)
	s.mu.Unlock()

	// Send to report channel
	select {
	case s.reportChan <- execution:
	default:
		// Channel is full, skip this report
	}
}

// GetExecutions returns all executions for a workflow
func (s *Scheduler) GetExecutions(workflowName string) []parser.WorkflowExecution {
	s.mu.RLock()
	defer s.mu.RUnlock()

	executions := s.executions[workflowName]
	// Return a copy to prevent race conditions
	result := make([]parser.WorkflowExecution, len(executions))
	copy(result, executions)
	return result
}

// GetAllExecutions returns all executions across all workflows
func (s *Scheduler) GetAllExecutions() map[string][]parser.WorkflowExecution {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Return a deep copy to prevent race conditions
	result := make(map[string][]parser.WorkflowExecution)
	for workflowName, executions := range s.executions {
		result[workflowName] = make([]parser.WorkflowExecution, len(executions))
		copy(result[workflowName], executions)
	}
	return result
}

// GetReportChannel returns the channel for receiving execution reports
func (s *Scheduler) GetReportChannel() <-chan parser.WorkflowExecution {
	return s.reportChan
}

// GetWorkflows returns all configured workflows
func (s *Scheduler) GetWorkflows() []parser.Workflow {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Return a copy
	result := make([]parser.Workflow, len(s.workflows))
	copy(result, s.workflows)
	return result
}

// ExecuteWorkflowNow executes a workflow immediately (for testing or manual triggers)
func (s *Scheduler) ExecuteWorkflowNow(workflowName string) (*parser.WorkflowExecution, error) {
	s.mu.RLock()
	var targetWorkflow *parser.Workflow
	for _, workflow := range s.workflows {
		if workflow.Name == workflowName {
			targetWorkflow = &workflow
			break
		}
	}
	s.mu.RUnlock()

	if targetWorkflow == nil {
		return nil, fmt.Errorf("workflow '%s' not found", workflowName)
	}

	execution := s.runner.ExecuteWorkflow(s.ctx, targetWorkflow)

	// Store execution result
	s.mu.Lock()
	s.executions[workflowName] = append(s.executions[workflowName], execution)
	s.mu.Unlock()

	return &execution, nil
}

// GetNextRunTimes returns the next run times for all workflows
func (s *Scheduler) GetNextRunTimes() map[string]time.Time {
	s.mu.RLock()
	defer s.mu.RUnlock()

	nextRuns := make(map[string]time.Time)

	for _, workflow := range s.workflows {
		schedule, err := cron.ParseStandard(workflow.Schedule)
		if err != nil {
			continue // Skip invalid schedules
		}

		nextRuns[workflow.Name] = schedule.Next(time.Now())
	}

	return nextRuns
}

// GetStats returns scheduler statistics
func (s *Scheduler) GetStats() SchedulerStats {
	s.mu.RLock()
	defer s.mu.RUnlock()

	stats := SchedulerStats{
		TotalWorkflows:       len(s.workflows),
		TotalExecutions:      0,
		SuccessfulExecutions: 0,
		FailedExecutions:     0,
		NextRuns:             make(map[string]time.Time),
	}

	for _, workflow := range s.workflows {
		executions := s.executions[workflow.Name]
		stats.TotalExecutions += len(executions)

		for _, execution := range executions {
			if execution.Status == "completed" {
				stats.SuccessfulExecutions++
			} else {
				stats.FailedExecutions++
			}
		}

		// Calculate next run time
		if schedule, err := cron.ParseStandard(workflow.Schedule); err == nil {
			stats.NextRuns[workflow.Name] = schedule.Next(time.Now())
		}
	}

	return stats
}

// SchedulerStats represents statistics about the scheduler
type SchedulerStats struct {
	TotalWorkflows       int                  `json:"total_workflows"`
	TotalExecutions      int                  `json:"total_executions"`
	SuccessfulExecutions int                  `json:"successful_executions"`
	FailedExecutions     int                  `json:"failed_executions"`
	NextRuns             map[string]time.Time `json:"next_runs"`
}
