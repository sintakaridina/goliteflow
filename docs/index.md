---
layout: default
title: GoliteFlow
description: A lightweight workflow scheduler and task orchestrator for Go applications
---

<div class="hero">
  <div class="hero-content">
    <h1>GoliteFlow</h1>
    <p class="lead">A lightweight workflow scheduler and task orchestrator for Go applications</p>
    
    <div class="hero-buttons">
      <a href="/getting-started" class="btn btn-primary">Get Started</a>
      <a href="https://github.com/sintakaridina/goliteflow" class="btn btn-secondary">View on GitHub</a>
    </div>
  </div>
</div>

## Features

<div class="features-grid">
  <div class="feature-card">
    <div class="feature-icon">
      <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
        <polyline points="14,2 14,8 20,8"></polyline>
        <line x1="16" y1="13" x2="8" y2="13"></line>
        <line x1="16" y1="17" x2="8" y2="17"></line>
        <polyline points="10,9 9,9 8,9"></polyline>
      </svg>
    </div>
    <h3>YAML Configuration</h3>
    <p>Define workflows and tasks in simple, human-readable YAML files with dependency management.</p>
  </div>
  
  <div class="feature-card">
    <div class="feature-icon">
      <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <circle cx="12" cy="12" r="10"></circle>
        <polyline points="12,6 12,12 16,14"></polyline>
      </svg>
    </div>
    <h3>Cron Scheduling</h3>
    <p>Built-in scheduler using standard cron syntax for flexible task scheduling.</p>
  </div>
  
  <div class="feature-card">
    <div class="feature-icon">
      <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <path d="M1 4v6h6"></path>
        <path d="M3.51 15a9 9 0 1 0 2.13-9.36L1 10"></path>
      </svg>
    </div>
    <h3>Retry Logic</h3>
    <p>Configurable retry mechanisms with exponential backoff for robust task execution.</p>
  </div>
  
  <div class="feature-card">
    <div class="feature-icon">
      <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <path d="M3 3v18h18"></path>
        <path d="M18.7 8l-5.1 5.2-2.8-2.7L7 14.3"></path>
      </svg>
    </div>
    <h3>HTML Reports</h3>
    <p>Beautiful, interactive HTML reports with execution history and detailed task logs.</p>
  </div>
  
  <div class="feature-card">
    <div class="feature-icon">
      <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <rect x="2" y="3" width="20" height="14" rx="2" ry="2"></rect>
        <line x1="8" y1="21" x2="16" y2="21"></line>
        <line x1="12" y1="17" x2="12" y2="21"></line>
      </svg>
    </div>
    <h3>CLI Tool</h3>
    <p>Easy-to-use command-line interface for running and managing workflows.</p>
  </div>
  
  <div class="feature-card">
    <div class="feature-icon">
      <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"></path>
      </svg>
    </div>
    <h3>Go Library</h3>
    <p>Use as a Go library in your applications with a clean, simple API.</p>
  </div>
</div>

## Quick Start

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

## Why GoliteFlow?

<div class="comparison">
  <div class="pros">
    <h3>Advantages</h3>
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
    <h3>Perfect For</h3>
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

## Example Report

GoliteFlow generates beautiful HTML reports with:

- **Workflow Statistics**: Success rates, execution times
- **Task Details**: Individual task results, retry counts
- **Interactive Interface**: Expandable sections, search functionality
- **Error Logs**: Detailed stdout/stderr capture
- **Timeline View**: Execution history with timestamps

[View Sample Report](examples/sample-report.html)

## Production Ready

<div class="badges">
  <img src="https://img.shields.io/badge/Go-1.19%2B-blue" alt="Go Version">
  <img src="https://img.shields.io/badge/License-MIT-green" alt="License">
  <img src="https://img.shields.io/badge/Status-Production%20Ready-brightgreen" alt="Status">
  <img src="https://img.shields.io/github/stars/sintakaridina/goliteflow?style=social" alt="GitHub Stars">
</div>

- **Comprehensive Testing** - Full test coverage
- **Error Handling** - Robust error handling and validation
- **Documentation** - Complete documentation and examples
- **CI/CD** - Automated testing and building
- **Docker Support** - Container-ready
- **Cross-platform** - Windows, Linux, macOS support

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

## Documentation

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
/* Reset and base styles */
* {
  box-sizing: border-box;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
  line-height: 1.6;
  color: #333;
  margin: 0;
  padding: 0;
}

/* Hero section */
.hero {
  background: #ffffff;
  padding: 4rem 0;
  text-align: center;
  border-bottom: 1px solid #e1e5e9;
}

.hero-content {
  max-width: 800px;
  margin: 0 auto;
  padding: 0 2rem;
}

.hero h1 {
  font-size: 3.5rem;
  font-weight: 300;
  margin: 0 0 1rem 0;
  color: #2c3e50;
  letter-spacing: -0.02em;
}

.hero .lead {
  font-size: 1.25rem;
  color: #6c757d;
  margin-bottom: 2.5rem;
  font-weight: 400;
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
  border-radius: 6px;
  text-decoration: none;
  font-weight: 500;
  transition: all 0.2s ease;
  border: 1px solid transparent;
  font-size: 0.95rem;
}

.btn-primary {
  background: #007bff;
  color: white;
  border-color: #007bff;
}

.btn-primary:hover {
  background: #0056b3;
  border-color: #0056b3;
}

.btn-secondary {
  background: transparent;
  color: #007bff;
  border-color: #007bff;
}

.btn-secondary:hover {
  background: #007bff;
  color: white;
}

/* Features grid */
.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
  gap: 2rem;
  margin: 4rem 0;
  max-width: 1200px;
  margin-left: auto;
  margin-right: auto;
  padding: 0 2rem;
}

.feature-card {
  background: white;
  border: 1px solid #e1e5e9;
  border-radius: 8px;
  padding: 2rem;
  text-align: left;
  transition: all 0.2s ease;
}

.feature-card:hover {
  border-color: #007bff;
  box-shadow: 0 4px 12px rgba(0,123,255,0.1);
}

.feature-icon {
  width: 48px;
  height: 48px;
  background: #f8f9fa;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 1.5rem;
  color: #007bff;
}

.feature-card h3 {
  color: #2c3e50;
  margin: 0 0 1rem 0;
  font-size: 1.25rem;
  font-weight: 600;
}

.feature-card p {
  color: #6c757d;
  line-height: 1.6;
  margin: 0;
}

/* Comparison section */
.comparison {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 3rem;
  margin: 4rem 0;
  max-width: 1000px;
  margin-left: auto;
  margin-right: auto;
  padding: 0 2rem;
}

.pros, .use-cases {
  background: white;
  border: 1px solid #e1e5e9;
  border-radius: 8px;
  padding: 2rem;
}

.pros h3, .use-cases h3 {
  color: #2c3e50;
  margin: 0 0 1.5rem 0;
  font-size: 1.5rem;
  font-weight: 600;
}

.pros ul, .use-cases ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.pros li, .use-cases li {
  padding: 0.75rem 0;
  border-bottom: 1px solid #f1f3f4;
  color: #495057;
}

.pros li:last-child, .use-cases li:last-child {
  border-bottom: none;
}

/* Badges */
.badges {
  text-align: center;
  margin: 3rem 0;
}

.badges img {
  margin: 0 0.5rem;
  height: 24px;
}

/* Community section */
.community {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 2rem;
  margin: 4rem 0;
  max-width: 1000px;
  margin-left: auto;
  margin-right: auto;
  padding: 0 2rem;
}

.community-item {
  background: white;
  border: 1px solid #e1e5e9;
  border-radius: 8px;
  padding: 2rem;
  text-align: center;
  transition: all 0.2s ease;
}

.community-item:hover {
  border-color: #007bff;
  box-shadow: 0 4px 12px rgba(0,123,255,0.1);
}

.community-item h4 {
  color: #2c3e50;
  margin: 0 0 1rem 0;
  font-size: 1.1rem;
  font-weight: 600;
}

.community-item p {
  color: #6c757d;
  margin: 0;
}

/* Footer CTA */
.footer-cta {
  text-align: center;
  padding: 4rem 2rem;
  background: #f8f9fa;
  border-top: 1px solid #e1e5e9;
  margin: 4rem 0 0 0;
}

.footer-cta h3 {
  font-size: 2rem;
  font-weight: 300;
  margin: 0 0 1rem 0;
  color: #2c3e50;
}

.footer-cta p {
  font-size: 1.1rem;
  margin: 0 0 2rem 0;
  color: #6c757d;
}

.btn-large {
  padding: 1rem 2rem;
  font-size: 1.1rem;
}

/* Responsive design */
@media (max-width: 768px) {
  .hero h1 {
    font-size: 2.5rem;
  }
  
  .hero .lead {
    font-size: 1.1rem;
  }
  
  .hero-buttons {
    flex-direction: column;
    align-items: center;
  }
  
  .comparison {
    grid-template-columns: 1fr;
    gap: 2rem;
  }
  
  .features-grid {
    grid-template-columns: 1fr;
    padding: 0 1rem;
  }
  
  .community {
    grid-template-columns: 1fr;
    padding: 0 1rem;
  }
  
  .hero-content {
    padding: 0 1rem;
  }
  
  .comparison {
    padding: 0 1rem;
  }
}

@media (max-width: 480px) {
  .hero h1 {
    font-size: 2rem;
  }
  
  .feature-card {
    padding: 1.5rem;
  }
  
  .pros, .use-cases {
    padding: 1.5rem;
  }
}
</style>
