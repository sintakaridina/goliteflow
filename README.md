# GoliteFlow

[![Build Status](https://github.com/sintakaridina/goliteflow/workflows/CI/badge.svg)](https://github.com/sintakaridina/goliteflow/actions)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.19-blue.svg)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/sintakaridina/goliteflow)](https://goreportcard.com/report/github.com/sintakaridina/goliteflow)
[![GitHub Downloads](https://img.shields.io/github/downloads/sintakaridina/goliteflow/total.svg)](https://github.com/sintakaridina/goliteflow/releases)

A lightweight workflow scheduler and task orchestrator designed for **any programming language**. GoliteFlow executes tasks/workflows defined in YAML files with retry logic, conditional execution, monitoring, and cron-based scheduling.

**🎯 Perfect for:** Python, Node.js, PHP, Java, Ruby, Go, or any application that needs scheduled tasks and monitoring.

## ✨ Features

- **🌍 Language Agnostic** - Works with any programming language or shell command
- **📅 Cron Scheduling** - Built-in scheduler with cron expressions
- **🔄 Dependency Management** - Tasks can depend on other tasks
- **🔁 Retry Logic** - Automatic retries with configurable backoff
- **📊 Beautiful HTML Reports** - Real-time dashboard with execution history
- **🚀 Lightweight** - Single binary, no external dependencies
- **🛠️ Production Ready** - Report archival, cleanup, and enterprise features

## 🚀 Quick Start (5 Minutes)

### Step 1: Download Binary

Choose your platform and download the latest release:

```bash
# Linux/macOS
curl -L https://github.com/sintakaridina/goliteflow/releases/latest/download/goliteflow-linux-amd64 -o goliteflow
chmod +x goliteflow

# Windows (PowerShell)
curl -L https://github.com/sintakaridina/goliteflow/releases/latest/download/goliteflow-windows-amd64.exe -o goliteflow.exe

# macOS (Apple Silicon)
curl -L https://github.com/sintakaridina/goliteflow/releases/latest/download/goliteflow-darwin-arm64 -o goliteflow
chmod +x goliteflow
```

### Step 2: Create Your First Workflow

Create a `my-workflow.yml` file:

```yaml
version: "1.0"
workflows:
  - name: hello_world_demo
    schedule: "@manual" # Run manually for testing
    tasks:
      - id: greet
        command: "echo Hello from GoliteFlow!"

      - id: date_time
        command: "date"
        depends_on: ["greet"]

      - id: system_info
        command: "uname -a"
        depends_on: ["date_time"]
```

### Step 3: Run Your Workflow

```bash
# Run the workflow once
./goliteflow run --config=my-workflow.yml

# Generate beautiful HTML report
./goliteflow report-enhanced --output=report.html

# Open report in browser
open report.html  # macOS
xdg-open report.html  # Linux
start report.html  # Windows
```

### Step 4: Run Continuously (Production)

```bash
# Start daemon for continuous monitoring
./goliteflow daemon --config=my-workflow.yml

# The daemon will:
# ✅ Monitor schedules 24/7
# ✅ Auto-execute workflows when scheduled
# ✅ Auto-generate HTML reports after each execution
# ✅ Handle failures with retry logic
# ✅ Manage report archival and cleanup
```

**🎉 That's it!** You now have a production-ready task scheduler with beautiful monitoring dashboard.

## 💼 Real-World Examples

### Python Data Pipeline

```yaml
version: "1.0"
workflows:
  - name: daily_etl
    schedule: "0 2 * * *" # Every day at 2 AM
    tasks:
      - id: extract_data
        command: "python scripts/extract_from_api.py"
        retry_count: 3

      - id: transform_data
        command: "python scripts/clean_and_transform.py"
        depends_on: ["extract_data"]

      - id: load_to_database
        command: "python scripts/load_to_postgres.py"
        depends_on: ["transform_data"]

      - id: send_report
        command: "python scripts/email_summary.py"
        depends_on: ["load_to_database"]
```

### Node.js API Health Monitoring

```yaml
version: "1.0"
workflows:
  - name: api_health_check
    schedule: "*/5 * * * *" # Every 5 minutes
    tasks:
      - id: check_api_health
        command: "node monitoring/health-check.js"

      - id: alert_on_failure
        command: "node monitoring/send-slack-alert.js"
        condition: "on_failure" # Only run if health check fails
```

### Multi-Language DevOps Pipeline

```yaml
version: "1.0"
workflows:
  - name: deployment_pipeline
    schedule: "0 0 * * 1" # Weekly on Monday
    tasks:
      - id: pull_latest_code
        command: "git pull origin main"

      - id: install_dependencies
        command: "npm install"
        depends_on: ["pull_latest_code"]

      - id: run_tests
        command: "python -m pytest tests/"
        depends_on: ["install_dependencies"]

      - id: build_application
        command: "npm run build"
        depends_on: ["run_tests"]

      - id: deploy_to_staging
        command: "bash deploy/staging.sh"
        depends_on: ["build_application"]
```

## 📊 Enhanced HTML Reports

GoliteFlow generates beautiful, production-ready HTML reports with:

### 📈 **Dashboard Features:**

- **Execution Timeline** - Visual timeline of task execution
- **Success/Failure Rates** - Statistics and trends
- **Task Dependencies** - Visual dependency graph
- **Performance Metrics** - Execution times and bottlenecks
- **Error Logs** - Detailed error messages and stack traces
- **Archive Management** - Automatic report rotation and cleanup

### 🔧 **Report Configuration:**

```yaml
# Configure report behavior in your workflow
version: "1.0"
reporting:
  max_executions: 50 # Keep latest 50 executions in main report
  archive_after_days: 30 # Archive reports older than 30 days
  cleanup_after_days: 90 # Delete archived reports after 90 days
  page_size: 20 # Executions per page

workflows:
  # ... your workflows
```

### 📋 **Report Management Commands:**

```bash
# Generate enhanced report (recommended)
./goliteflow report-enhanced --output=report.html

# Configure report limits
./goliteflow report-enhanced --max-executions=100 --page-size=25

# Manage report archives
./goliteflow report-manage stats      # Show statistics
./goliteflow report-manage archive    # Archive old reports
./goliteflow report-manage cleanup    # Clean up old archives
```

## 🛠️ Installation Options

### Option 1: Download Binary (Recommended)

Download the pre-built binary for your platform from [GitHub Releases](https://github.com/sintakaridina/goliteflow/releases/latest).

### Option 2: Install as Go Module

If you're building a Go application:

```bash
go get github.com/sintakaridina/goliteflow
```

### Option 3: Build from Source

```bash
git clone https://github.com/sintakaridina/goliteflow.git
cd goliteflow
make build
```

## ⚙️ Configuration Reference

### Basic Workflow Structure

```yaml
version: "1.0"
workflows:
  - name: "workflow_name"
    schedule: "0 */6 * * *" # Cron expression or @manual
    max_concurrent_tasks: 3
    timeout: "30m"

    tasks:
      - id: "task_1"
        command: "your-command here"
        working_dir: "/path/to/directory"
        env:
          API_KEY: "your-api-key"
          ENV: "production"
        timeout: "10m"
        retry_count: 3
        retry_delay: "5s"

      - id: "task_2"
        command: "another-command"
        depends_on: ["task_1"]
        condition: "on_success" # on_success, on_failure, always
```

### Schedule Expressions

| Expression    | Description       | Example        |
| ------------- | ----------------- | -------------- |
| `@manual`     | Run manually only | For testing    |
| `@daily`      | Once per day      | `0 0 * * *`    |
| `@hourly`     | Once per hour     | `0 * * * *`    |
| `*/5 * * * *` | Every 5 minutes   | API monitoring |
| `0 2 * * *`   | Daily at 2 AM     | Backup jobs    |
| `0 9 * * 1`   | Mondays at 9 AM   | Weekly reports |

### Report Management Configuration

| Setting              | Default            | Description                          |
| -------------------- | ------------------ | ------------------------------------ |
| `max_executions`     | 50                 | Max executions in main report        |
| `archive_after_days` | 30                 | Archive reports after N days         |
| `cleanup_after_days` | 90                 | Delete archived reports after N days |
| `page_size`          | 20                 | Executions per page                  |
| `report_dir`         | `reports/`         | Main report directory                |
| `archive_dir`        | `reports/archive/` | Archive directory                    |

## 🚀 Production Deployment

### Linux/macOS (systemd)

1. **Create service file:** `/etc/systemd/system/goliteflow.service`

```ini
[Unit]
Description=GoliteFlow Daemon
After=network.target

[Service]
Type=simple
User=goliteflow
WorkingDirectory=/opt/goliteflow
ExecStart=/opt/goliteflow/goliteflow daemon --config=/opt/goliteflow/production.yml
Restart=always
RestartSec=10

[Install]
WantedBy=multi-user.target
```

2. **Enable and start:**

```bash
sudo systemctl enable goliteflow
sudo systemctl start goliteflow
sudo systemctl status goliteflow
```

### Docker Deployment

```dockerfile
FROM alpine:latest
RUN apk --no-cache add ca-certificates python3 nodejs
WORKDIR /app
COPY goliteflow-linux-amd64 goliteflow
COPY config.yml .
RUN chmod +x goliteflow
CMD ["./goliteflow", "daemon", "--config=config.yml"]
```

### Windows Service

```powershell
# Install as Windows service
.\goliteflow.exe daemon --install --config=production.yml

# Start service
net start goliteflow
```

## 🔍 Monitoring & Troubleshooting

### Check Daemon Status

```bash
# View daemon logs
./goliteflow daemon --config=config.yml --log-level=debug

# Check specific workflow status
./goliteflow status --workflow=my_workflow

# View execution history
./goliteflow report-manage stats
```

### Common Issues

**Issue: Tasks failing with "command not found"**

```yaml
# Solution: Use full paths or set working directory
tasks:
  - command: "/usr/bin/python3 /full/path/to/script.py"
    working_dir: "/path/to/project"
```

**Issue: HTML report too slow to load**

```bash
# Solution: Reduce max executions in reports
./goliteflow report-enhanced --max-executions=25
```

**Issue: Disk space growing**

```bash
# Solution: Configure automatic cleanup
./goliteflow report-manage cleanup --older-than=60d
```

## 🤝 Integration Examples

### Integrate with Existing Applications

#### Python Flask App

```python
# In your Flask app
import subprocess

def trigger_goliteflow_task():
    result = subprocess.run(['./goliteflow', 'run', '--config=tasks.yml'])
    return result.returncode == 0
```

#### Node.js Express App

```javascript
// In your Express app
const { exec } = require("child_process");

app.post("/trigger-workflow", (req, res) => {
  exec("./goliteflow run --config=tasks.yml", (error, stdout, stderr) => {
    if (error) {
      res.status(500).json({ error: error.message });
    } else {
      res.json({ success: true, output: stdout });
    }
  });
});
```

### Web Dashboard Integration

Access reports via HTTP server:

```bash
# Start with web server
./goliteflow daemon --config=config.yml --web-port=8080

# Access dashboard at:
# http://localhost:8080/reports/latest.html
```

## 📚 Examples

Explore comprehensive examples in the [`examples/`](examples/) directory:

- **Quick Demo** (`quick-demo.yml`) - 2-minute setup test
- **File Processing** (`file-processing-workflow.yml`) - ETL pipeline with Python
- **API Monitoring** (`api-monitoring-workflow.yml`) - Health check automation
- **Backup & Cleanup** (`backup-cleanup-workflow.yml`) - Maintenance tasks

```bash
cd examples/
../goliteflow run --config=quick-demo.yml
../goliteflow report-enhanced --output=demo-report.html
```

## 💡 Tips for Beginners

### Start Simple

1. **Begin with `@manual` schedule** for testing
2. **Use simple commands** like `echo` first
3. **Check HTML reports** to understand execution flow
4. **Add complexity gradually** (dependencies, retries, etc.)

### Best Practices

- ✅ **Use full paths** in commands
- ✅ **Set working directories** for scripts
- ✅ **Handle errors gracefully** with conditions
- ✅ **Monitor with HTML reports** regularly
- ✅ **Use environment variables** for configuration
- ✅ **Test workflows manually** before scheduling

### Common Patterns

```yaml
# Error notification pattern
- id: main_task
  command: "python important_job.py"

- id: notify_on_error
  command: "python send_alert.py"
  depends_on: ["main_task"]
  condition: "on_failure"

# Backup pattern
- id: create_backup
  command: "mysqldump mydb > backup.sql"

- id: upload_backup
  command: "aws s3 cp backup.sql s3://backups/"
  depends_on: ["create_backup"]

- id: cleanup_local
  command: "rm backup.sql"
  depends_on: ["upload_backup"]
```

## 📖 Documentation

- **[Report Management Guide](docs/report-management.md)** - Comprehensive guide for production report management
- **[Examples Directory](examples/)** - Real-world workflow examples
- **[Build Instructions](BUILD.md)** - How to build from source

## 🆘 Support

- **GitHub Issues**: [Report bugs or request features](https://github.com/sintakaridina/goliteflow/issues)
- **Discussions**: [Ask questions and share workflows](https://github.com/sintakaridina/goliteflow/discussions)
- **Examples**: Check the `examples/` directory for real-world use cases

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

**🚀 Ready to get started?** Download the binary, create a simple YAML file, and run your first workflow in under 5 minutes!
