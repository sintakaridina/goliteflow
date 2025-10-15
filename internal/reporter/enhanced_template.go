package reporter

const enhancedHtmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>GoliteFlow Enhanced Report</title>
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
            background: #f8f9fa;
        }

        .container {
            max-width: 1200px;
            margin: 0 auto;
            padding: 20px;
        }

        .header {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            padding: 2rem;
            border-radius: 10px;
            margin-bottom: 2rem;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        }

        .header h1 {
            font-size: 2.5rem;
            margin-bottom: 0.5rem;
        }

        .header p {
            font-size: 1.1rem;
            opacity: 0.9;
        }

        .stats-grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
            gap: 1.5rem;
            margin-bottom: 2rem;
        }

        .stat-card {
            background: white;
            padding: 1.5rem;
            border-radius: 10px;
            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
            border-left: 4px solid #667eea;
        }

        .stat-number {
            font-size: 2rem;
            font-weight: bold;
            color: #667eea;
            margin-bottom: 0.5rem;
        }

        .stat-label {
            color: #6c757d;
            font-size: 0.9rem;
        }

        .management-info {
            background: #e3f2fd;
            border: 1px solid #2196f3;
            border-radius: 8px;
            padding: 1rem;
            margin-bottom: 2rem;
        }

        .management-info h3 {
            color: #1976d2;
            margin-bottom: 0.5rem;
        }

        .management-info ul {
            list-style: none;
            padding-left: 1rem;
        }

        .management-info li {
            margin: 0.25rem 0;
            color: #424242;
        }

        .management-info li:before {
            content: "ℹ️ ";
            margin-right: 0.5rem;
        }

        .executions-section {
            background: white;
            border-radius: 10px;
            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
            overflow: hidden;
        }

        .section-header {
            background: #f8f9fa;
            padding: 1.5rem;
            border-bottom: 1px solid #e9ecef;
        }

        .section-title {
            font-size: 1.5rem;
            color: #333;
            margin-bottom: 0.5rem;
        }

        .section-subtitle {
            color: #6c757d;
            font-size: 0.9rem;
        }

        .execution-table {
            width: 100%;
            border-collapse: collapse;
        }

        .execution-table th {
            background: #f8f9fa;
            padding: 1rem;
            text-align: left;
            font-weight: 600;
            color: #333;
            border-bottom: 2px solid #e9ecef;
        }

        .execution-table td {
            padding: 1rem;
            border-bottom: 1px solid #e9ecef;
        }

        .execution-table tr:hover {
            background: #f8f9fa;
        }

        .status-badge {
            display: inline-block;
            padding: 0.25rem 0.75rem;
            border-radius: 20px;
            font-size: 0.8rem;
            font-weight: 500;
            text-transform: uppercase;
        }

        .status-success {
            background: #d4edda;
            color: #155724;
        }

        .status-danger {
            background: #f8d7da;
            color: #721c24;
        }

        .status-warning {
            background: #fff3cd;
            color: #856404;
        }

        .status-secondary {
            background: #e2e3e5;
            color: #383d41;
        }

        .duration {
            font-family: 'Courier New', monospace;
            font-size: 0.9rem;
            color: #495057;
        }

        .timestamp {
            color: #6c757d;
            font-size: 0.9rem;
        }

        .workflow-name {
            font-weight: 600;
            color: #667eea;
        }

        .pagination {
            display: flex;
            justify-content: center;
            align-items: center;
            gap: 1rem;
            padding: 2rem;
            background: #f8f9fa;
        }

        .pagination a {
            padding: 0.5rem 1rem;
            background: #667eea;
            color: white;
            text-decoration: none;
            border-radius: 5px;
            transition: background 0.2s;
        }

        .pagination a:hover {
            background: #5a6fd8;
        }

        .pagination .disabled {
            background: #e9ecef;
            color: #6c757d;
            cursor: not-allowed;
        }

        .footer {
            text-align: center;
            margin-top: 3rem;
            padding: 2rem;
            color: #6c757d;
            font-size: 0.9rem;
        }

        .archive-notice {
            background: #fff3cd;
            border: 1px solid #ffeaa7;
            border-radius: 8px;
            padding: 1rem;
            margin-bottom: 2rem;
        }

        .archive-notice h4 {
            color: #856404;
            margin-bottom: 0.5rem;
        }

        .archive-notice p {
            color: #856404;
            margin: 0;
        }

        .no-data {
            text-align: center;
            padding: 3rem;
            color: #6c757d;
        }

        .no-data h3 {
            margin-bottom: 1rem;
        }

        @media (max-width: 768px) {
            .container {
                padding: 10px;
            }

            .header h1 {
                font-size: 2rem;
            }

            .stats-grid {
                grid-template-columns: 1fr;
            }

            .execution-table {
                font-size: 0.9rem;
            }

            .execution-table th,
            .execution-table td {
                padding: 0.5rem;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <!-- Header -->
        <div class="header">
            <h1>🚀 GoliteFlow Enhanced Report</h1>
            <p>Generated on {{formatTime .GeneratedAt}} • Intelligent Report Management</p>
        </div>

        <!-- Statistics -->
        <div class="stats-grid">
            <div class="stat-card">
                <div class="stat-number">{{.Stats.Total}}</div>
                <div class="stat-label">Total Executions</div>
            </div>
            <div class="stat-card">
                <div class="stat-number">{{.Stats.Completed}}</div>
                <div class="stat-label">Successful</div>
            </div>
            <div class="stat-card">
                <div class="stat-number">{{.Stats.Failed}}</div>
                <div class="stat-label">Failed</div>
            </div>
            <div class="stat-card">
                <div class="stat-number">{{printf "%.1f%%" .Stats.SuccessRate}}</div>
                <div class="stat-label">Success Rate</div>
            </div>
        </div>

        <!-- Report Management Info -->
        <div class="management-info">
            <h3>📋 Report Management</h3>
            <ul>
                <li>Showing latest {{.Config.MaxExecutions}} executions (older executions archived)</li>
                <li>Reports archived after {{.Config.ArchiveAfterDays}} days</li>
                <li>Archived reports cleaned up after {{.Config.CleanupAfterDays}} days</li>
                {{if .Config.EnablePagination}}
                <li>Pagination enabled with {{.Config.PageSize}} executions per page</li>
                {{end}}
            </ul>
        </div>

        {{if gt (len .RecentExecutions) 0}}
        <!-- Recent Executions -->
        <div class="executions-section">
            <div class="section-header">
                <h2 class="section-title">📊 Recent Executions</h2>
                <p class="section-subtitle">Latest {{len .RecentExecutions}} workflow executions</p>
            </div>

            <table class="execution-table">
                <thead>
                    <tr>
                        <th>Workflow</th>
                        <th>Start Time</th>
                        <th>Status</th>
                        <th>Execution ID</th>
                    </tr>
                </thead>
                <tbody>
                    {{range .RecentExecutions}}
                    <tr>
                        <td>
                            <span class="workflow-name">{{.WorkflowID}}</span>
                        </td>
                        <td>
                            <span class="timestamp">{{formatTime .StartTime}}</span>
                        </td>
                        <td>
                            <span class="status-badge status-{{statusColor .Status}}">{{.Status}}</span>
                        </td>
                        <td>
                            <code>{{truncateString .ExecutionID 12}}</code>
                        </td>
                    </tr>
                    {{end}}
                </tbody>
            </table>

            {{if .Config.EnablePagination}}
            <div class="pagination">
                <a href="#" class="disabled">← Previous</a>
                <span>Page 1 of 1</span>
                <a href="#" class="disabled">Next →</a>
            </div>
            {{end}}
        </div>

        <!-- Workflow Summary -->
        {{if .WorkflowSummary}}
        <div class="executions-section" style="margin-top: 2rem;">
            <div class="section-header">
                <h2 class="section-title">📈 Workflow Summary</h2>
                <p class="section-subtitle">Performance overview by workflow</p>
            </div>

            <table class="execution-table">
                <thead>
                    <tr>
                        <th>Workflow Name</th>
                        <th>Total Runs</th>
                        <th>Success Rate</th>
                        <th>Last Run</th>
                    </tr>
                </thead>
                <tbody>
                    {{range $name, $summary := .WorkflowSummary}}
                    <tr>
                        <td>
                            <span class="workflow-name">{{$summary.Name}}</span>
                        </td>
                        <td>{{$summary.TotalRuns}}</td>
                        <td>
                            <span class="status-badge {{if ge $summary.SuccessRate 90.0}}status-success{{else if ge $summary.SuccessRate 70.0}}status-warning{{else}}status-danger{{end}}">
                                {{printf "%.1f%%" $summary.SuccessRate}}
                            </span>
                        </td>
                        <td>
                            <span class="timestamp">{{formatTime $summary.LastRun}}</span>
                        </td>
                    </tr>
                    {{end}}
                </tbody>
            </table>
        </div>
        {{end}}

        {{else}}
        <!-- No Data State -->
        <div class="no-data">
            <h3>📋 No Recent Executions</h3>
            <p>No workflow executions found in the recent timeframe.</p>
            <p>Run some workflows to see execution data here.</p>
        </div>
        {{end}}

        <!-- Archive Notice -->
        <div class="archive-notice">
            <h4>📦 Archived Data</h4>
            <p>
                Older execution data is automatically archived to preserve performance. 
                Use <code>goliteflow report-manage stats</code> to view archive statistics.
            </p>
        </div>

        <!-- Footer -->
        <div class="footer">
            <p>
                Generated by <strong>GoliteFlow Enhanced Reporter</strong> • 
                Intelligent report management keeps your reports fast and organized
            </p>
        </div>
    </div>
</body>
</html>
`
