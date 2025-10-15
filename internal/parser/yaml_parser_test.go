package parser

import (
	"testing"
)

func TestYAMLParser_ParseFile(t *testing.T) {
	parser := NewYAMLParser()

	tests := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		{
			name:     "valid simple workflow",
			filename: "../../testdata/simple-workflow.yml",
			wantErr:  false,
		},
		{
			name:     "invalid workflow",
			filename: "../../testdata/invalid-workflow.yml",
			wantErr:  true,
		},
		{
			name:     "non-existent file",
			filename: "../../testdata/non-existent.yml",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config, err := parser.ParseFile(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("YAMLParser.ParseFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && config == nil {
				t.Error("Expected config to be non-nil for valid file")
			}
		})
	}
}

func TestYAMLParser_ValidateConfig(t *testing.T) {
	parser := NewYAMLParser()

	tests := []struct {
		name    string
		config  *WorkflowConfig
		wantErr bool
	}{
		{
			name: "valid config",
			config: &WorkflowConfig{
				Version: "1.0",
				Workflows: []Workflow{
					{
						Name:     "test",
						Schedule: "0 0 * * *",
						Tasks: []Task{
							{ID: "task1", Command: "echo hello"},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "missing version",
			config: &WorkflowConfig{
				Workflows: []Workflow{
					{
						Name:     "test",
						Schedule: "0 0 * * *",
						Tasks: []Task{
							{ID: "task1", Command: "echo hello"},
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "empty workflows",
			config: &WorkflowConfig{
				Version:   "1.0",
				Workflows: []Workflow{},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := parser.ValidateConfig(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("YAMLParser.ValidateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestYAMLParser_TopologicalSort(t *testing.T) {
	parser := NewYAMLParser()

	workflow := &Workflow{
		Name:     "test",
		Schedule: "0 0 * * *",
		Tasks: []Task{
			{ID: "task1", Command: "echo task1"},
			{ID: "task2", Command: "echo task2", DependsOn: []string{"task1"}},
			{ID: "task3", Command: "echo task3", DependsOn: []string{"task2"}},
		},
	}

	sorted, err := parser.TopologicalSort(workflow)
	if err != nil {
		t.Fatalf("TopologicalSort() error = %v", err)
	}

	// Check that dependencies are respected
	expectedOrder := []string{"task1", "task2", "task3"}
	if len(sorted) != len(expectedOrder) {
		t.Fatalf("Expected %d tasks, got %d", len(expectedOrder), len(sorted))
	}

	for i, task := range sorted {
		if task.ID != expectedOrder[i] {
			t.Errorf("Expected task %s at position %d, got %s", expectedOrder[i], i, task.ID)
		}
	}
}

func TestYAMLParser_TopologicalSort_CircularDependency(t *testing.T) {
	parser := NewYAMLParser()

	workflow := &Workflow{
		Name:     "test",
		Schedule: "0 0 * * *",
		Tasks: []Task{
			{ID: "task1", Command: "echo task1", DependsOn: []string{"task2"}},
			{ID: "task2", Command: "echo task2", DependsOn: []string{"task1"}},
		},
	}

	_, err := parser.TopologicalSort(workflow)
	if err == nil {
		t.Error("Expected error for circular dependency, got nil")
	}
}
