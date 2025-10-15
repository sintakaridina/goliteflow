# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial release of GoliteFlow
- YAML-based workflow configuration
- Cron-based task scheduling
- Task dependency management
- Retry logic with exponential backoff
- HTML report generation
- CLI tool with multiple commands
- Library interface for Go applications
- Comprehensive logging with zerolog
- Unit tests for all components
- Docker support
- GitHub Actions CI/CD pipeline

### Features
- **Workflow Configuration**: Define workflows and tasks in YAML files
- **Scheduling**: Built-in cron scheduler for task execution
- **Retry Logic**: Configurable retry mechanisms with backoff
- **Dependencies**: Task execution order with dependency management
- **Monitoring**: Beautiful HTML reports with execution history
- **CLI Tool**: Command-line interface for workflow management
- **Library**: Use as a Go library in applications
- **Logging**: Structured logging with multiple levels
- **Testing**: Comprehensive test coverage
- **Documentation**: Complete documentation and examples

### Technical Details
- Go 1.19+ support
- Zero external database dependencies
- Self-contained HTML reports
- Cross-platform support (Linux, Windows, macOS)
- Docker containerization
- GitHub Actions CI/CD
- MIT License

## [1.0.0] - 2024-10-15

### Added
- Initial release
- Core workflow execution engine
- YAML parser with validation
- Task runner with retry logic
- Cron-based scheduler
- HTML reporter with interactive interface
- CLI tool with run, report, and validate commands
- Library interface for Go applications
- Comprehensive test suite
- Documentation and examples
- Docker support
- CI/CD pipeline

### Security
- Non-root Docker user
- Input validation for YAML configuration
- Safe command execution with timeouts
- No external network dependencies for core functionality

### Performance
- Efficient task execution with goroutines
- Minimal memory footprint
- Fast YAML parsing
- Optimized HTML report generation

---

## Release Notes

### v1.0.0 - Initial Release

This is the first stable release of GoliteFlow, a lightweight workflow scheduler and task orchestrator designed for monolithic or small applications.

#### Key Features

1. **YAML Configuration**: Simple, human-readable workflow definitions
2. **Cron Scheduling**: Flexible scheduling with standard cron syntax
3. **Task Dependencies**: Define execution order with dependency management
4. **Retry Logic**: Robust error handling with configurable retries
5. **HTML Reports**: Beautiful, interactive execution reports
6. **CLI Tool**: Easy-to-use command-line interface
7. **Library Interface**: Use as a Go library in your applications
8. **Zero Dependencies**: No external database or web server required

#### Getting Started

```bash
# Install
go get github.com/sintakaridina/goliteflow

# Use as library
import "github.com/sintakaridina/goliteflow"

# Use as CLI
goliteflow run --config=workflows.yml
```

#### Example Workflow

```yaml
version: "1.0"
workflows:
  - name: daily_summary
    schedule: "0 7 * * *"
    tasks:
      - id: fetch_data
        command: "curl -s https://api.example.com/data"
        retry: 3
      - id: process_data
        depends_on: ["fetch_data"]
        command: "go run process.go"
        retry: 2
      - id: send_report
        depends_on: ["process_data"]
        command: "bash send_report.sh"
        retry: 1
```

#### What's Next

- Web-based dashboard
- REST API for remote management
- Plugin system for custom task types
- Metrics and monitoring integration
- Kubernetes operator
- More output formats (JSON, CSV)

#### Breaking Changes

None - this is the initial release.

#### Migration Guide

N/A - this is the initial release.

#### Contributors

- Initial development and design
- Comprehensive test suite
- Documentation and examples
- CI/CD pipeline setup
- Docker support

#### Acknowledgments

- [robfig/cron](https://github.com/robfig/cron) for cron scheduling
- [rs/zerolog](https://github.com/rs/zerolog) for structured logging
- [spf13/cobra](https://github.com/spf13/cobra) for CLI framework
- [go-yaml](https://github.com/go-yaml/yaml) for YAML parsing

---

For more information, see the [README.md](README.md) and [CONTRIBUTING.md](CONTRIBUTING.md) files.
