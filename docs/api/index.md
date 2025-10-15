---
layout: default
title: Library API
description: Go library API reference for GoliteFlow
---

# Library API Reference

This document provides a complete reference for GoliteFlow's Go library API.

## 📋 Table of Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [API Reference](#api-reference)
- [Examples](#examples)
- [Error Handling](#error-handling)

## 🚀 Installation

```bash
go get github.com/sintakaridina/goliteflow
```

## ⚡ Quick Start

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
}
```

## 📚 API Reference

### Package Functions

#### `Run(configFile string) error`

Execute workflows from a configuration file once.

**Parameters:**
- `configFile` (string): Path to the YAML configuration file

**Returns:**
- `error`: Any error that occurred during execution

**Example:**
```go
err := goliteflow.Run("workflows.yml")
if err != nil {
    log.Fatal(err)
}
```

#### `RunWithReport(configFile, reportFile string) error`

Execute workflows and generate an HTML report.

**Parameters:**
- `configFile` (string): Path to the YAML configuration file
- `reportFile` (string): Path for the output HTML report

**Returns:**
- `error`: Any error that occurred during execution

**Example:**
```go
err := goliteflow.RunWithReport("workflows.yml", "report.html")
if err != nil {
    log.Fatal(err)
}
```

#### `ValidateConfig(configFile string) error`

Validate a workflow configuration file.

**Parameters:**
- `configFile` (string): Path to the YAML configuration file

**Returns:**
- `error`: Any validation error

**Example:**
```go
err := goliteflow.ValidateConfig("workflows.yml")
if err != nil {
    log.Printf("Configuration error: %v", err)
}
```

### GoliteFlow Struct

The main struct for advanced usage and control.

#### `New() *GoliteFlow`

Create a new GoliteFlow instance.

**Returns:**
- `*GoliteFlow`: New GoliteFlow instance

**Example:**
```go
gf := goliteflow.New()
```

#### `LoadConfig(filename string) error`

Load workflow configuration from a YAML file.

**Parameters:**
- `filename` (string): Path to the YAML configuration file

**Returns:**
- `error`: Any error that occurred during loading

**Example:**
```go
gf := goliteflow.New()
err := gf.LoadConfig("workflows.yml")
if err != nil {
    log.Fatal(err)
}
```

#### `Start() error`

Start the workflow scheduler.

**Returns:**
- `error`: Any error that occurred during startup

**Example:**
```go
err := gf.Start()
if err != nil {
    log.Fatal(err)
}
defer gf.Stop()
```

#### `Stop()`

Stop the workflow scheduler.

**Example:**
```go
gf.Stop()
```

#### `Run() error`

Execute workflows once (non-daemon mode).

**Returns:**
- `error`: Any error that occurred during execution

**Example:**
```go
err := gf.Run()
if err != nil {
    log.Fatal(err)
}
```

#### `RunWithContext(ctx context.Context) error`

Execute workflows with a context for cancellation.

**Parameters:**
- `ctx` (context.Context): Context for cancellation and timeout

**Returns:**
- `error`: Any error that occurred during execution

**Example:**
```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

err := gf.RunWithContext(ctx)
if err != nil {
    log.Fatal(err)
}
```

#### `GenerateReport(outputFile string) error`

Generate an HTML report of execution history.

**Parameters:**
- `outputFile` (string): Path for the output HTML report

**Returns:**
- `error`: Any error that occurred during report generation

**Example:**
```go
err := gf.GenerateReport("report.html")
if err != nil {
    log.Fatal(err)
}
```

#### `GetStats() *SchedulerStats`

Get scheduler statistics.

**Returns:**
- `*SchedulerStats`: Scheduler statistics or nil if not started

**Example:**
```go
stats := gf.GetStats()
if stats != nil {
    log.Printf("Total workflows: %d", stats.TotalWorkflows)
    log.Printf("Successful executions: %d", stats.SuccessfulExecutions)
}
```

#### `GetExecutions(workflowName string) []WorkflowExecution`

Get execution history for a specific workflow.

**Parameters:**
- `workflowName` (string): Name of the workflow

**Returns:**
- `[]WorkflowExecution`: List of executions for the workflow

**Example:**
```go
executions := gf.GetExecutions("daily_backup")
for _, exec := range executions {
    log.Printf("Execution: %s, Status: %s", exec.StartTime, exec.Status)
}
```

#### `GetNextRunTimes() map[string]time.Time`

Get the next scheduled run times for all workflows.

**Returns:**
- `map[string]time.Time`: Map of workflow names to next run times

**Example:**
```go
nextRuns := gf.GetNextRunTimes()
for workflow, nextRun := range nextRuns {
    log.Printf("Next run for %s: %s", workflow, nextRun)
}
```

#### `SetLogLevel(level zerolog.Level)`

Set the logging level.

**Parameters:**
- `level` (zerolog.Level): Logging level (Debug, Info, Warn, Error, Fatal)

**Example:**
```go
gf.SetLogLevel(zerolog.DebugLevel)
```

#### `GetLogger() *Logger`

Get the logger instance.

**Returns:**
- `*Logger`: Logger instance

**Example:**
```go
logger := gf.GetLogger()
logger.Info("Custom log message")
```

### Data Types

#### `SchedulerStats`

Statistics about the scheduler.

```go
type SchedulerStats struct {
    TotalWorkflows       int                    `json:"total_workflows"`
    TotalExecutions      int                    `json:"total_executions"`
    SuccessfulExecutions int                    `json:"successful_executions"`
    FailedExecutions     int                    `json:"failed_executions"`
    NextRuns             map[string]time.Time   `json:"next_runs"`
}
```

#### `WorkflowExecution`

Execution state of a workflow.

```go
type WorkflowExecution struct {
    WorkflowID   string            `json:"workflow_id"`
    StartTime    time.Time         `json:"start_time"`
    EndTime      time.Time         `json:"end_time"`
    Duration     time.Duration     `json:"duration"`
    Status       string            `json:"status"` // running, completed, failed
    TaskResults  []ExecutionResult `json:"task_results"`
    ErrorMessage string            `json:"error_message,omitempty"`
}
```

#### `ExecutionResult`

Result of a task execution.

```go
type ExecutionResult struct {
    TaskID      string        `json:"task_id"`
    WorkflowID  string        `json:"workflow_id"`
    StartTime   time.Time     `json:"start_time"`
    EndTime     time.Time     `json:"end_time"`
    Duration    time.Duration `json:"duration"`
    ExitCode    int           `json:"exit_code"`
    Success     bool          `json:"success"`
    RetryCount  int           `json:"retry_count"`
    Stdout      string        `json:"stdout"`
    Stderr      string        `json:"stderr"`
    Error       string        `json:"error,omitempty"`
}
```

## 🔧 Examples

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
    "github.com/rs/zerolog"
    "github.com/sintakaridina/goliteflow"
)

func main() {
    // Create instance
    gf := goliteflow.New()
    
    // Set log level
    gf.SetLogLevel(zerolog.DebugLevel)
    
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
    
    // Run with timeout
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
    
    // Print statistics
    stats := gf.GetStats()
    if stats != nil {
        log.Printf("Total workflows: %d", stats.TotalWorkflows)
        log.Printf("Successful executions: %d", stats.SuccessfulExecutions)
    }
}
```

### Daemon Mode

```go
package main

import (
    "log"
    "os"
    "os/signal"
    "syscall"
    "github.com/sintakaridina/goliteflow"
)

func main() {
    gf := goliteflow.New()
    
    err := gf.LoadConfig("workflows.yml")
    if err != nil {
        log.Fatal(err)
    }
    
    err = gf.Start()
    if err != nil {
        log.Fatal(err)
    }
    defer gf.Stop()
    
    // Wait for signal
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    
    log.Println("GoliteFlow daemon started. Press Ctrl+C to stop.")
    <-sigChan
    log.Println("Shutting down...")
}
```

### Custom Logging

```go
package main

import (
    "log"
    "github.com/rs/zerolog"
    "github.com/sintakaridina/goliteflow"
)

func main() {
    gf := goliteflow.New()
    
    // Get logger and add custom fields
    logger := gf.GetLogger()
    logger = logger.WithField("service", "my-app")
    logger.Info("Starting application")
    
    err := gf.LoadConfig("workflows.yml")
    if err != nil {
        logger.Error("Failed to load config", "error", err)
        return
    }
    
    err = gf.Run()
    if err != nil {
        logger.Error("Failed to run workflows", "error", err)
        return
    }
    
    logger.Info("Workflows completed successfully")
}
```

### Monitoring and Statistics

```go
package main

import (
    "log"
    "time"
    "github.com/sintakaridina/goliteflow"
)

func main() {
    gf := goliteflow.New()
    
    err := gf.LoadConfig("workflows.yml")
    if err != nil {
        log.Fatal(err)
    }
    
    err = gf.Start()
    if err != nil {
        log.Fatal(err)
    }
    defer gf.Stop()
    
    // Monitor execution
    go func() {
        ticker := time.NewTicker(5 * time.Minute)
        defer ticker.Stop()
        
        for range ticker.C {
            stats := gf.GetStats()
            if stats != nil {
                log.Printf("Stats - Workflows: %d, Executions: %d, Success: %d, Failed: %d",
                    stats.TotalWorkflows,
                    stats.TotalExecutions,
                    stats.SuccessfulExecutions,
                    stats.FailedExecutions)
            }
        }
    }()
    
    // Run workflows
    err = gf.Run()
    if err != nil {
        log.Fatal(err)
    }
}
```

## ⚠️ Error Handling

### Common Errors

**Configuration Errors:**
```go
err := gf.LoadConfig("workflows.yml")
if err != nil {
    // Handle configuration errors
    log.Printf("Configuration error: %v", err)
    return
}
```

**Execution Errors:**
```go
err := gf.Run()
if err != nil {
    // Handle execution errors
    log.Printf("Execution error: %v", err)
    return
}
```

**Context Cancellation:**
```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

err := gf.RunWithContext(ctx)
if err == context.DeadlineExceeded {
    log.Println("Execution timed out")
} else if err == context.Canceled {
    log.Println("Execution was canceled")
} else if err != nil {
    log.Printf("Execution error: %v", err)
}
```

### Best Practices

1. **Always check errors** from all API calls
2. **Use context** for cancellation and timeouts
3. **Set appropriate log levels** for your use case
4. **Handle graceful shutdown** with signal handling
5. **Monitor statistics** for production deployments

## 📚 Related Documentation

- [Getting Started](/getting-started) - Quick setup guide
- [Configuration Reference](/configuration) - YAML configuration details
- [CLI Reference](/cli-reference) - Command-line interface
- [Examples](/examples) - Real-world use cases
- [Deployment Guide](/deployment) - Production deployment

---

<div class="next-steps">
  <h3>Ready to use the library?</h3>
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
