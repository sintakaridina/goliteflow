---
layout: default
title: GoliteFlow - Language Agnostic Workflow Orchestrator
description: A lightweight workflow scheduler and task orchestrator for ANY programming language. Execute tasks/workflows defined in YAML files with retry logic, conditional execution, enhanced reports, and production-ready monitoring.
keywords: workflow, scheduler, goliteflow, python, nodejs, php, java, ruby, task orchestrator, cron, yaml, automation, lightweight, monitoring, reports
author: GoliteFlow Team
---

<div class="hero">
  <h1>🚀 GoliteFlow</h1>
  <p class="lead">Lightweight workflow orchestrator for <strong>any programming language</strong></p>
  <p class="hero-subtitle">Python • Node.js • PHP • Java • Ruby • Go • Shell Commands</p>
  
  <div class="hero-buttons">
    <a href="/goliteflow/getting-started" class="btn btn-primary">🏃‍♂️ Quick Start (5 min)</a>
    <a href="https://github.com/sintakaridina/goliteflow" class="btn btn-secondary">📂 View on GitHub</a>
  </div>
  
  <div class="hero-stats">
    <span class="stat"><strong>Zero Dependencies</strong> • Single Binary</span>
    <span class="stat"><strong>Production Ready</strong> • Enterprise Reports</span>
    <span class="stat"><strong>Cross Platform</strong> • Linux, Windows, macOS</span>
  </div>
</div></div>

## ✨ Key Features

<div class="features-grid">
  <div class="feature-card">
    <div class="feature-icon">🌍</div>
    <h3>Language Agnostic</h3>
    <p>Works with Python, Node.js, PHP, Java, Ruby, Go, or any shell command. No code changes required.</p>
  </div>
  
  <div class="feature-card">
    <div class="feature-icon">📊</div>
    <h3>Enhanced Reports</h3>
    <p>Production-ready HTML dashboards with automatic archival, pagination, and enterprise scaling features.</p>
  </div>
  
  <div class="feature-card">
    <div class="feature-icon">⏰</div>
    <h3>Smart Scheduling</h3>
    <p>Cron-based scheduling with dependency management, retry logic, and conditional execution.</p>
  </div>
  
  <div class="feature-card">
    <div class="feature-icon">🚀</div>
    <h3>Zero Dependencies</h3>
    <p>Single binary deployment. No databases, web servers, or complex setup required.</p>
  </div>
  
  <div class="feature-card">
    <div class="feature-icon">🔄</div>
    <h3>Daemon Mode</h3>
    <p>Continuous monitoring with automatic HTML report updates after each execution.</p>
  </div>
  
  <div class="feature-card">
    <div class="feature-icon">🛠️</div>
    <h3>Production Ready</h3>
    <p>Report management, archival, cleanup, and enterprise features for long-running deployments.</p>
  </div>
</div>

## 🚀 Quick Start (5 Minutes)

### Step 1: Download Binary

Choose your platform:

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

Create `my-workflow.yml`:

```yaml
version: "1.0"
workflows:
  - name: hello_world_demo
    schedule: "@manual" # Run manually for testing
    tasks:
      - id: greet
        command: "echo Hello from GoliteFlow!"

      - id: check_system
        command: "python --version && node --version"
        depends_on: ["greet"]

      - id: list_files
        command: "ls -la" # or "dir" on Windows
        depends_on: ["check_system"]
```

### Step 3: Run & Monitor

```bash
# Run the workflow
./goliteflow run --config=my-workflow.yml

# Generate enhanced HTML report
./goliteflow report-enhanced --output=dashboard.html

# Start continuous daemon (production)
./goliteflow daemon --config=my-workflow.yml
```

### Real-World Examples

**Python Data Pipeline:**

```yaml
version: "1.0"
workflows:
  - name: daily_etl
    schedule: "0 2 * * *" # Every day at 2 AM
    tasks:
      - id: extract_data
        command: "python scripts/extract_from_api.py"
        retry_count: 3
      - id: process_data
        command: "python scripts/transform_data.py"
        depends_on: ["extract_data"]
      - id: send_report
        command: "python scripts/email_summary.py"
        depends_on: ["process_data"]
```

**Node.js API Monitoring:**

```yaml
version: "1.0"
workflows:
  - name: api_health_check
    schedule: "*/5 * * * *" # Every 5 minutes
    tasks:
      - id: health_check
        command: "node monitoring/health-check.js"
      - id: alert_on_failure
        command: "node monitoring/send-alert.js"
        condition: "on_failure"
```

## 🎯 Why Choose GoliteFlow?

<div class="comparison">
  <div class="pros">
    <h3>✅ Key Advantages</h3>
    <ul>
      <li><strong>🌍 Any Language</strong> - Python, Node.js, PHP, Java, Ruby, Go, shell commands</li>
      <li><strong>📦 Zero Setup</strong> - Single binary, no databases or web servers</li>
      <li><strong>⚡ Lightning Fast</strong> - 5-minute setup from zero to production</li>
      <li><strong>📊 Enterprise Reports</strong> - Automatic archival and scaling</li>
      <li><strong>🔒 Production Ready</strong> - Built-in reliability and monitoring</li>
      <li><strong>🛠️ Developer Friendly</strong> - YAML config, clear documentation</li>
    </ul>
  </div>
  
  <div class="use-cases">
    <h3>🎨 Perfect Use Cases</h3>
    <ul>
      <li>🐍 Python data processing & ML pipelines</li>
      <li>🟢 Node.js API monitoring & automation</li>
      <li>🐘 PHP application maintenance tasks</li>
      <li>☕ Java batch processing & reports</li>
      <li>💎 Ruby deployment & backup scripts</li>
      <li>🔧 DevOps automation & CI/CD</li>
      <li>📈 ETL pipelines & data workflows</li>
      <li>🚨 Health checks & alerting systems</li>
    </ul>
  </div>
</div>

## 📊 vs Other Solutions

| Feature               | GoliteFlow   | Airflow           | Prefect     | Temporal          |
| --------------------- | ------------ | ----------------- | ----------- | ----------------- |
| **Setup Time**        | ⚡ 5 minutes | 🕒 Hours          | 🕐 30+ min  | 🕒 Hours          |
| **Language Support**  | 🌍 Any       | 🐍 Python         | 🐍 Python   | 📦 SDK Required   |
| **Dependencies**      | ✅ None      | 🗄️ Database + Web | 🗄️ Database | 🗄️ Database + Web |
| **Resource Usage**    | 🪶 Minimal   | 🏗️ Heavy          | 🏢 Medium   | 🏗️ Heavy          |
| **HTML Reports**      | 📊 Built-in  | ❌ External       | ❌ External | ❌ External       |
| **Beginner Friendly** | ✅ Very Easy | ❌ Complex        | 🟡 Medium   | ❌ Complex        |

## 📊 Enhanced HTML Reports

GoliteFlow generates **production-ready HTML dashboards** with enterprise features:

### 🎨 Dashboard Features

- **📈 Real-time Statistics** - Success rates, execution trends, performance metrics
- **🔍 Interactive Timeline** - Visual workflow execution history
- **📋 Task Details** - Individual results, retry attempts, error logs
- **🎯 Dependency Graph** - Visual task dependency mapping
- **📱 Responsive Design** - Works on desktop and mobile

### 🚀 Enterprise Management

- **🔄 Automatic Rotation** - Limits main report to recent executions
- **📦 Monthly Archival** - Historical data organized by month
- **🧹 Auto Cleanup** - Configurable retention policies
- **⚡ Fast Loading** - Constant performance regardless of history size
- **📊 Analytics Dashboard** - Comprehensive workflow analytics

### 🛠️ Report Commands

```bash
# Generate enhanced report (recommended)
./goliteflow report-enhanced --output=dashboard.html

# Configure report management
./goliteflow report-enhanced --max-executions=100 --archive-after=30

# Manage archives
./goliteflow report-manage stats      # View statistics
./goliteflow report-manage cleanup    # Clean old archives
```

[📸 View Live Demo Report →](https://sintakaridina.github.io/goliteflow/examples/complete-report.html)

## 🚀 Production Deployment

<div class="deployment-options">
  <div class="deployment-card">
    <h3>🐧 Linux/macOS (systemd)</h3>
    <pre><code># Install as system service
sudo systemctl enable goliteflow
sudo systemctl start goliteflow
    </code></pre>
  </div>
  
  <div class="deployment-card">
    <h3>🪟 Windows Service</h3>
    <pre><code># Install as Windows service
.\goliteflow.exe daemon --install --config=production.yml
net start goliteflow
    </code></pre>
  </div>
  
  <div class="deployment-card">
    <h3>🐳 Docker Container</h3>
    <pre><code># Run in container
docker run -v $(pwd):/workflows \
  sintakaridina/goliteflow:latest \
  daemon --config=/workflows/config.yml
    </code></pre>
  </div>
</div>

### 📊 Monitoring & Scaling

- **📈 Real-time Dashboard** - Live HTML reports with auto-refresh
- **🔄 Report Rotation** - Automatic archival prevents size growth
- **📦 Archive Management** - Monthly organization with configurable cleanup
- **🚨 Health Monitoring** - Built-in status checks and alerting
- **⚡ Performance** - Handles thousands of executions efficiently

<div class="badges">
  <img src="https://img.shields.io/badge/Go-1.19%2B-blue" alt="Go Version">
  <img src="https://img.shields.io/badge/License-MIT-green" alt="License">
  <img src="https://img.shields.io/badge/Status-Production%20Ready-brightgreen" alt="Status">
  <img src="https://img.shields.io/github/stars/sintakaridina/goliteflow?style=social" alt="GitHub Stars">
</div>

## Community

<div class="community">
  <div class="community-item">
    <h4>Found a Bug?</h4>
    <p>Report issues on <a href="https://github.com/sintakaridina/goliteflow/issues">GitHub Issues</a></p>
  </div>
  
  <div class="community-item">
    <h4>Have an Idea?</h4>
    <p>Suggest features in <a href="https://github.com/sintakaridina/goliteflow/discussions">GitHub Discussions</a></p>
  </div>
  
  <div class="community-item">
    <h4>Want to Contribute?</h4>
    <p>Check out our <a href="/contributing">Contributing Guide</a></p>
  </div>
</div>

## 📖 Documentation

<div class="docs-grid">
  <div class="doc-card">
    <h4>🏃‍♂️ <a href="/goliteflow/getting-started">Getting Started</a></h4>
    <p>5-minute setup guide for beginners</p>
  </div>
  
  <div class="doc-card">
    <h4>⚙️ <a href="/goliteflow/configuration">Configuration Reference</a></h4>
    <p>Complete YAML configuration options</p>
  </div>
  
  <div class="doc-card">
    <h4>🖥️ <a href="/goliteflow/cli-reference">CLI Reference</a></h4>
    <p>All command-line options and examples</p>
  </div>
  
  <div class="doc-card">
    <h4>📊 <a href="/goliteflow/report-management">Report Management</a></h4>
    <p>Enterprise report features and scaling</p>
  </div>
  
  <div class="doc-card">
    <h4>💼 <a href="https://github.com/sintakaridina/goliteflow/tree/main/examples">Real Examples</a></h4>
    <p>Python, Node.js, DevOps workflows</p>
  </div>
  
  <div class="doc-card">
    <h4>🤝 <a href="/goliteflow/contributing">Contributing</a></h4>
    <p>Help improve GoliteFlow</p>
  </div>
</div>

---

<div class="footer-cta">
  <h2>🎉 Ready to Automate Your Workflows?</h2>
  <p>Join developers using GoliteFlow for Python, Node.js, PHP, and more!</p>
  <div class="cta-buttons">
    <a href="/goliteflow/getting-started" class="btn btn-primary btn-large">🚀 Start in 5 Minutes</a>
    <a href="https://github.com/sintakaridina/goliteflow/releases" class="btn btn-secondary btn-large">📦 Download Binary</a>
  </div>
  
  <div class="stats-footer">
    <span>⭐ Zero Dependencies</span>
    <span>🌍 Any Language</span>
    <span>📊 Enhanced Reports</span>
    <span>🚀 Production Ready</span>
  </div>
</div>
