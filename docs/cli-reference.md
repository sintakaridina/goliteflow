---
layout: default
title: CLI Reference
description: Complete command-line interface reference for GoliteFlow. All commands, options, and examples.
keywords: goliteflow cli, goliteflow commands, goliteflow options, goliteflow reference, golang cli tool
author: GoliteFlow Team
---

# CLI Reference

This document provides a complete reference for GoliteFlow's command-line interface.

## 📋 Table of Contents

- [Installation](#installation)
- [Global Options](#global-options)
- [Commands](#commands)
- [Examples](#examples)
- [Exit Codes](#exit-codes)

## 🚀 Installation

### Go Install

```bash
go install github.com/sintakaridina/goliteflow@latest
```

### Download Binary

```bash
# Linux/macOS
curl -L https://github.com/sintakaridina/goliteflow/releases/latest/download/goliteflow-linux-amd64 -o goliteflow
chmod +x goliteflow

# Windows
curl -L https://github.com/sintakaridina/goliteflow/releases/latest/download/goliteflow-windows-amd64.exe -o goliteflow.exe
```

### Docker

```bash
docker pull sintakaridina/goliteflow:latest
```

## 🌐 Global Options

These options are available for all commands:

| Option | Short | Description | Default |
|--------|-------|-------------|---------|
| `--config` | `-c` | Configuration file path | `lite-workflows.yml` |
| `--verbose` | `-v` | Enable verbose logging | `false` |
| `--help` | `-h` | Show help information | - |

### Usage

```bash
goliteflow [global-options] <command> [command-options]
```

## 📝 Commands

### `run` - Execute Workflows

Execute workflows from a configuration file.

#### Syntax

```bash
goliteflow run [options]
```

#### Options

| Option | Short | Description | Default |
|--------|-------|-------------|---------|
| `--daemon` | `-d` | Run as daemon (continuous execution) | `false` |
| `--output` | `-o` | Output file for HTML report | `report.html` |

#### Examples

```bash
# Run workflows once
goliteflow run --config=workflows.yml

# Run as daemon
goliteflow run --config=workflows.yml --daemon

# Enable verbose logging
goliteflow run --config=workflows.yml --verbose

# Specify custom output file
goliteflow run --config=workflows.yml --output=my-report.html

# Run with all options
goliteflow run --config=workflows.yml --daemon --verbose --output=daemon-report.html
```

#### Behavior

- **One-time execution**: By default, runs all workflows once and exits
- **Daemon mode**: Continuously runs workflows according to their schedules
- **Report generation**: Automatically generates HTML report after execution
- **Signal handling**: Responds to SIGINT/SIGTERM for graceful shutdown

### `validate` - Validate Configuration

Validate a workflow configuration file for syntax and structure errors.

#### Syntax

```bash
goliteflow validate [options]
```

#### Options

| Option | Short | Description | Default |
|--------|-------|-------------|---------|
| `--verbose` | `-v` | Enable verbose output | `false` |

#### Examples

```bash
# Basic validation
goliteflow validate --config=workflows.yml

# Verbose validation
goliteflow validate --config=workflows.yml --verbose

# Validate with custom config
goliteflow validate --config=my-workflows.yml
```

#### Output

**Success:**
```
Configuration is valid!
Found 3 workflows:
  - daily_backup (schedule: 0 2 * * *, tasks: 2)
  - hourly_cleanup (schedule: 0 * * * *, tasks: 1)
  - weekly_report (schedule: 0 9 * * 1, tasks: 3)
```

**Error:**
```
Error: validation failed: workflow[0].task[1]: dependency 'nonexistent_task' not found
```

### `report` - Generate HTML Report

Generate an HTML report from execution data.

#### Syntax

```bash
goliteflow report [options]
```

#### Options

| Option | Short | Description | Default |
|--------|-------|-------------|---------|
| `--output` | `-o` | Output file for HTML report | `report.html` |

#### Examples

```bash
# Generate default report
goliteflow report

# Generate custom report
goliteflow report --output=my-report.html

# Generate with verbose logging
goliteflow report --output=report.html --verbose
```

#### Output

Generates a self-contained HTML file with:
- Workflow execution statistics
- Task execution details
- Error logs and retry information
- Interactive interface for exploring data

## 🔧 Examples

### Basic Workflow Execution

```bash
# 1. Create a simple workflow
cat > simple.yml << EOF
version: "1.0"
workflows:
  - name: hello
    schedule: "0 9 * * *"
    tasks:
      - id: greet
        command: "echo 'Hello World'"
EOF

# 2. Validate the configuration
goliteflow validate --config=simple.yml

# 3. Run the workflow
goliteflow run --config=simple.yml

# 4. View the report
open report.html
```

### Daemon Mode

```bash
# Start daemon with custom config
goliteflow run --config=production-workflows.yml --daemon

# In another terminal, check the process
ps aux | grep goliteflow

# Stop the daemon
kill <pid>
```

### Docker Usage

```bash
# Run with Docker
docker run -v $(pwd):/workspace sintakaridina/goliteflow:latest run --config=/workspace/workflows.yml

# Run as daemon with Docker
docker run -d -v $(pwd):/workspace --name goliteflow sintakaridina/goliteflow:latest run --config=/workspace/workflows.yml --daemon

# Generate report with Docker
docker run -v $(pwd):/workspace sintakaridina/goliteflow:latest report --output=/workspace/report.html
```

### Advanced Configuration

```bash
# Multiple configuration files
goliteflow run --config=workflows1.yml
goliteflow run --config=workflows2.yml

# Verbose execution for debugging
goliteflow run --config=workflows.yml --verbose

# Custom output location
goliteflow run --config=workflows.yml --output=/var/log/goliteflow/report.html
```

## 📊 Exit Codes

GoliteFlow uses standard exit codes to indicate success or failure:

| Code | Description |
|------|-------------|
| `0` | Success |
| `1` | General error |
| `2` | Configuration error |
| `3` | Execution error |
| `4` | Validation error |

### Examples

```bash
# Check exit code
goliteflow run --config=workflows.yml
echo $?  # Should be 0 for success

# Use in scripts
if goliteflow validate --config=workflows.yml; then
    echo "Configuration is valid"
    goliteflow run --config=workflows.yml
else
    echo "Configuration has errors"
    exit 1
fi
```

## 🔍 Troubleshooting

### Common Issues

**1. Command not found**
```bash
# Check if goliteflow is in PATH
which goliteflow

# Add to PATH if needed
export PATH=$PATH:/path/to/goliteflow
```

**2. Permission denied**
```bash
# Make executable (Linux/macOS)
chmod +x goliteflow

# Run with sudo if needed
sudo goliteflow run --config=workflows.yml
```

**3. Configuration file not found**
```bash
# Check file exists
ls -la workflows.yml

# Use absolute path
goliteflow run --config=/full/path/to/workflows.yml
```

**4. Port already in use (daemon mode)**
```bash
# Check for existing processes
ps aux | grep goliteflow

# Kill existing process
kill <pid>
```

### Debug Mode

Enable verbose logging for detailed debugging:

```bash
# Verbose validation
goliteflow validate --config=workflows.yml --verbose

# Verbose execution
goliteflow run --config=workflows.yml --verbose

# Verbose report generation
goliteflow report --output=report.html --verbose
```

### Log Analysis

Check the generated HTML report for:
- Task execution details
- Error messages and stack traces
- Retry attempts and timing
- Workflow execution timeline

## 🎯 Best Practices

### 1. Configuration Management

```bash
# Use descriptive config names
goliteflow run --config=production-workflows.yml
goliteflow run --config=staging-workflows.yml

# Validate before running
goliteflow validate --config=workflows.yml && goliteflow run --config=workflows.yml
```

### 2. Output Management

```bash
# Use timestamped reports
goliteflow run --config=workflows.yml --output=report-$(date +%Y%m%d).html

# Organize reports by environment
mkdir -p reports/production reports/staging
goliteflow run --config=prod-workflows.yml --output=reports/production/report.html
```

### 3. Process Management

```bash
# Use systemd for daemon mode (Linux)
sudo systemctl start goliteflow

# Use supervisor for process management
supervisorctl start goliteflow

# Use Docker for isolation
docker run -d --name goliteflow --restart unless-stopped \
  -v $(pwd):/workspace \
  sintakaridina/goliteflow:latest run --config=/workspace/workflows.yml --daemon
```

### 4. Monitoring

```bash
# Check process status
ps aux | grep goliteflow

# Monitor log files
tail -f /var/log/goliteflow.log

# Check report generation
ls -la report*.html
```

## 📚 Related Documentation

- [Getting Started](/getting-started) - Quick setup guide
- [Configuration Reference](/configuration) - YAML configuration details
- [Library API](/api) - Go library documentation
- [Examples](/examples) - Real-world use cases
- [Deployment Guide](/deployment) - Production deployment

---

<div class="next-steps">
  <h3>Ready to use the CLI?</h3>
  <p>Start with the getting started guide or explore examples.</p>
  <a href="/getting-started" class="btn btn-primary">Getting Started</a>
  <a href="/examples" class="btn btn-secondary">View Examples</a>
</div>

<style>
.next-steps {
  text-align: center;
  padding: 2rem;
  background: #f8f9fa;
  border-radius: 1rem;
  margin: 2rem 0;
}

.btn {
  display: inline-block;
  padding: 0.75rem 1.5rem;
  margin: 0.5rem;
  border-radius: 0.5rem;
  text-decoration: none;
  font-weight: 600;
  transition: all 0.3s ease;
}

.btn-primary {
  background: #007bff;
  color: white;
}

.btn-primary:hover {
  background: #0056b3;
  transform: translateY(-2px);
}

.btn-secondary {
  background: transparent;
  color: #007bff;
  border: 2px solid #007bff;
}

.btn-secondary:hover {
  background: #007bff;
  color: white;
}
</style>
