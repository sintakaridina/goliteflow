---
layout: default
title: Examples
description: Real-world examples and use cases for GoliteFlow
---

# Examples

This section provides real-world examples and use cases for GoliteFlow.

## 📋 Table of Contents

- [Basic Examples](#basic-examples)
- [Data Processing](#data-processing)
- [Backup & Maintenance](#backup--maintenance)
- [CI/CD Workflows](#cicd-workflows)
- [Monitoring & Alerting](#monitoring--alerting)
- [API Integration](#api-integration)

## 🚀 Basic Examples

### Simple Hello World

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

### File Operations

```yaml
version: "1.0"
workflows:
  - name: file_operations
    schedule: "0 */6 * * *"  # Every 6 hours
    tasks:
      - id: create_backup_dir
        command: "mkdir -p /backups/$(date +%Y%m%d)"
        retry: 1
      - id: backup_important_files
        depends_on: ["create_backup_dir"]
        command: "cp -r /important /backups/$(date +%Y%m%d)/"
        retry: 2
        timeout: "30m"
      - id: compress_backup
        depends_on: ["backup_important_files"]
        command: "tar -czf /backups/$(date +%Y%m%d).tar.gz /backups/$(date +%Y%m%d)"
        retry: 1
        timeout: "1h"
```

## 📊 Data Processing

### ETL Pipeline

```yaml
version: "1.0"
workflows:
  - name: etl_pipeline
    schedule: "0 2 * * *"  # Daily at 2 AM
    tasks:
      - id: extract_data
        command: "python extract.py --source=api --output=raw_data.json"
        retry: 3
        timeout: "1h"
      
      - id: validate_data
        depends_on: ["extract_data"]
        command: "python validate.py --input=raw_data.json --schema=schema.json"
        retry: 2
        timeout: "30m"
      
      - id: transform_data
        depends_on: ["validate_data"]
        command: "python transform.py --input=raw_data.json --output=processed_data.csv"
        retry: 2
        timeout: "2h"
      
      - id: load_data
        depends_on: ["transform_data"]
        command: "python load.py --input=processed_data.csv --database=production"
        retry: 3
        timeout: "1h"
      
      - id: send_completion_notification
        depends_on: ["load_data"]
        command: "curl -X POST https://hooks.slack.com/services/... -d 'ETL completed successfully'"
        retry: 1
        timeout: "30s"
```

### Data Backup and Archival

```yaml
version: "1.0"
workflows:
  - name: data_backup
    schedule: "0 1 * * *"  # Daily at 1 AM
    tasks:
      - id: backup_database
        command: "pg_dump mydb > backup_$(date +%Y%m%d).sql"
        retry: 2
        timeout: "2h"
      
      - id: backup_files
        command: "tar -czf files_$(date +%Y%m%d).tar.gz /data"
        retry: 2
        timeout: "3h"
      
      - id: upload_to_s3
        depends_on: ["backup_database", "backup_files"]
        command: "aws s3 cp backup_$(date +%Y%m%d).sql s3://backups/database/"
        retry: 3
        timeout: "1h"
      
      - id: upload_files_to_s3
        depends_on: ["backup_database", "backup_files"]
        command: "aws s3 cp files_$(date +%Y%m%d).tar.gz s3://backups/files/"
        retry: 3
        timeout: "2h"
      
      - id: cleanup_old_backups
        depends_on: ["upload_to_s3", "upload_files_to_s3"]
        command: "find /backups -name '*.sql' -mtime +7 -delete"
        retry: 1
        timeout: "10m"
```

## 🔧 Backup & Maintenance

### System Maintenance

```yaml
version: "1.0"
workflows:
  - name: system_maintenance
    schedule: "0 3 * * 0"  # Weekly on Sunday at 3 AM
    tasks:
      - id: update_packages
        command: "apt update && apt upgrade -y"
        retry: 2
        timeout: "1h"
      
      - id: cleanup_logs
        command: "find /var/log -name '*.log' -mtime +30 -delete"
        retry: 1
        timeout: "10m"
      
      - id: cleanup_temp_files
        command: "find /tmp -type f -mtime +7 -delete"
        retry: 1
        timeout: "5m"
      
      - id: defragment_disk
        depends_on: ["cleanup_logs", "cleanup_temp_files"]
        command: "fstrim -av"
        retry: 1
        timeout: "30m"
      
      - id: restart_services
        depends_on: ["defragment_disk"]
        command: "systemctl restart nginx postgresql"
        retry: 2
        timeout: "5m"
```

### Database Maintenance

```yaml
version: "1.0"
workflows:
  - name: database_maintenance
    schedule: "0 4 * * 0"  # Weekly on Sunday at 4 AM
    tasks:
      - id: backup_database
        command: "pg_dump mydb > weekly_backup_$(date +%Y%m%d).sql"
        retry: 2
        timeout: "2h"
      
      - id: vacuum_database
        depends_on: ["backup_database"]
        command: "psql mydb -c 'VACUUM ANALYZE;'"
        retry: 2
        timeout: "1h"
      
      - id: reindex_database
        depends_on: ["vacuum_database"]
        command: "psql mydb -c 'REINDEX DATABASE mydb;'"
        retry: 2
        timeout: "3h"
      
      - id: update_statistics
        depends_on: ["reindex_database"]
        command: "psql mydb -c 'ANALYZE;'"
        retry: 1
        timeout: "30m"
```

## 🔄 CI/CD Workflows

### Automated Testing

```yaml
version: "1.0"
workflows:
  - name: automated_testing
    schedule: "0 */4 * * *"  # Every 4 hours
    tasks:
      - id: pull_latest_code
        command: "git pull origin main"
        retry: 2
        timeout: "5m"
      
      - id: run_unit_tests
        depends_on: ["pull_latest_code"]
        command: "go test ./..."
        retry: 1
        timeout: "30m"
      
      - id: run_integration_tests
        depends_on: ["run_unit_tests"]
        command: "docker-compose -f docker-compose.test.yml up --abort-on-container-exit"
        retry: 1
        timeout: "1h"
      
      - id: generate_coverage_report
        depends_on: ["run_integration_tests"]
        command: "go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out -o coverage.html"
        retry: 1
        timeout: "15m"
      
      - id: upload_artifacts
        depends_on: ["generate_coverage_report"]
        command: "aws s3 cp coverage.html s3://ci-artifacts/coverage-$(date +%Y%m%d-%H%M).html"
        retry: 2
        timeout: "10m"
```

### Deployment Pipeline

```yaml
version: "1.0"
workflows:
  - name: deployment_pipeline
    schedule: "0 18 * * 1-5"  # Weekdays at 6 PM
    tasks:
      - id: build_application
        command: "docker build -t myapp:latest ."
        retry: 2
        timeout: "30m"
      
      - id: run_tests
        depends_on: ["build_application"]
        command: "docker run --rm myapp:latest go test ./..."
        retry: 1
        timeout: "15m"
      
      - id: push_to_registry
        depends_on: ["run_tests"]
        command: "docker push myapp:latest"
        retry: 2
        timeout: "20m"
      
      - id: deploy_to_staging
        depends_on: ["push_to_registry"]
        command: "kubectl set image deployment/myapp myapp=myapp:latest -n staging"
        retry: 2
        timeout: "10m"
      
      - id: run_smoke_tests
        depends_on: ["deploy_to_staging"]
        command: "python smoke_tests.py --env=staging"
        retry: 2
        timeout: "15m"
      
      - id: deploy_to_production
        depends_on: ["run_smoke_tests"]
        command: "kubectl set image deployment/myapp myapp=myapp:latest -n production"
        retry: 2
        timeout: "10m"
```

## 📡 Monitoring & Alerting

### Health Checks

```yaml
version: "1.0"
workflows:
  - name: health_checks
    schedule: "*/5 * * * *"  # Every 5 minutes
    tasks:
      - id: check_web_service
        command: "curl -f http://localhost:8080/health"
        retry: 2
        timeout: "30s"
      
      - id: check_database
        command: "psql mydb -c 'SELECT 1;'"
        retry: 2
        timeout: "30s"
      
      - id: check_disk_space
        command: "df -h | awk '$5 > 80 {print $0}'"
        retry: 1
        timeout: "10s"
      
      - id: check_memory_usage
        command: "free | awk 'NR==2{printf \"%.2f%%\", $3*100/$2 }'"
        retry: 1
        timeout: "10s"
      
      - id: send_alert_if_failed
        depends_on: ["check_web_service", "check_database", "check_disk_space", "check_memory_usage"]
        command: "if [ $? -ne 0 ]; then curl -X POST https://hooks.slack.com/services/... -d 'Health check failed'; fi"
        retry: 1
        timeout: "30s"
```

### Performance Monitoring

```yaml
version: "1.0"
workflows:
  - name: performance_monitoring
    schedule: "0 */1 * * *"  # Every hour
    tasks:
      - id: collect_metrics
        command: "python collect_metrics.py --output=metrics.json"
        retry: 2
        timeout: "5m"
      
      - id: analyze_performance
        depends_on: ["collect_metrics"]
        command: "python analyze_performance.py --input=metrics.json"
        retry: 1
        timeout: "10m"
      
      - id: generate_report
        depends_on: ["analyze_performance"]
        command: "python generate_report.py --output=performance_report.html"
        retry: 1
        timeout: "5m"
      
      - id: send_report
        depends_on: ["generate_report"]
        command: "mail -s 'Performance Report' admin@company.com < performance_report.html"
        retry: 1
        timeout: "1m"
```

## 🔌 API Integration

### Third-party API Integration

```yaml
version: "1.0"
workflows:
  - name: api_integration
    schedule: "0 */2 * * *"  # Every 2 hours
    tasks:
      - id: fetch_weather_data
        command: "curl -s 'https://api.openweathermap.org/data/2.5/weather?q=London&appid=YOUR_API_KEY' > weather.json"
        retry: 3
        timeout: "30s"
      
      - id: process_weather_data
        depends_on: ["fetch_weather_data"]
        command: "python process_weather.py --input=weather.json --output=processed_weather.csv"
        retry: 2
        timeout: "5m"
      
      - id: store_in_database
        depends_on: ["process_weather_data"]
        command: "python store_weather.py --input=processed_weather.csv"
        retry: 2
        timeout: "2m"
      
      - id: send_notification
        depends_on: ["store_in_database"]
        command: "python send_notification.py --message='Weather data updated successfully'"
        retry: 1
        timeout: "30s"
```

### Webhook Processing

```yaml
version: "1.0"
workflows:
  - name: webhook_processing
    schedule: "*/10 * * * *"  # Every 10 minutes
    tasks:
      - id: check_webhook_queue
        command: "python check_webhook_queue.py --queue=incoming_webhooks"
        retry: 2
        timeout: "1m"
      
      - id: process_webhooks
        depends_on: ["check_webhook_queue"]
        command: "python process_webhooks.py --batch-size=100"
        retry: 2
        timeout: "10m"
      
      - id: update_status
        depends_on: ["process_webhooks"]
        command: "python update_webhook_status.py --status=processed"
        retry: 1
        timeout: "2m"
      
      - id: cleanup_processed
        depends_on: ["update_status"]
        command: "python cleanup_processed_webhooks.py --older-than=1h"
        retry: 1
        timeout: "5m"
```

## 🎯 Best Practices

### 1. Error Handling

```yaml
version: "1.0"
workflows:
  - name: robust_workflow
    schedule: "0 2 * * *"
    tasks:
      - id: critical_task
        command: "python critical_process.py"
        retry: 5  # High retry count for critical tasks
        timeout: "2h"
      
      - id: fallback_task
        depends_on: ["critical_task"]
        command: "python fallback_process.py"
        retry: 1
        timeout: "30m"
      
      - id: notification_task
        depends_on: ["critical_task", "fallback_task"]
        command: "python send_notification.py"
        retry: 3
        timeout: "1m"
```

### 2. Resource Management

```yaml
version: "1.0"
workflows:
  - name: resource_aware_workflow
    schedule: "0 3 * * *"
    tasks:
      - id: check_resources
        command: "python check_system_resources.py"
        retry: 1
        timeout: "1m"
      
      - id: adjust_workload
        depends_on: ["check_resources"]
        command: "python adjust_workload.py --based-on=resources"
        retry: 1
        timeout: "5m"
      
      - id: execute_workload
        depends_on: ["adjust_workload"]
        command: "python execute_workload.py"
        retry: 2
        timeout: "1h"
```

### 3. Monitoring and Logging

```yaml
version: "1.0"
workflows:
  - name: monitored_workflow
    schedule: "0 4 * * *"
    tasks:
      - id: start_logging
        command: "echo 'Workflow started at $(date)' >> workflow.log"
        retry: 1
        timeout: "10s"
      
      - id: main_process
        depends_on: ["start_logging"]
        command: "python main_process.py 2>&1 | tee -a workflow.log"
        retry: 2
        timeout: "2h"
      
      - id: log_completion
        depends_on: ["main_process"]
        command: "echo 'Workflow completed at $(date)' >> workflow.log"
        retry: 1
        timeout: "10s"
```

---

<div class="next-steps">
  <h3>Ready to build your workflows?</h3>
  <p>Start with the getting started guide or explore the configuration reference.</p>
  <a href="/getting-started" class="btn btn-primary">Getting Started</a>
  <a href="/configuration" class="btn btn-secondary">Configuration Reference</a>
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
