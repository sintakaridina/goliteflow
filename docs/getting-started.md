---
layout: default
title: Getting Started - 5 Minute Setup
description: Complete beginner guide for GoliteFlow. Learn how to install, configure, and run workflows for any progra## Real-World Examples

### Python Data Pipelineng language in 5 minutes.
keywords: goliteflow installation, goliteflow setup, python workflows, nodejs automation, workflow scheduler, task orchestrator
author: GoliteFlow Team
---

# Getting Started (5 Minutes)

Get up and running with GoliteFlow in just **5 minutes**. Works with **any programming language** - Python, Node.js, PHP, Java, Ruby, Go, or shell commands **Need Real Examples?**

1. [Python data pipelines](https://github.com/sintakaridina/goliteflow/tree/main/examples)
2. [Node.js API monitoring](https://github.com/sintakaridina/goliteflow/tree/main/examples)
3. [DevOps automation](https://github.com/sintakaridina/goliteflow/tree/main/examples)

---

<div class="next-steps-cta">
  <h3>Congratulations!</h3>
  <p>You now have GoliteFlow running with enhanced reports and production features.</p>
  <div class="cta-buttons">
    <a href="/goliteflow/configuration" class="btn btn-primary">Advanced Configuration</a>
    <a href="/goliteflow/report-management" class="btn btn-secondary">Report Management</a>
  </div>
</div>

## Prerequisites

- **Any programming language** (Python, Node.js, PHP, etc.) - Optional
- **Basic YAML knowledge** - We'll teach you as we go
- **Command line access** - Terminal, PowerShell, or CMD

**No Go installation required!** GoliteFlow is distributed as a single binary.

## Step 1: Download Binary (1 minute)

Choose your platform and run **one command**:

### Linux/macOS

```bash
curl -L https://github.com/sintakaridina/goliteflow/releases/latest/download/goliteflow-linux-amd64 -o goliteflow && chmod +x goliteflow
```

### Windows (PowerShell)

```powershell
curl -L https://github.com/sintakaridina/goliteflow/releases/latest/download/goliteflow-windows-amd64.exe -o goliteflow.exe
```

### macOS (Apple Silicon)

```bash
curl -L https://github.com/sintakaridina/goliteflow/releases/latest/download/goliteflow-darwin-arm64 -o goliteflow && chmod +x goliteflow
```

### Docker (Optional)

```bash
docker pull sintakaridina/goliteflow:latest
```

**Done!** You now have GoliteFlow installed. No dependencies, no configuration needed.

## Step 2: Create Your First Workflow (2 minutes)

Create a file called `hello-world.yml`:

```yaml
version: "1.0"
workflows:
  - name: hello_world_demo
    schedule: "@manual" # Run manually for testing
    tasks:
      - id: greet
        command: "echo Hello from GoliteFlow!"

      - id: check_system
        command: "echo System info && date"
        depends_on: ["greet"]

      - id: list_files
        command: "ls -la" # Use "dir" on Windows
        depends_on: ["check_system"]
```

**Explanation:**

- `@manual` - Run only when triggered manually (great for testing)
- `depends_on` - Task runs only after specified tasks complete
- `command` - Any shell command (works with Python, Node.js, etc.)

## Step 3: Run Your Workflow (1 minute)

### Test Your Configuration

```bash
./goliteflow validate --config=hello-world.yml
```

### Run the Workflow

```bash
./goliteflow run --config=hello-world.yml
```

**Expected output:**

```
Starting GoliteFlow
Loaded 1 workflows from hello-world.yml
Executing workflow: hello_world_demo
Workflow 'hello_world_demo' completed with status: completed
Report generated: report.html
```

### Step 4: View the Report

Open the generated `report.html` file in your browser to see the execution details.

## CLI Commands

## Step 4: View Enhanced Reports (1 minute)

Generate a **production-ready HTML dashboard**:

```bash
# Generate enhanced report (recommended)
./goliteflow report-enhanced --output=dashboard.html

# Open in your browser
open dashboard.html      # macOS
xdg-open dashboard.html  # Linux
start dashboard.html     # Windows
```

**You'll see:**

- **Beautiful Dashboard** - Modern, responsive design
- **Execution Statistics** - Success rates, timing, trends
- **Interactive Timeline** - Click to expand task details
- **Task Details** - Individual results, logs, retry attempts
- **Auto Management** - Reports stay fast with archival

## Step 5: Production Mode (1 minute)

Start **continuous monitoring** with automatic report updates:

```bash
# Update schedule to run continuously
# Edit hello-world.yml and change:
# schedule: "@manual"        # Remove this
# schedule: "*/5 * * * *"    # Add this (every 5 minutes)

# Start daemon mode
./goliteflow daemon --config=hello-world.yml
```

**Now it runs automatically!**

- Executes every 5 minutes
- Updates HTML report after each run
- Archives old data automatically
- Cleans up storage automatically

## Enhanced CLI Commands

### Core Commands

```bash
# Run workflows
./goliteflow run --config=my-workflow.yml

# Start continuous daemon
./goliteflow daemon --config=my-workflow.yml

# Validate configuration
./goliteflow validate --config=my-workflow.yml
```

### Enhanced Report Commands

```bash
# Generate enhanced dashboard (recommended)
./goliteflow report-enhanced --output=dashboard.html

# Configure report limits
./goliteflow report-enhanced --max-executions=100 --page-size=25

# Manage report archives
./goliteflow report-manage stats      # View statistics
./goliteflow report-manage archive    # Archive old data
./goliteflow report-manage cleanup    # Clean up storage
```

## Real-World Examples

### Python Data Pipeline

Create `python-etl.yml`:

```yaml
version: "1.0"
workflows:
  - name: daily_data_processing
    schedule: "0 2 * * *" # Daily at 2 AM
    tasks:
      - id: extract_data
        command: "python scripts/extract_from_api.py"
        retry_count: 3
        timeout: "10m"

      - id: transform_data
        command: "python scripts/clean_and_transform.py"
        depends_on: ["extract_data"]
        timeout: "15m"

      - id: load_to_db
        command: "python scripts/load_to_postgres.py"
        depends_on: ["transform_data"]

      - id: send_report
        command: "python scripts/email_daily_summary.py"
        depends_on: ["load_to_db"]
        condition: "on_success"
```

### Node.js API Monitoring

Create `nodejs-monitoring.yml`:

```yaml
version: "1.0"
workflows:
  - name: api_health_check
    schedule: "*/5 * * * *" # Every 5 minutes
    tasks:
      - id: check_main_api
        command: "node monitoring/check-api-health.js"

      - id: check_database
        command: "node monitoring/check-db-connection.js"

      - id: alert_on_failure
        command: "node monitoring/send-slack-alert.js"
        condition: "on_failure" # Only if health checks fail
```

### PHP Application Tasks

Create `php-maintenance.yml`:

```yaml
version: "1.0"
workflows:
  - name: weekly_maintenance
    schedule: "0 3 * * 0" # Sundays at 3 AM
    tasks:
      - id: clear_cache
        command: "php artisan cache:clear"

      - id: optimize_database
        command: "php artisan migrate:status && php artisan db:optimize"
        depends_on: ["clear_cache"]

      - id: backup_files
        command: "php scripts/backup-files.php"
        depends_on: ["optimize_database"]
```

### Java Batch Processing

Create `java-batch.yml`:

```yaml
version: "1.0"
workflows:
  - name: monthly_reports
    schedule: "0 1 1 * *" # 1st day of month at 1 AM
    tasks:
      - id: generate_reports
        command: "java -jar batch-processor.jar --monthly-report"
        timeout: "30m"

      - id: compress_files
        command: "java -jar file-compressor.jar --input=reports/ --output=archive/"
        depends_on: ["generate_reports"]
```

## Go Library Integration (Optional)

If you're building a Go application, you can integrate GoliteFlow directly:

```go
package main

import (
    "log"
    "github.com/sintakaridina/goliteflow"
)

func main() {
    // Execute workflow from Go code
    err := goliteflow.Run("workflows.yml")
    if err != nil {
        log.Fatal("Workflow failed:", err)
    }
}

```

## Configuration Tips

### Schedule Patterns

| Pattern       | Description       | Example Use Case            |
| ------------- | ----------------- | --------------------------- |
| `@manual`     | Run manually only | Testing, debugging          |
| `@daily`      | Once per day      | Daily reports, backups      |
| `@hourly`     | Once per hour     | Log rotation, health checks |
| `*/5 * * * *` | Every 5 minutes   | API monitoring              |
| `0 2 * * *`   | Daily at 2 AM     | ETL jobs, heavy processing  |
| `0 9 * * 1`   | Mondays at 9 AM   | Weekly reports              |

### Task Dependencies

```yaml
# Sequential processing
tasks:
  - id: step1
    command: "python extract.py"
  - id: step2
    depends_on: ["step1"]
    command: "python transform.py"
  - id: step3
    depends_on: ["step2"]
    command: "python load.py"

# Parallel + Final
tasks:
  - id: task_a
    command: "python process_a.py"
  - id: task_b
    command: "python process_b.py"
  - id: combine
    depends_on: ["task_a", "task_b"]
    command: "python combine_results.py"
```

### Error Handling

```yaml
tasks:
  - id: main_job
    command: "python important_task.py"
    retry_count: 3
    timeout: "30m"

  - id: cleanup_on_success
    depends_on: ["main_job"]
    command: "python cleanup.py"
    condition: "on_success"

  - id: alert_on_failure
    depends_on: ["main_job"]
    command: "python send_alert.py"
    condition: "on_failure"
```

## Docker Usage

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

## Troubleshooting

### Common Issues

**1. "command not found" error**

````bash
## Production Deployment

### Linux/macOS (systemd)
```bash
# Create service file: /etc/systemd/system/goliteflow.service
[Unit]
Description=GoliteFlow Daemon
After=network.target

[Service]
Type=simple
User=goliteflow
WorkingDirectory=/opt/goliteflow
ExecStart=/opt/goliteflow/goliteflow daemon --config=/opt/goliteflow/production.yml
Restart=always

[Install]
WantedBy=multi-user.target

# Enable and start
sudo systemctl enable goliteflow
sudo systemctl start goliteflow
````

### Windows Service

```powershell
# Install as Windows service
.\goliteflow.exe daemon --install --config=production.yml
net start goliteflow
```

### Docker Production

```bash
# Create production image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY goliteflow-linux-amd64 goliteflow
COPY production.yml .
RUN chmod +x goliteflow
CMD ["./goliteflow", "daemon", "--config=production.yml"]
```

## Troubleshooting

### Common Issues & Solutions

**Issue: "command not found" error**

```yaml
# Problem: Relative paths
command: "python script.py"

# Solution: Full paths
command: "/usr/bin/python3 /full/path/to/script.py"
working_dir: "/path/to/project"
which goliteflow

# Or use full path
/path/to/goliteflow run --config=workflows.yml
```

**2. "invalid cron expression" error**

**Issue: Invalid cron expression**

```bash
# Problem: Wrong cron syntax
schedule: "0 25 * * *"  # Invalid: 25 is not valid hour

# Solution: Valid cron expressions
schedule: "0 2 * * *"   # Daily at 2 AM
schedule: "*/15 * * * *" # Every 15 minutes
# Use crontab.guru to validate expressions
```

**Issue: HTML report too slow**

```bash
# Problem: Too many executions in report
./goliteflow report-enhanced --output=report.html

# Solution: Limit executions
./goliteflow report-enhanced --max-executions=25 --output=report.html
```

**Issue: Task fails with "permission denied"**

```yaml
# Problem: Script not executable
tasks:
  - command: "./script.sh"

# Solution: Make executable or use interpreter
tasks:
  - command: "chmod +x script.sh && ./script.sh"
  # OR
  - command: "bash script.sh"
```

**Issue: Python/Node.js not found**

```yaml
# Problem: PATH issues
tasks:
  - command: "python script.py"

# Solution: Full paths or set environment
tasks:
  - command: "/usr/bin/python3 script.py"
  - env:
      PATH: "/usr/local/bin:/usr/bin:/bin"
```

### Getting Help

- **[Configuration Reference](/goliteflow/configuration)** - Complete YAML guide
- **[CLI Reference](/goliteflow/cli-reference)** - All commands explained
- **[Report Management](/goliteflow/report-management)** - Production scaling
- **[GitHub Issues](https://github.com/sintakaridina/goliteflow/issues)** - Report bugs
- **[Discussions](https://github.com/sintakaridina/goliteflow/discussions)** - Ask questions

## Next Steps

**Choose your path:**

### **Ready for Production?**

1. [Deploy as system service](/goliteflow/configuration#production-deployment)
2. [Configure report management](/goliteflow/report-management)
3. [Set up monitoring and alerts](/goliteflow/configuration#monitoring)

### **Want to Customize?**

1. [Advanced YAML configuration](/goliteflow/configuration)
2. [Report customization options](/goliteflow/report-management)
3. [Integration with existing apps](/goliteflow/cli-reference)

### � **Need Real Examples?**

1. [Python data pipelines](https://github.com/sintakaridina/goliteflow/tree/main/examples)
2. [Node.js API monitoring](https://github.com/sintakaridina/goliteflow/tree/main/examples)
3. [DevOps automation](https://github.com/sintakaridina/goliteflow/tree/main/examples)

---

<div class="next-steps-cta">
  <h3>🎉 Congratulations!</h3>
  <p>You now have GoliteFlow running with enhanced reports and production features.</p>
  <div class="cta-buttons">
    <a href="/goliteflow/configuration" class="btn btn-primary">⚙️ Advanced Configuration</a>
    <a href="/goliteflow/report-management" class="btn btn-secondary">📊 Report Management</a>
  </div>
</div>
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
