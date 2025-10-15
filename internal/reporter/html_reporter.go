package reporter

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"time"

	"github.com/sintakaridina/goliteflow/internal/parser"
)

// HTMLReporter generates HTML reports for workflow executions
type HTMLReporter struct {
	template *template.Template
}

// NewHTMLReporter creates a new HTML reporter
func NewHTMLReporter() (*HTMLReporter, error) {
	tmpl, err := template.New("report").Parse(htmlTemplate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML template: %w", err)
	}

	return &HTMLReporter{
		template: tmpl,
	}, nil
}

// GenerateReport generates an HTML report from execution data
func (hr *HTMLReporter) GenerateReport(executions map[string][]parser.WorkflowExecution, outputPath string) error {
	report := hr.buildReport(executions)

	var buf bytes.Buffer
	if err := hr.template.Execute(&buf, report); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	// Ensure output directory exists
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Write to file
	if err := os.WriteFile(outputPath, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write report file: %w", err)
	}

	return nil
}

// buildReport builds the report data structure
func (hr *HTMLReporter) buildReport(executions map[string][]parser.WorkflowExecution) ReportData {
	report := ReportData{
		GeneratedAt:     time.Now(),
		TotalWorkflows:  len(executions),
		SuccessfulRuns:  0,
		FailedRuns:      0,
		WorkflowResults: []WorkflowReport{},
	}

	for workflowName, workflowExecutions := range executions {
		workflowReport := WorkflowReport{
			Name:        workflowName,
			TotalRuns:   len(workflowExecutions),
			SuccessRate: 0,
			LastRun:     time.Time{},
			Executions:  []ExecutionReport{},
		}

		successCount := 0
		for _, execution := range workflowExecutions {
			if execution.Status == "completed" {
				successCount++
			}

			if execution.StartTime.After(workflowReport.LastRun) {
				workflowReport.LastRun = execution.StartTime
			}

			execReport := ExecutionReport{
				StartTime:    execution.StartTime,
				EndTime:      execution.EndTime,
				Duration:     execution.Duration,
				Status:       execution.Status,
				ErrorMessage: execution.ErrorMessage,
				TaskResults:  []TaskReport{},
			}

			for _, taskResult := range execution.TaskResults {
				taskReport := TaskReport{
					TaskID:     taskResult.TaskID,
					StartTime:  taskResult.StartTime,
					EndTime:    taskResult.EndTime,
					Duration:   taskResult.Duration,
					ExitCode:   taskResult.ExitCode,
					Success:    taskResult.Success,
					RetryCount: taskResult.RetryCount,
					Stdout:     taskResult.Stdout,
					Stderr:     taskResult.Stderr,
					Error:      taskResult.Error,
				}
				execReport.TaskResults = append(execReport.TaskResults, taskReport)
			}

			workflowReport.Executions = append(workflowReport.Executions, execReport)
		}

		if workflowReport.TotalRuns > 0 {
			workflowReport.SuccessRate = float64(successCount) / float64(workflowReport.TotalRuns) * 100
		}

		report.WorkflowResults = append(report.WorkflowResults, workflowReport)

		if successCount == workflowReport.TotalRuns {
			report.SuccessfulRuns++
		} else {
			report.FailedRuns++
		}
	}

	return report
}

// ReportData represents the data structure for the HTML report
type ReportData struct {
	GeneratedAt     time.Time
	TotalWorkflows  int
	SuccessfulRuns  int
	FailedRuns      int
	WorkflowResults []WorkflowReport
}

// WorkflowReport represents a workflow in the report
type WorkflowReport struct {
	Name        string
	TotalRuns   int
	SuccessRate float64
	LastRun     time.Time
	Executions  []ExecutionReport
}

// ExecutionReport represents an execution in the report
type ExecutionReport struct {
	StartTime    time.Time
	EndTime      time.Time
	Duration     time.Duration
	Status       string
	ErrorMessage string
	TaskResults  []TaskReport
}

// TaskReport represents a task in the report
type TaskReport struct {
	TaskID     string
	StartTime  time.Time
	EndTime    time.Time
	Duration   time.Duration
	ExitCode   int
	Success    bool
	RetryCount int
	Stdout     string
	Stderr     string
	Error      string
}

// HTML template with embedded CSS and JavaScript
const htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>GoliteFlow Execution Report</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            line-height: 1.6;
            color: #333;
            background-color: #f5f5f5;
        }
        
        .container {
            max-width: 1200px;
            margin: 0 auto;
            padding: 20px;
        }
        
        .header {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            padding: 30px;
            border-radius: 10px;
            margin-bottom: 30px;
            text-align: center;
        }
        
        .header h1 {
            font-size: 2.5em;
            margin-bottom: 10px;
        }
        
        .header .subtitle {
            font-size: 1.2em;
            opacity: 0.9;
        }
        
        .stats {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
            gap: 20px;
            margin-bottom: 30px;
        }
        
        .stat-card {
            background: white;
            padding: 20px;
            border-radius: 10px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
            text-align: center;
        }
        
        .stat-number {
            font-size: 2em;
            font-weight: bold;
            color: #667eea;
        }
        
        .stat-label {
            color: #666;
            margin-top: 5px;
        }
        
        .workflow {
            background: white;
            border-radius: 10px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
            margin-bottom: 20px;
            overflow: hidden;
        }
        
        .workflow-header {
            background: #f8f9fa;
            padding: 20px;
            border-bottom: 1px solid #e9ecef;
            cursor: pointer;
            display: flex;
            justify-content: space-between;
            align-items: center;
        }
        
        .workflow-header:hover {
            background: #e9ecef;
        }
        
        .workflow-name {
            font-size: 1.3em;
            font-weight: bold;
            color: #333;
        }
        
        .workflow-stats {
            display: flex;
            gap: 20px;
            font-size: 0.9em;
            color: #666;
        }
        
        .workflow-content {
            display: none;
            padding: 20px;
        }
        
        .workflow-content.expanded {
            display: block;
        }
        
        .execution {
            border: 1px solid #e9ecef;
            border-radius: 8px;
            margin-bottom: 15px;
            overflow: hidden;
        }
        
        .execution-header {
            background: #f8f9fa;
            padding: 15px;
            cursor: pointer;
            display: flex;
            justify-content: space-between;
            align-items: center;
        }
        
        .execution-header:hover {
            background: #e9ecef;
        }
        
        .status {
            padding: 4px 12px;
            border-radius: 20px;
            font-size: 0.8em;
            font-weight: bold;
            text-transform: uppercase;
        }
        
        .status.completed {
            background: #d4edda;
            color: #155724;
        }
        
        .status.failed {
            background: #f8d7da;
            color: #721c24;
        }
        
        .status.running {
            background: #fff3cd;
            color: #856404;
        }
        
        .execution-content {
            display: none;
            padding: 15px;
        }
        
        .execution-content.expanded {
            display: block;
        }
        
        .task {
            border: 1px solid #e9ecef;
            border-radius: 6px;
            margin-bottom: 10px;
            overflow: hidden;
        }
        
        .task-header {
            background: #f8f9fa;
            padding: 12px;
            cursor: pointer;
            display: flex;
            justify-content: space-between;
            align-items: center;
        }
        
        .task-header:hover {
            background: #e9ecef;
        }
        
        .task-content {
            display: none;
            padding: 12px;
            background: #fafafa;
        }
        
        .task-content.expanded {
            display: block;
        }
        
        .log-section {
            margin-top: 10px;
        }
        
        .log-section h4 {
            margin-bottom: 5px;
            color: #666;
        }
        
        .log-content {
            background: #2d3748;
            color: #e2e8f0;
            padding: 10px;
            border-radius: 4px;
            font-family: 'Courier New', monospace;
            font-size: 0.9em;
            white-space: pre-wrap;
            max-height: 200px;
            overflow-y: auto;
        }
        
        .timestamp {
            color: #666;
            font-size: 0.9em;
        }
        
        .duration {
            color: #667eea;
            font-weight: bold;
        }
        
        .retry-badge {
            background: #ffc107;
            color: #212529;
            padding: 2px 6px;
            border-radius: 10px;
            font-size: 0.7em;
            font-weight: bold;
        }
        
        .toggle-icon {
            transition: transform 0.3s ease;
        }
        
        .toggle-icon.rotated {
            transform: rotate(180deg);
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>GoliteFlow Execution Report</h1>
            <div class="subtitle">Generated on {{.GeneratedAt.Format "2006-01-02 15:04:05 MST"}}</div>
        </div>
        
        <div class="stats">
            <div class="stat-card">
                <div class="stat-number">{{.TotalWorkflows}}</div>
                <div class="stat-label">Total Workflows</div>
            </div>
            <div class="stat-card">
                <div class="stat-number">{{.SuccessfulRuns}}</div>
                <div class="stat-label">Successful Runs</div>
            </div>
            <div class="stat-card">
                <div class="stat-number">{{.FailedRuns}}</div>
                <div class="stat-label">Failed Runs</div>
            </div>
        </div>
        
        {{range .WorkflowResults}}
        <div class="workflow">
            <div class="workflow-header" onclick="toggleWorkflow('{{.Name}}')">
                <div class="workflow-name">{{.Name}}</div>
                <div class="workflow-stats">
                    <span>{{.TotalRuns}} runs</span>
                    <span>{{printf "%.1f" .SuccessRate}}% success</span>
                    <span>Last: {{.LastRun.Format "2006-01-02 15:04"}}</span>
                </div>
                <span class="toggle-icon" id="icon-{{.Name}}">▼</span>
            </div>
            <div class="workflow-content" id="content-{{.Name}}">
                {{$workflowName := .Name}}
                {{range .Executions}}
                <div class="execution">
                    <div class="execution-header" onclick="toggleExecution('{{$workflowName}}-{{.StartTime.Unix}}')">
                        <div>
                            <span class="status {{.Status}}">{{.Status}}</span>
                            <span class="timestamp">{{.StartTime.Format "2006-01-02 15:04:05"}}</span>
                            <span class="duration">{{.Duration}}</span>
                        </div>
                        <span class="toggle-icon" id="exec-icon-{{$workflowName}}-{{.StartTime.Unix}}">▼</span>
                    </div>
                    <div class="execution-content" id="exec-content-{{$workflowName}}-{{.StartTime.Unix}}">
                        {{if .ErrorMessage}}
                        <div style="color: #721c24; background: #f8d7da; padding: 10px; border-radius: 4px; margin-bottom: 10px;">
                            <strong>Error:</strong> {{.ErrorMessage}}
                        </div>
                        {{end}}
                        {{$executionStartTime := .StartTime}}
                        {{range .TaskResults}}
                        <div class="task">
                            <div class="task-header" onclick="toggleTask('{{$workflowName}}-{{$executionStartTime.Unix}}-{{.TaskID}}')">
                                <div>
                                    <span class="status {{if .Success}}completed{{else}}failed{{end}}">{{.TaskID}}</span>
                                    {{if .RetryCount}}<span class="retry-badge">{{.RetryCount}} retries</span>{{end}}
                                    <span class="duration">{{.Duration}}</span>
                                </div>
                                <span class="toggle-icon" id="task-icon-{{$workflowName}}-{{$executionStartTime.Unix}}-{{.TaskID}}">▼</span>
                            </div>
                            <div class="task-content" id="task-content-{{$workflowName}}-{{$executionStartTime.Unix}}-{{.TaskID}}">
                                <div><strong>Exit Code:</strong> {{.ExitCode}}</div>
                                <div><strong>Start Time:</strong> {{.StartTime.Format "2006-01-02 15:04:05"}}</div>
                                <div><strong>End Time:</strong> {{.EndTime.Format "2006-01-02 15:04:05"}}</div>
                                {{if .Error}}
                                <div class="log-section">
                                    <h4>Error:</h4>
                                    <div class="log-content">{{.Error}}</div>
                                </div>
                                {{end}}
                                {{if .Stdout}}
                                <div class="log-section">
                                    <h4>Stdout:</h4>
                                    <div class="log-content">{{.Stdout}}</div>
                                </div>
                                {{end}}
                                {{if .Stderr}}
                                <div class="log-section">
                                    <h4>Stderr:</h4>
                                    <div class="log-content">{{.Stderr}}</div>
                                </div>
                                {{end}}
                            </div>
                        </div>
                        {{end}}
                    </div>
                </div>
                {{end}}
            </div>
        </div>
        {{end}}
    </div>
    
    <script>
        function toggleWorkflow(workflowName) {
            const content = document.getElementById('content-' + workflowName);
            const icon = document.getElementById('icon-' + workflowName);
            
            if (content.classList.contains('expanded')) {
                content.classList.remove('expanded');
                icon.classList.remove('rotated');
            } else {
                content.classList.add('expanded');
                icon.classList.add('rotated');
            }
        }
        
        function toggleExecution(executionId) {
            const content = document.getElementById('exec-content-' + executionId);
            const icon = document.getElementById('exec-icon-' + executionId);
            
            if (content.classList.contains('expanded')) {
                content.classList.remove('expanded');
                icon.classList.remove('rotated');
            } else {
                content.classList.add('expanded');
                icon.classList.add('rotated');
            }
        }
        
        function toggleTask(taskId) {
            const content = document.getElementById('task-content-' + taskId);
            const icon = document.getElementById('task-icon-' + taskId);
            
            if (content.classList.contains('expanded')) {
                content.classList.remove('expanded');
                icon.classList.remove('rotated');
            } else {
                content.classList.add('expanded');
                icon.classList.add('rotated');
            }
        }
    </script>
</body>
</html>`
