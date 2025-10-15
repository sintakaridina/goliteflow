# GoliteFlow

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.19-blue.svg)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/sintakaridina/goliteflow)](https://goreportcard.com/report/github.com/sintakaridina/goliteflow)

A lightweight workflow scheduler and task orchestrator designed for monolithic or small applications. GoliteFlow executes tasks/workflows defined in YAML files with retry logic, conditional execution, monitoring, and cron-based scheduling.

## ✨ Features

- **YAML-based Configuration**: Define workflows and tasks in simple YAML files
- **Cron Scheduling**: Built-in scheduler using cron syntax for task scheduling
- **Retry Logic**: Configurable retry mechanisms with exponential backoff
- **Task Dependencies**: Define task execution order with dependency management
- **HTML Monitoring**: Generate beautiful HTML reports with execution history
- **CLI Tool**: Command-line interface for running and managing workflows
- **Library Interface**: Use as a Go library in your applications
- **Zero External Dependencies**: No database or web server required
- **Comprehensive Logging**: Built-in logging with zerolog

## 🚀 Installation

### As a Go Module

```bash
go get github.com/sintakaridina/goliteflow
```

### Build from Source

```bash
git clone https://github.com/sintakaridina/goliteflow.git
cd goliteflow
go build -o goliteflow cmd/goliteflow/main.go
```

## 📖 Usage

### Library Usage

```go
package main

import (
    "log"
    "github.com/sintakaridina/goliteflow"
)

func main() {
    // Simple usage
    err := goliteflow.Run("workflows.yml")
    if err != nil {
        log.Fatal(err)
    }
    
    // Advanced usage with report generation
    err = goliteflow.RunWithReport("workflows.yml", "report.html")
    if err != nil {
        log.Fatal(err)
    }
    
    // Full control
    gf := goliteflow.New()
    if err := gf.LoadConfig("workflows.yml"); err != nil {
        log.Fatal(err)
    }
    
    if err := gf.Start(); err != nil {
        log.Fatal(err)
    }
    defer gf.Stop()
    
    // Run workflows once
    if err := gf.Run(); err != nil {
        log.Fatal(err)
    }
    
    // Generate report
    if err := gf.GenerateReport("report.html"); err != nil {
        log.Fatal(err)
    }
}
```

### CLI Usage

```bash
# Run workflows from configuration file
goliteflow run --config=workflows.yml

# Run as daemon (continuous execution)
goliteflow run --config=workflows.yml --daemon

# Generate HTML report
goliteflow report --output=report.html

# Validate configuration file
goliteflow validate --config=workflows.yml

# Enable verbose logging
goliteflow run --config=workflows.yml --verbose
```

## 📋 Configuration

### Workflow Configuration (YAML)

Create a `lite-workflows.yml` file:

```yaml
version: "1.0"
workflows:
  - name: daily_summary
    schedule: "0 7 * * *"  # Run daily at 7:00 AM
    tasks:
      - id: fetch_data
        command: "curl -s https://api.example.com/data"
        retry: 3
        timeout: "30s"
      - id: process_data
        depends_on: ["fetch_data"]
        command: "go run process.go"
        retry: 2
        timeout: "10s"
      - id: send_report
        depends_on: ["process_data"]
        command: "bash send_report.sh"
        retry: 1
        timeout: "5s"

  - name: hourly_cleanup
    schedule: "0 * * * *"  # Run every hour
    tasks:
      - id: cleanup_temp_files
        command: "rm -rf /tmp/old_files"
        retry: 1
        timeout: "30s"
      - id: log_cleanup_status
        depends_on: ["cleanup_temp_files"]
        command: "echo 'Cleanup completed at $(date)'"
        retry: 1
        timeout: "5s"
```

### Configuration Schema

#### Workflow Configuration
- `version`: Configuration version (required)
- `workflows`: List of workflow definitions (required)

#### Workflow Definition
- `name`: Unique workflow name (required)
- `schedule`: Cron expression for scheduling (required)
- `tasks`: List of tasks to execute (required)

#### Task Definition
- `id`: Unique task identifier (required)
- `command`: Command to execute (required)
- `retry`: Number of retry attempts (optional, default: 1)
- `depends_on`: List of task IDs this task depends on (optional)
- `timeout`: Task timeout duration (optional, default: 30m)

### Cron Schedule Format

GoliteFlow uses the standard cron format with seconds precision:

```
┌───────────── second (0 - 59)
│ ┌───────────── minute (0 - 59)
│ │ ┌───────────── hour (0 - 23)
│ │ │ ┌───────────── day of month (1 - 31)
│ │ │ │ ┌───────────── month (1 - 12)
│ │ │ │ │ ┌───────────── day of week (0 - 6) (Sunday to Saturday)
│ │ │ │ │ │
* * * * * *
```

Examples:
- `"0 0 * * *"` - Daily at midnight
- `"0 7 * * *"` - Daily at 7:00 AM
- `"0 * * * *"` - Every hour
- `"*/30 * * * *"` - Every 30 minutes
- `"0 2 * * 0"` - Every Sunday at 2:00 AM

## 📊 HTML Reports

GoliteFlow generates comprehensive HTML reports containing:

- **Workflow Statistics**: Total workflows, successful/failed runs
- **Execution History**: Detailed execution logs with timestamps
- **Task Results**: Individual task outcomes, retry counts, and logs
- **Error Details**: Captured stdout, stderr, and error messages
- **Interactive Interface**: Expandable sections for detailed viewing

### Report Features

- **Self-contained**: All CSS and JavaScript embedded
- **Responsive Design**: Works on desktop and mobile devices
- **Real-time Updates**: Reports can be regenerated with latest data
- **Export Ready**: HTML format for easy sharing and archiving

## 🧪 Testing

Run the test suite:

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/parser
go test ./internal/executor
go test ./internal/scheduler
```

## 📁 Project Structure

```
goliteflow/
├── cmd/
│   └── goliteflow/          # CLI application
│       └── main.go
├── internal/
│   ├── parser/              # YAML configuration parsing
│   │   ├── models.go
│   │   ├── yaml_parser.go
│   │   └── yaml_parser_test.go
│   ├── executor/            # Task execution engine
│   │   ├── runner.go
│   │   └── runner_test.go
│   ├── scheduler/           # Cron-based scheduling
│   │   ├── scheduler.go
│   │   └── scheduler_test.go
│   ├── reporter/            # HTML report generation
│   │   └── html_reporter.go
│   └── logger/              # Logging utilities
│       └── logger.go
├── examples/                # Example configurations
│   └── lite-workflows.yml
├── testdata/                # Test configuration files
│   ├── simple-workflow.yml
│   └── invalid-workflow.yml
├── goliteflow.go            # Main library interface
├── goliteflow_test.go       # Library tests
├── go.mod
├── go.sum
└── README.md
```

## 🔧 Development

### Prerequisites

- Go 1.19 or later
- Git

### Building

```bash
# Clone the repository
git clone https://github.com/sintakaridina/goliteflow.git
cd goliteflow

# Install dependencies
go mod tidy

# Build the CLI tool
go build -o goliteflow cmd/goliteflow/main.go

# Run tests
go test ./...

# Format code
go fmt ./...

# Lint code
go vet ./...
```

### Adding New Features

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Make your changes
4. Add tests for new functionality
5. Run tests: `go test ./...`
6. Commit changes: `git commit -m 'Add amazing feature'`
7. Push to branch: `git push origin feature/amazing-feature`
8. Open a Pull Request

## 📝 Examples

### Simple Workflow

```yaml
version: "1.0"
workflows:
  - name: hello_world
    schedule: "0 9 * * *"  # Daily at 9 AM
    tasks:
      - id: greet
        command: "echo 'Hello, World!'"
        retry: 1
```

### Complex Workflow with Dependencies

```yaml
version: "1.0"
workflows:
  - name: data_pipeline
    schedule: "0 2 * * *"  # Daily at 2 AM
    tasks:
      - id: download_data
        command: "wget https://example.com/data.csv"
        retry: 3
        timeout: "300s"
      - id: validate_data
        depends_on: ["download_data"]
        command: "python validate.py data.csv"
        retry: 2
        timeout: "60s"
      - id: process_data
        depends_on: ["validate_data"]
        command: "python process.py data.csv"
        retry: 2
        timeout: "120s"
      - id: upload_results
        depends_on: ["process_data"]
        command: "aws s3 cp results.json s3://bucket/"
        retry: 3
        timeout: "180s"
```

### Error Handling and Retries

```yaml
version: "1.0"
workflows:
  - name: robust_task
    schedule: "0 */6 * * *"  # Every 6 hours
    tasks:
      - id: api_call
        command: "curl -f https://unreliable-api.com/data"
        retry: 5  # Will retry up to 5 times
        timeout: "30s"
      - id: fallback_task
        depends_on: ["api_call"]
        command: "echo 'API call completed or failed after retries'"
        retry: 1
```

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guidelines](CONTRIBUTING.md) for details.

### Development Workflow

1. **Fork and Clone**: Fork the repository and clone your fork
2. **Create Branch**: Create a feature branch from `dev`
3. **Make Changes**: Implement your feature or fix
4. **Add Tests**: Ensure your code is well-tested
5. **Run Tests**: Verify all tests pass
6. **Submit PR**: Create a pull request to the `dev` branch

### Code Style

- Follow Go conventions and idioms
- Use `gofmt` for formatting
- Add comments for exported functions
- Write tests for new functionality
- Keep functions small and focused

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [robfig/cron](https://github.com/robfig/cron) for cron scheduling
- [rs/zerolog](https://github.com/rs/zerolog) for structured logging
- [spf13/cobra](https://github.com/spf13/cobra) for CLI framework
- [go-yaml](https://github.com/go-yaml/yaml) for YAML parsing

## 📞 Support

- **Issues**: [GitHub Issues](https://github.com/sintakaridina/goliteflow/issues)
- **Discussions**: [GitHub Discussions](https://github.com/sintakaridina/goliteflow/discussions)
- **Documentation**: [Wiki](https://github.com/sintakaridina/goliteflow/wiki)

---

**Made with ❤️ for the Go community**
