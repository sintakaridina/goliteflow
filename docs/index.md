---
layout: default
title: GoliteFlow
description: A lightweight workflow scheduler and task orchestrator for Go applications
---

# GoliteFlow

<div class="hero">
  <h1>🚀 Lightweight Workflow Orchestrator</h1>
  <p class="lead">Execute tasks and workflows defined in YAML files with retry logic, scheduling, and beautiful HTML reports.</p>
  
  <div class="hero-buttons">
    <a href="/getting-started" class="btn btn-primary">Get Started</a>
    <a href="https://github.com/sintakaridina/goliteflow" class="btn btn-secondary">View on GitHub</a>
  </div>
</div>

## ✨ Features

<div class="features-grid">
  <div class="feature-card">
    <h3>📝 YAML Configuration</h3>
    <p>Define workflows and tasks in simple, human-readable YAML files with dependency management.</p>
  </div>
  
  <div class="feature-card">
    <h3>⏰ Cron Scheduling</h3>
    <p>Built-in scheduler using standard cron syntax for flexible task scheduling.</p>
  </div>
  
  <div class="feature-card">
    <h3>🔄 Retry Logic</h3>
    <p>Configurable retry mechanisms with exponential backoff for robust task execution.</p>
  </div>
  
  <div class="feature-card">
    <h3>📊 HTML Reports</h3>
    <p>Beautiful, interactive HTML reports with execution history and detailed task logs.</p>
  </div>
  
  <div class="feature-card">
    <h3>🖥️ CLI Tool</h3>
    <p>Easy-to-use command-line interface for running and managing workflows.</p>
  </div>
  
  <div class="feature-card">
    <h3>📚 Go Library</h3>
    <p>Use as a Go library in your applications with a clean, simple API.</p>
  </div>
</div>

## 🚀 Quick Start

### Installation

```bash
# Install via Go
go get github.com/sintakaridina/goliteflow

# Or download binary from releases
curl -L https://github.com/sintakaridina/goliteflow/releases/latest/download/goliteflow-windows-amd64.exe -o goliteflow.exe
```

### Basic Usage

1. **Create a workflow file** (`workflows.yml`):

```yaml
version: "1.0"
workflows:
  - name: daily_backup
    schedule: "0 2 * * *"  # Daily at 2 AM
    tasks:
      - id: backup_data
        command: "tar -czf backup.tar.gz /data"
        retry: 3
      - id: upload_backup
        depends_on: ["backup_data"]
        command: "aws s3 cp backup.tar.gz s3://my-bucket/"
        retry: 2
```

2. **Run the workflow**:

```bash
# Validate configuration
goliteflow validate --config=workflows.yml

# Run once
goliteflow run --config=workflows.yml

# Run as daemon
goliteflow run --config=workflows.yml --daemon

# Generate report
goliteflow report --output=report.html
```

3. **Use as Go library**:

```go
package main

import "github.com/sintakaridina/goliteflow"

func main() {
    // Simple usage
    goliteflow.Run("workflows.yml")
    
    // With report generation
    goliteflow.RunWithReport("workflows.yml", "report.html")
}
```

## 📈 Why GoliteFlow?

<div class="comparison">
  <div class="pros">
    <h3>✅ Advantages</h3>
    <ul>
      <li><strong>Zero Dependencies</strong> - No external database or web server required</li>
      <li><strong>Lightweight</strong> - Minimal resource usage, perfect for small applications</li>
      <li><strong>Simple</strong> - Easy YAML configuration, no complex setup</li>
      <li><strong>Self-contained</strong> - HTML reports with embedded CSS/JS</li>
      <li><strong>Fast</strong> - Efficient execution with goroutines</li>
      <li><strong>Reliable</strong> - Built-in retry logic and error handling</li>
    </ul>
  </div>
  
  <div class="use-cases">
    <h3>🎯 Perfect For</h3>
    <ul>
      <li>Data processing pipelines</li>
      <li>Backup and maintenance tasks</li>
      <li>CI/CD workflows</li>
      <li>Monitoring and alerting</li>
      <li>File processing and ETL</li>
      <li>API integrations</li>
    </ul>
  </div>
</div>

## 📊 Example Report

GoliteFlow generates beautiful HTML reports with:

- **Workflow Statistics**: Success rates, execution times
- **Task Details**: Individual task results, retry counts
- **Interactive Interface**: Expandable sections, search functionality
- **Error Logs**: Detailed stdout/stderr capture
- **Timeline View**: Execution history with timestamps

[View Sample Report](examples/sample-report.html)

## 🏆 Production Ready

<div class="badges">
  <img src="https://img.shields.io/badge/Go-1.19%2B-blue" alt="Go Version">
  <img src="https://img.shields.io/badge/License-MIT-green" alt="License">
  <img src="https://img.shields.io/badge/Status-Production%20Ready-brightgreen" alt="Status">
  <img src="https://img.shields.io/github/stars/sintakaridina/goliteflow?style=social" alt="GitHub Stars">
</div>

- ✅ **Comprehensive Testing** - Full test coverage
- ✅ **Error Handling** - Robust error handling and validation
- ✅ **Documentation** - Complete documentation and examples
- ✅ **CI/CD** - Automated testing and building
- ✅ **Docker Support** - Container-ready
- ✅ **Cross-platform** - Windows, Linux, macOS support

## 🤝 Community

<div class="community">
  <div class="community-item">
    <h4>🐛 Found a Bug?</h4>
    <p>Report issues on <a href="https://github.com/sintakaridina/goliteflow/issues">GitHub Issues</a></p>
  </div>
  
  <div class="community-item">
    <h4>💡 Have an Idea?</h4>
    <p>Suggest features in <a href="https://github.com/sintakaridina/goliteflow/discussions">GitHub Discussions</a></p>
  </div>
  
  <div class="community-item">
    <h4>🔧 Want to Contribute?</h4>
    <p>Check out our <a href="/contributing">Contributing Guide</a></p>
  </div>
</div>

## 📚 Documentation

- [Getting Started](/getting-started) - Quick setup guide
- [Configuration Reference](/configuration) - YAML configuration details
- [CLI Reference](/cli-reference) - Command-line interface
- [Library API](/api) - Go library documentation
- [Examples](/examples) - Real-world use cases
- [Deployment Guide](/deployment) - Production deployment

---

<div class="footer-cta">
  <h3>Ready to get started?</h3>
  <p>GoliteFlow makes workflow orchestration simple and reliable.</p>
  <a href="/getting-started" class="btn btn-primary btn-large">Start Building</a>
</div>

<style>
.hero {
  text-align: center;
  padding: 3rem 0;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  margin: -2rem -1rem 3rem -1rem;
  border-radius: 0 0 1rem 1rem;
}

.hero h1 {
  font-size: 3rem;
  margin-bottom: 1rem;
}

.lead {
  font-size: 1.25rem;
  margin-bottom: 2rem;
  opacity: 0.9;
}

.hero-buttons {
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
}

.btn {
  display: inline-block;
  padding: 0.75rem 1.5rem;
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
  color: white;
  border: 2px solid white;
}

.btn-secondary:hover {
  background: white;
  color: #667eea;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
  margin: 3rem 0;
}

.feature-card {
  background: #f8f9fa;
  padding: 2rem;
  border-radius: 1rem;
  border: 1px solid #e9ecef;
  transition: transform 0.3s ease;
}

.feature-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 25px rgba(0,0,0,0.1);
}

.feature-card h3 {
  color: #495057;
  margin-bottom: 1rem;
}

.comparison {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 2rem;
  margin: 3rem 0;
}

.pros, .use-cases {
  background: #f8f9fa;
  padding: 2rem;
  border-radius: 1rem;
}

.pros h3, .use-cases h3 {
  color: #495057;
  margin-bottom: 1rem;
}

.badges {
  text-align: center;
  margin: 2rem 0;
}

.badges img {
  margin: 0 0.5rem;
}

.community {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 2rem;
  margin: 3rem 0;
}

.community-item {
  text-align: center;
  padding: 2rem;
  background: #f8f9fa;
  border-radius: 1rem;
}

.footer-cta {
  text-align: center;
  padding: 3rem;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border-radius: 1rem;
  margin: 3rem 0;
}

.btn-large {
  padding: 1rem 2rem;
  font-size: 1.25rem;
}

@media (max-width: 768px) {
  .hero h1 {
    font-size: 2rem;
  }
  
  .comparison {
    grid-template-columns: 1fr;
  }
  
  .hero-buttons {
    flex-direction: column;
    align-items: center;
  }
}
</style>
