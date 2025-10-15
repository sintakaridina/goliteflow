---
layout: default
title: CLI Reference - Complete Command Guide
description: Complete command-line interface reference for GoliteFlow. All commands, enhanced reports, production options, and examples for any programming language.
keywords: goliteflow cli, goliteflow commands, workflow orchestrator, enhanced reports, python automation, nodejs scheduler
author: GoliteFlow Team
---

# 🖥️ CLI Reference

Complete command-line interface reference for GoliteFlow. Works with **any programming language** - Python, Node.js, PHP, Java, Ruby, Go, and shell commands.

## 📋 Quick Navigation

- [🚀 Installation](#installation)
- [⚙️ Core Commands](#core-commands)
- [📊 Enhanced Reports](#enhanced-reports)
- [🛠️ Report Management](#report-management)
- [💼 Real Examples](#real-examples)
- [📖 Configuration](#configuration)

## 🚀 Installation

**Choose your platform:**

```bash
# Linux/macOS
curl -L https://github.com/sintakaridina/goliteflow/releases/latest/download/goliteflow-linux-amd64 -o goliteflow && chmod +x goliteflow

# Windows (PowerShell)
curl -L https://github.com/sintakaridina/goliteflow/releases/latest/download/goliteflow-windows-amd64.exe -o goliteflow.exe

# macOS (Apple Silicon)
curl -L https://github.com/sintakaridina/goliteflow/releases/latest/download/goliteflow-darwin-arm64 -o goliteflow && chmod +x goliteflow
```

## ⚙️ Core Commands

### 🏃 `run` - Execute Workflows

Run workflows once or continuously with automatic scheduling.

**Syntax:**

```bash
./goliteflow run --config=<file> [options]
```

**Options:**
| Option | Description | Default |
|--------|-------------|---------|
| `--config`, `-c` | YAML configuration file | `lite-workflows.yml` |
| `--verbose`, `-v` | Enable debug logging | `false` |

**Examples:**

```bash
# Run workflows once (testing)
./goliteflow run --config=my-workflow.yml

# Run with debug logs
./goliteflow run --config=my-workflow.yml --verbose

# Works with any language
./goliteflow run --config=python-etl.yml
./goliteflow run --config=nodejs-api.yml
./goliteflow run --config=php-maintenance.yml
```

### � `daemon` - Continuous Mode

Start continuous monitoring with automatic scheduling and report updates.

**Syntax:**

```bash
./goliteflow daemon --config=<file> [options]
```

**Features:**

- ⏰ **Auto Scheduling** - Executes workflows based on cron expressions
- 📊 **Auto Reports** - Updates HTML reports after each execution
- 🔄 **Continuous** - Runs 24/7 in background
- 🛡️ **Resilient** - Automatic restarts on errors

**Examples:**

```bash
# Start production daemon
./goliteflow daemon --config=production.yml

# Daemon with debug logging
./goliteflow daemon --config=production.yml --verbose
```

### ✅ `validate` - Configuration Check

Validate YAML configuration before running workflows.

**Syntax:**

```bash
./goliteflow validate --config=<file>
```

**Examples:**

```bash
# Validate configuration
./goliteflow validate --config=my-workflow.yml

# Validate with verbose output
./goliteflow validate --config=my-workflow.yml --verbose
```

## 📊 Enhanced Reports

### 🎨 `report-enhanced` - Production Dashboard

Generate **production-ready HTML dashboards** with enterprise features.

**Syntax:**

```bash
./goliteflow report-enhanced --output=<file> [options]
```

**Enterprise Features:**

- 📈 **Real-time Statistics** - Success rates, performance metrics
- 🔄 **Automatic Rotation** - Limits to recent executions (configurable)
- 📦 **Monthly Archival** - Historical data organized by month
- 🧹 **Auto Cleanup** - Configurable retention policies
- ⚡ **Fast Performance** - Constant loading time regardless of history
- 📱 **Responsive Design** - Works on desktop and mobile

**Options:**
| Option | Description | Default |
|--------|-------------|---------|
| `--output`, `-o` | HTML report output file | `report.html` |
| `--max-executions` | Max executions in main report | `50` |
| `--archive-after` | Archive reports after N days | `30` |
| `--cleanup-after` | Delete archives after N days | `90` |
| `--page-size` | Executions per page | `20` |
| `--pagination` | Enable pagination | `true` |
| `--report-dir` | Reports directory | `reports/` |
| `--archive-dir` | Archive directory | `reports/archive/` |

**Examples:**

```bash
# Generate enhanced report (recommended)
./goliteflow report-enhanced --output=dashboard.html

# High-volume production setup
./goliteflow report-enhanced \
  --max-executions=25 \
  --archive-after=7 \
  --cleanup-after=30 \
  --output=production-dashboard.html

# Medium-volume setup
./goliteflow report-enhanced \
  --max-executions=100 \
  --page-size=25 \
  --output=dashboard.html

# Custom directories
./goliteflow report-enhanced \
  --report-dir="custom/reports" \
  --archive-dir="custom/archive" \
  --output=custom-report.html
```

### 📋 `report` - Basic Report (Legacy)

Generate basic HTML report for compatibility.

**Syntax:**

```bash
./goliteflow report --output=<file>
```

**Examples:**

```bash
# Basic report (simple, no management features)
./goliteflow report --output=simple-report.html
```

**💡 Recommendation:** Use `report-enhanced` for production deployments.

## 🛠️ Report Management

### 📊 `report-manage stats` - View Statistics

Display comprehensive report and execution statistics.

**Syntax:**

```bash
./goliteflow report-manage stats [options]
```

**Options:**
| Option | Description | Default |
|--------|-------------|---------|
| `--report-dir` | Reports directory | `reports/` |
| `--archive-dir` | Archive directory | `reports/archive/` |

**Examples:**

```bash
# View statistics
./goliteflow report-manage stats

# Custom directories
./goliteflow report-manage stats \
  --report-dir="custom/reports" \
  --archive-dir="custom/archive"
```

**Sample Output:**

```
📊 GoliteFlow Report Statistics
================================
Total Executions: 1,250
Completed: 1,180
Failed: 70
Recent (7 days): 45
Success Rate: 94.4%

📁 Storage Information
Report Directory: reports
Archive Directory: reports/archive
Archive Size: 125 MB
```

### 📦 `report-manage archive` - Archive Reports

Manually archive old reports based on age threshold.

**Syntax:**

```bash
./goliteflow report-manage archive [options]
```

**Options:**
| Option | Description | Default |
|--------|-------------|---------|
| `--days` | Archive reports older than N days | `30` |
| `--report-dir` | Reports directory | `reports/` |
| `--archive-dir` | Archive directory | `reports/archive/` |

**Examples:**

```bash
# Archive reports older than 30 days (default)
./goliteflow report-manage archive

# Archive after 15 days
./goliteflow report-manage archive --days=15

# Custom directories
./goliteflow report-manage archive \
  --days=7 \
  --report-dir="custom/reports" \
  --archive-dir="custom/archive"
```

### 🧹 `report-manage cleanup` - Clean Archives

Remove very old archived reports to free disk space.

**Syntax:**

```bash
./goliteflow report-manage cleanup [options]
```

**Options:**
| Option | Description | Default |
|--------|-------------|---------|
| `--days` | Delete archives older than N days | `90` |
| `--archive-dir` | Archive directory | `reports/archive/` |

**Examples:**

```bash
# Cleanup archives older than 90 days (default)
./goliteflow report-manage cleanup

# Cleanup after 60 days
./goliteflow report-manage cleanup --days=60

# Custom archive directory
./goliteflow report-manage cleanup \
  --days=30 \
  --archive-dir="custom/archive"
```

## 💼 Real Examples

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

## � Configuration

### 🌍 Global Options

Available for all commands:

| Option      | Short | Description             | Default              |
| ----------- | ----- | ----------------------- | -------------------- |
| `--config`  | `-c`  | YAML configuration file | `lite-workflows.yml` |
| `--verbose` | `-v`  | Enable debug logging    | `false`              |
| `--help`    | `-h`  | Show command help       | -                    |

### 📅 Schedule Patterns

| Expression    | Description           | Use Case           |
| ------------- | --------------------- | ------------------ |
| `@manual`     | Manual execution only | Testing, debugging |
| `@daily`      | Daily at midnight     | Reports, backups   |
| `@hourly`     | Every hour            | Log rotation       |
| `*/5 * * * *` | Every 5 minutes       | API monitoring     |
| `0 2 * * *`   | Daily at 2 AM         | ETL jobs           |
| `0 9 * * 1`   | Mondays at 9 AM       | Weekly reports     |

### 🔧 Production Tips

**Daemon Management:**

```bash
# Linux/macOS - Create systemd service
sudo systemctl enable goliteflow
sudo systemctl start goliteflow

# Windows - Install as service
.\goliteflow.exe daemon --install --config=production.yml

# Docker - Production container
docker run -d --name goliteflow \
  -v $(pwd):/workspace \
  sintakaridina/goliteflow:latest \
  daemon --config=/workspace/production.yml
```

**Report Management:**

```bash
# Automated maintenance (cron/task scheduler)
# Daily archival
0 2 * * * /usr/local/bin/goliteflow report-manage archive

# Weekly cleanup
0 3 * * 0 /usr/local/bin/goliteflow report-manage cleanup
```

## 📊 Exit Codes

| Code | Status              | Description                           |
| ---- | ------------------- | ------------------------------------- |
| `0`  | ✅ Success          | All operations completed successfully |
| `1`  | ❌ Error            | General execution error               |
| `2`  | ⚠️ Config Error     | YAML configuration invalid            |
| `3`  | 🔧 Workflow Error   | Workflow execution failed             |
| `4`  | 📝 Validation Error | Pre-execution validation failed       |

**Usage in Scripts:**

```bash
# Check validation before running
if ./goliteflow validate --config=production.yml; then
    ./goliteflow daemon --config=production.yml
else
    echo "❌ Configuration invalid!"
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
