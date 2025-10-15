---
layout: default
title: Getting Started
description: Quick start guide for GoliteFlow
---

# Getting Started

This guide will help you get up and running with GoliteFlow in just a few minutes.

## Prerequisites

- **Go 1.19+** (for library usage)
- **Basic knowledge** of YAML and command-line tools
- **Git** (for installation)

## Installation

### Option 1: Go Install (Recommended)

```bash
go install github.com/sintakaridina/goliteflow@latest
```

### Option 2: Download Binary

Download the latest release for your platform:

```bash
# Linux/macOS
curl -L https://github.com/sintakaridina/goliteflow/releases/latest/download/goliteflow-linux-amd64 -o goliteflow
chmod +x goliteflow

# Windows
curl -L https://github.com/sintakaridina/goliteflow/releases/latest/download/goliteflow-windows-amd64.exe -o goliteflow.exe
```

### Option 3: Docker

```bash
docker pull sintakaridina/goliteflow:latest
```

## Your First Workflow

Let's create a simple workflow to get you started.

### Step 1: Create a Workflow File

Create a file called `my-first-workflow.yml`:

```yaml
version: "1.0"
workflows:
  - name: hello_world
    schedule: "0 9 * * *"  # Daily at 9 AM
    tasks:
      - id: greet
        command: "echo 'Hello from GoliteFlow!'"
        retry: 1
      - id: log_time
        depends_on: ["greet"]
        command: "echo 'Current time: $(date)'"
        retry: 1
```

### Step 2: Validate Your Configuration

```bash
goliteflow validate --config=my-first-workflow.yml
```

You should see:
```
Configuration is valid!
Found 1 workflows:
  - hello_world (schedule: 0 9 * * *, tasks: 2)
```

### Step 3: Run Your Workflow

```bash
goliteflow run --config=my-first-workflow.yml
```

You should see output like:
```
Starting GoliteFlow
Loaded 1 workflows from my-first-workflow.yml
Scheduler started successfully
Running workflows once...
Executing workflow: hello_world
Workflow 'hello_world' completed with status: completed
Report generated: report.html
```

### Step 4: View the Report

Open the generated `report.html` file in your browser to see the execution details.

## CLI Commands

GoliteFlow provides several CLI commands:

### `run` - Execute Workflows

```bash
# Run once (default)
goliteflow run --config=workflows.yml

# Run as daemon (continuous)
goliteflow run --config=workflows.yml --daemon

# Enable verbose logging
goliteflow run --config=workflows.yml --verbose

# Specify output file
goliteflow run --config=workflows.yml --output=my-report.html
```

### `validate` - Check Configuration

```bash
# Validate configuration file
goliteflow validate --config=workflows.yml

# Enable verbose output
goliteflow validate --config=workflows.yml --verbose
```

### `report` - Generate Reports

```bash
# Generate HTML report
goliteflow report --output=report.html
```

## 📚 Using as a Go Library

### Basic Usage

```go
package main

import (
    "log"
    "github.com/sintakaridina/goliteflow"
)

func main() {
    // Simple execution
    err := goliteflow.Run("workflows.yml")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Advanced Usage

```go
package main

import (
    "context"
    "log"
    "time"
    "github.com/sintakaridina/goliteflow"
)

func main() {
    // Create GoliteFlow instance
    gf := goliteflow.New()
    
    // Load configuration
    err := gf.LoadConfig("workflows.yml")
    if err != nil {
        log.Fatal(err)
    }
    
    // Start scheduler
    err = gf.Start()
    if err != nil {
        log.Fatal(err)
    }
    defer gf.Stop()
    
    // Run with context for cancellation
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    err = gf.RunWithContext(ctx)
    if err != nil {
        log.Fatal(err)
    }
    
    // Generate report
    err = gf.GenerateReport("report.html")
    if err != nil {
        log.Fatal(err)
    }
}
```

## 🎨 Configuration Basics

### Workflow Structure

```yaml
version: "1.0"
workflows:
  - name: workflow_name
    schedule: "cron_expression"
    tasks:
      - id: task_id
        command: "command_to_execute"
        retry: number_of_retries
        depends_on: ["other_task_ids"]
        timeout: "duration"
```

### Cron Schedule Examples

```yaml
# Every minute
schedule: "* * * * *"

# Every hour
schedule: "0 * * * *"

# Daily at 2 AM
schedule: "0 2 * * *"

# Every Monday at 9 AM
schedule: "0 9 * * 1"

# Every 15 minutes
schedule: "*/15 * * * *"
```

### Task Dependencies

```yaml
tasks:
  - id: download_data
    command: "wget https://example.com/data.csv"
    retry: 3
  
  - id: process_data
    depends_on: ["download_data"]
    command: "python process.py data.csv"
    retry: 2
  
  - id: send_notification
    depends_on: ["process_data"]
    command: "curl -X POST https://hooks.slack.com/..."
    retry: 1
```

## 🐳 Docker Usage

### Using Pre-built Image

```bash
# Run with volume mount
docker run -v $(pwd):/workspace sintakaridina/goliteflow:latest run --config=/workspace/workflows.yml

# Run as daemon
docker run -d -v $(pwd):/workspace --name goliteflow sintakaridina/goliteflow:latest run --config=/workspace/workflows.yml --daemon
```

### Custom Dockerfile

```dockerfile
FROM sintakaridina/goliteflow:latest

# Copy your workflows
COPY workflows.yml /app/workflows.yml

# Run your workflows
CMD ["run", "--config=/app/workflows.yml", "--daemon"]
```

## 🔍 Troubleshooting

### Common Issues

**1. "command not found" error**
```bash
# Make sure goliteflow is in your PATH
which goliteflow

# Or use full path
/path/to/goliteflow run --config=workflows.yml
```

**2. "invalid cron expression" error**
```bash
# Check your cron syntax
# Use online cron validators like crontab.guru
```

**3. "permission denied" error**
```bash
# Make binary executable (Linux/macOS)
chmod +x goliteflow

# Or run with sudo if needed
sudo goliteflow run --config=workflows.yml
```

**4. "workflow failed" error**
```bash
# Check the generated report.html for detailed error logs
# Enable verbose logging
goliteflow run --config=workflows.yml --verbose
```

### Getting Help

- **Documentation**: Check the [Configuration Reference](/configuration)
- **Issues**: Report bugs on [GitHub Issues](https://github.com/sintakaridina/goliteflow/issues)
- **Discussions**: Ask questions in [GitHub Discussions](https://github.com/sintakaridina/goliteflow/discussions)
- **Email**: Contact the maintainer

## Next Steps

Now that you have GoliteFlow running, here's what you can do next:

1. **Explore Examples**: Check out [real-world examples](/examples)
2. **Learn Configuration**: Read the [Configuration Reference](/configuration)
3. **CLI Reference**: Learn all [CLI commands](/cli-reference)
4. **API Documentation**: Explore the [Go library API](/api)
5. **Deploy to Production**: Follow the [Deployment Guide](/deployment)

## 📚 Additional Resources

- [Configuration Reference](/configuration) - Detailed YAML configuration
- [CLI Reference](/cli-reference) - Complete command reference
- [Library API](/api) - Go library documentation
- [Examples](/examples) - Real-world use cases
- [Deployment Guide](/deployment) - Production deployment

---

<div class="next-steps">
  <h3>Ready for more?</h3>
  <p>Explore advanced configuration and real-world examples.</p>
  <a href="/configuration" class="btn btn-primary">Configuration Reference</a>
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
