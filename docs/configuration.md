---
layout: default
title: Configuration Reference
description: Complete YAML configuration reference for GoliteFlow
---

# Configuration Reference

This document provides a complete reference for GoliteFlow's YAML configuration format.

## 📋 Table of Contents

- [Basic Structure](#basic-structure)
- [Workflow Configuration](#workflow-configuration)
- [Task Configuration](#task-configuration)
- [Cron Schedule Format](#cron-schedule-format)
- [Examples](#examples)
- [Validation Rules](#validation-rules)

## 🏗️ Basic Structure

```yaml
version: "1.0"
workflows:
  - name: workflow_name
    schedule: "cron_expression"
    tasks:
      - id: task_id
        command: "command_to_execute"
        # ... other task options
```

### Root Level Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `version` | string | ✅ | Configuration version (currently "1.0") |
| `workflows` | array | ✅ | List of workflow definitions |

## 🔄 Workflow Configuration

Each workflow represents a collection of tasks that run on a schedule.

```yaml
workflows:
  - name: daily_backup
    schedule: "0 2 * * *"
    tasks:
      - id: backup_files
        command: "tar -czf backup.tar.gz /data"
        retry: 3
```

### Workflow Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | ✅ | Unique workflow identifier |
| `schedule` | string | ✅ | Cron expression for scheduling |
| `tasks` | array | ✅ | List of tasks to execute |

## ⚙️ Task Configuration

Tasks are individual commands that execute within a workflow.

```yaml
tasks:
  - id: download_data
    command: "curl -s https://api.example.com/data"
    retry: 3
    timeout: "30s"
    depends_on: ["previous_task"]
```

### Task Fields

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `id` | string | ✅ | - | Unique task identifier |
| `command` | string | ✅ | - | Command to execute |
| `retry` | integer | ❌ | 1 | Number of retry attempts |
| `timeout` | string | ❌ | "30m" | Task timeout duration |
| `depends_on` | array | ❌ | [] | List of task IDs this task depends on |

### Task Dependencies

Tasks can depend on other tasks using the `depends_on` field:

```yaml
tasks:
  - id: step1
    command: "echo 'Step 1'"
  
  - id: step2
    depends_on: ["step1"]
    command: "echo 'Step 2'"
  
  - id: step3
    depends_on: ["step1", "step2"]
    command: "echo 'Step 3'"
```

**Execution Order**: step1 → step2 → step3

### Retry Configuration

Tasks can be configured to retry on failure:

```yaml
tasks:
  - id: unreliable_task
    command: "curl -f https://unreliable-api.com/data"
    retry: 5  # Will retry up to 5 times
```

**Retry Behavior**:
- Exponential backoff: 1s, 2s, 4s, 8s, 16s...
- Maximum backoff: 5 minutes
- All retries must fail for task to be marked as failed

### Timeout Configuration

Tasks can have custom timeouts:

```yaml
tasks:
  - id: long_running_task
    command: "python long_script.py"
    timeout: "2h"  # 2 hours timeout
```

**Timeout Format**: Go duration format
- `"30s"` - 30 seconds
- `"5m"` - 5 minutes
- `"1h"` - 1 hour
- `"2h30m"` - 2 hours 30 minutes

## ⏰ Cron Schedule Format

GoliteFlow uses standard cron format with 5 fields:

```
┌───────────── minute (0 - 59)
│ ┌───────────── hour (0 - 23)
│ │ ┌───────────── day of month (1 - 31)
│ │ │ ┌───────────── month (1 - 12)
│ │ │ │ ┌───────────── day of week (0 - 6) (Sunday to Saturday)
│ │ │ │ │
* * * * *
```

### Common Schedule Examples

| Schedule | Description |
|----------|-------------|
| `"* * * * *"` | Every minute |
| `"0 * * * *"` | Every hour |
| `"0 0 * * *"` | Daily at midnight |
| `"0 9 * * *"` | Daily at 9 AM |
| `"0 2 * * 0"` | Every Sunday at 2 AM |
| `"*/15 * * * *"` | Every 15 minutes |
| `"0 0 1 * *"` | Monthly on the 1st at midnight |
| `"0 9 * * 1-5"` | Weekdays at 9 AM |
| `"0 0 * * 0"` | Every Sunday at midnight |

### Special Characters

| Character | Description | Example |
|-----------|-------------|---------|
| `*` | Any value | `* * * * *` (every minute) |
| `,` | List of values | `0 9,17 * * *` (9 AM and 5 PM) |
| `-` | Range of values | `0 9-17 * * *` (9 AM to 5 PM) |
| `/` | Step values | `*/15 * * * *` (every 15 minutes) |

## 📝 Examples

### Basic Workflow

```yaml
version: "1.0"
workflows:
  - name: simple_backup
    schedule: "0 2 * * *"
    tasks:
      - id: create_backup
        command: "tar -czf backup-$(date +%Y%m%d).tar.gz /data"
        retry: 2
        timeout: "1h"
```

### Complex Workflow with Dependencies

```yaml
version: "1.0"
workflows:
  - name: data_pipeline
    schedule: "0 3 * * *"
    tasks:
      - id: download_data
        command: "wget -O data.csv https://api.example.com/export"
        retry: 3
        timeout: "30m"
      
      - id: validate_data
        depends_on: ["download_data"]
        command: "python validate.py data.csv"
        retry: 2
        timeout: "10m"
      
      - id: process_data
        depends_on: ["validate_data"]
        command: "python process.py data.csv"
        retry: 2
        timeout: "1h"
      
      - id: upload_results
        depends_on: ["process_data"]
        command: "aws s3 cp results.json s3://my-bucket/"
        retry: 3
        timeout: "15m"
      
      - id: send_notification
        depends_on: ["upload_results"]
        command: "curl -X POST https://hooks.slack.com/services/..."
        retry: 1
        timeout: "30s"
```

### Multiple Workflows

```yaml
version: "1.0"
workflows:
  - name: hourly_cleanup
    schedule: "0 * * * *"
    tasks:
      - id: cleanup_temp
        command: "rm -rf /tmp/old_files"
        retry: 1
        timeout: "5m"
  
  - name: daily_backup
    schedule: "0 2 * * *"
    tasks:
      - id: backup_database
        command: "pg_dump mydb > backup.sql"
        retry: 2
        timeout: "30m"
      
      - id: compress_backup
        depends_on: ["backup_database"]
        command: "gzip backup.sql"
        retry: 1
        timeout: "5m"
  
  - name: weekly_report
    schedule: "0 9 * * 1"
    tasks:
      - id: generate_report
        command: "python generate_weekly_report.py"
        retry: 2
        timeout: "2h"
      
      - id: email_report
        depends_on: ["generate_report"]
        command: "mail -s 'Weekly Report' admin@company.com < report.pdf"
        retry: 1
        timeout: "1m"
```

### Error Handling Example

```yaml
version: "1.0"
workflows:
  - name: robust_pipeline
    schedule: "0 4 * * *"
    tasks:
      - id: fetch_data
        command: "curl -f https://unreliable-api.com/data"
        retry: 5  # Retry up to 5 times
        timeout: "2m"
      
      - id: process_data
        depends_on: ["fetch_data"]
        command: "python process.py"
        retry: 3
        timeout: "30m"
      
      - id: fallback_notification
        depends_on: ["fetch_data"]
        command: "echo 'Data fetch failed, using fallback'"
        retry: 1
        timeout: "10s"
```

## ✅ Validation Rules

### Configuration Validation

- **Version**: Must be "1.0"
- **Workflows**: Must have at least one workflow
- **Workflow Names**: Must be unique within the configuration
- **Task IDs**: Must be unique within each workflow

### Workflow Validation

- **Name**: Required, non-empty string
- **Schedule**: Required, valid cron expression
- **Tasks**: Must have at least one task

### Task Validation

- **ID**: Required, non-empty string
- **Command**: Required, non-empty string
- **Retry**: Must be non-negative integer
- **Timeout**: Must be valid Go duration format
- **Depends On**: Must reference existing task IDs in the same workflow

### Dependency Validation

- **No Circular Dependencies**: Tasks cannot depend on themselves directly or indirectly
- **Valid References**: All dependencies must reference existing task IDs
- **Same Workflow**: Dependencies must be within the same workflow

## 🔧 Best Practices

### 1. Naming Conventions

```yaml
# Good: Descriptive names
- name: daily_database_backup
  tasks:
    - id: backup_postgres
    - id: backup_redis

# Avoid: Generic names
- name: workflow1
  tasks:
    - id: task1
    - id: task2
```

### 2. Error Handling

```yaml
# Good: Appropriate retry counts
- id: api_call
  command: "curl -f https://api.example.com/data"
  retry: 3  # Reasonable for network calls

- id: file_operation
  command: "cp file.txt backup/"
  retry: 1  # File operations usually succeed or fail immediately
```

### 3. Timeout Configuration

```yaml
# Good: Realistic timeouts
- id: download_large_file
  command: "wget https://example.com/large-file.zip"
  timeout: "30m"  # Allow time for large downloads

- id: quick_validation
  command: "python validate.py"
  timeout: "2m"   # Quick operations
```

### 4. Dependency Design

```yaml
# Good: Clear dependency chain
tasks:
  - id: download
    command: "wget data.csv"
  
  - id: validate
    depends_on: ["download"]
    command: "python validate.py data.csv"
  
  - id: process
    depends_on: ["validate"]
    command: "python process.py data.csv"
```

## 🚨 Common Mistakes

### 1. Circular Dependencies

```yaml
# ❌ Wrong: Circular dependency
tasks:
  - id: task_a
    depends_on: ["task_b"]
    command: "echo A"
  
  - id: task_b
    depends_on: ["task_a"]
    command: "echo B"
```

### 2. Invalid Cron Expressions

```yaml
# ❌ Wrong: Invalid cron format
schedule: "every day at 9am"  # Not valid cron

# ✅ Correct: Valid cron format
schedule: "0 9 * * *"  # Daily at 9 AM
```

### 3. Missing Dependencies

```yaml
# ❌ Wrong: Task depends on non-existent task
tasks:
  - id: process_data
    depends_on: ["download_data"]  # download_data doesn't exist
    command: "python process.py"
```

### 4. Invalid Timeout Format

```yaml
# ❌ Wrong: Invalid duration format
timeout: "2 hours"  # Not valid Go duration

# ✅ Correct: Valid duration format
timeout: "2h"  # 2 hours
```

---

<div class="next-steps">
  <h3>Ready to create your workflows?</h3>
  <p>Check out real-world examples and start building.</p>
  <a href="/examples" class="btn btn-primary">View Examples</a>
  <a href="/cli-reference" class="btn btn-secondary">CLI Reference</a>
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
