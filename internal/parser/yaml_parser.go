package parser

import (
	"fmt"
	"io"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// YAMLParser handles parsing of workflow configuration files
type YAMLParser struct{}

// NewYAMLParser creates a new YAML parser instance
func NewYAMLParser() *YAMLParser {
	return &YAMLParser{}
}

// ParseFile parses a YAML configuration file
func (p *YAMLParser) ParseFile(filename string) (*WorkflowConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()

	return p.ParseReader(file)
}

// ParseReader parses YAML from an io.Reader
func (p *YAMLParser) ParseReader(reader io.Reader) (*WorkflowConfig, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read data: %w", err)
	}

	return p.ParseBytes(data)
}

// ParseBytes parses YAML from byte data
func (p *YAMLParser) ParseBytes(data []byte) (*WorkflowConfig, error) {
	var config WorkflowConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	if err := p.ValidateConfig(&config); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	return &config, nil
}

// ValidateConfig validates the workflow configuration
func (p *YAMLParser) ValidateConfig(config *WorkflowConfig) error {
	if config.Version == "" {
		return fmt.Errorf("version is required")
	}

	if len(config.Workflows) == 0 {
		return fmt.Errorf("at least one workflow is required")
	}

	// Validate each workflow
	for i, workflow := range config.Workflows {
		if err := p.ValidateWorkflow(&workflow, i); err != nil {
			return err
		}
	}

	return nil
}

// ValidateWorkflow validates a single workflow
func (p *YAMLParser) ValidateWorkflow(workflow *Workflow, index int) error {
	if workflow.Name == "" {
		return fmt.Errorf("workflow[%d]: name is required", index)
	}

	if workflow.Schedule == "" {
		return fmt.Errorf("workflow[%d]: schedule is required", index)
	}

	if len(workflow.Tasks) == 0 {
		return fmt.Errorf("workflow[%d]: at least one task is required", index)
	}

	// Validate tasks and their dependencies
	taskIDs := make(map[string]bool)
	for i, task := range workflow.Tasks {
		if err := p.ValidateTask(&task, i, index); err != nil {
			return err
		}
		taskIDs[task.ID] = true
	}

	// Validate dependencies
	for i, task := range workflow.Tasks {
		for _, depID := range task.DependsOn {
			if !taskIDs[depID] {
				return fmt.Errorf("workflow[%d].task[%d]: dependency '%s' not found", index, i, depID)
			}
		}
	}

	return nil
}

// ValidateTask validates a single task
func (p *YAMLParser) ValidateTask(task *Task, taskIndex, workflowIndex int) error {
	if task.ID == "" {
		return fmt.Errorf("workflow[%d].task[%d]: id is required", workflowIndex, taskIndex)
	}

	if task.Command == "" {
		return fmt.Errorf("workflow[%d].task[%d]: command is required", workflowIndex, taskIndex)
	}

	if task.Retry < 0 {
		return fmt.Errorf("workflow[%d].task[%d]: retry count cannot be negative", workflowIndex, taskIndex)
	}

	// Validate timeout format if provided
	if task.Timeout != "" {
		if _, err := time.ParseDuration(task.Timeout); err != nil {
			return fmt.Errorf("workflow[%d].task[%d]: invalid timeout format '%s': %w", workflowIndex, taskIndex, task.Timeout, err)
		}
	}

	return nil
}

// GetTaskDependencies returns a map of task dependencies for topological sorting
func (p *YAMLParser) GetTaskDependencies(workflow *Workflow) map[string][]string {
	deps := make(map[string][]string)
	for _, task := range workflow.Tasks {
		deps[task.ID] = task.DependsOn
	}
	return deps
}

// TopologicalSort sorts tasks by their dependencies
func (p *YAMLParser) TopologicalSort(workflow *Workflow) ([]Task, error) {
	deps := p.GetTaskDependencies(workflow)
	visited := make(map[string]bool)
	temp := make(map[string]bool)
	result := []Task{}

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

		// Find the task and add it to result
		for _, task := range workflow.Tasks {
			if task.ID == taskID {
				result = append(result, task)
				break
			}
		}
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
