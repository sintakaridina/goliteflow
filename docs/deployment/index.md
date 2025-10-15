---
layout: default
title: Deployment Guide
description: Production deployment guide for GoliteFlow
---

# Deployment Guide

This guide covers deploying GoliteFlow in production environments.

## 📋 Table of Contents

- [Docker Deployment](#docker-deployment)
- [Kubernetes Deployment](#kubernetes-deployment)
- [Systemd Service](#systemd-service)
- [Cloud Deployment](#cloud-deployment)
- [Monitoring & Logging](#monitoring--logging)
- [Security Considerations](#security-considerations)

## 🐳 Docker Deployment

### Basic Docker Run

```bash
# Run with volume mount
docker run -v $(pwd):/workspace sintakaridina/goliteflow:latest run --config=/workspace/workflows.yml

# Run as daemon
docker run -d -v $(pwd):/workspace --name goliteflow sintakaridina/goliteflow:latest run --config=/workspace/workflows.yml --daemon
```

### Docker Compose

Create `docker-compose.yml`:

```yaml
version: '3.8'

services:
  goliteflow:
    image: sintakaridina/goliteflow:latest
    container_name: goliteflow
    restart: unless-stopped
    volumes:
      - ./workflows.yml:/app/workflows.yml:ro
      - ./reports:/app/reports
      - ./logs:/app/logs
    command: ["run", "--config=/app/workflows.yml", "--daemon", "--output=/app/reports/report.html"]
    environment:
      - TZ=UTC
    healthcheck:
      test: ["CMD", "goliteflow", "validate", "--config=/app/workflows.yml"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 40s
```

Deploy:

```bash
docker-compose up -d
```

### Custom Dockerfile

```dockerfile
FROM sintakaridina/goliteflow:latest

# Copy your workflows
COPY workflows.yml /app/workflows.yml

# Create directories for reports and logs
RUN mkdir -p /app/reports /app/logs

# Set working directory
WORKDIR /app

# Run your workflows
CMD ["run", "--config=/app/workflows.yml", "--daemon", "--output=/app/reports/report.html"]
```

Build and run:

```bash
docker build -t my-goliteflow .
docker run -d --name my-goliteflow my-goliteflow
```

## ☸️ Kubernetes Deployment

### Deployment Manifest

Create `goliteflow-deployment.yaml`:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: goliteflow
  labels:
    app: goliteflow
spec:
  replicas: 1
  selector:
    matchLabels:
      app: goliteflow
  template:
    metadata:
      labels:
        app: goliteflow
    spec:
      containers:
      - name: goliteflow
        image: sintakaridina/goliteflow:latest
        command: ["goliteflow", "run", "--config=/app/workflows.yml", "--daemon"]
        volumeMounts:
        - name: workflows
          mountPath: /app/workflows.yml
          subPath: workflows.yml
          readOnly: true
        - name: reports
          mountPath: /app/reports
        - name: logs
          mountPath: /app/logs
        env:
        - name: TZ
          value: "UTC"
        resources:
          requests:
            memory: "64Mi"
            cpu: "100m"
          limits:
            memory: "256Mi"
            cpu: "500m"
        livenessProbe:
          exec:
            command:
            - goliteflow
            - validate
            - --config=/app/workflows.yml
          initialDelaySeconds: 30
          periodSeconds: 30
        readinessProbe:
          exec:
            command:
            - goliteflow
            - validate
            - --config=/app/workflows.yml
          initialDelaySeconds: 5
          periodSeconds: 10
      volumes:
      - name: workflows
        configMap:
          name: goliteflow-workflows
      - name: reports
        persistentVolumeClaim:
          claimName: goliteflow-reports
      - name: logs
        persistentVolumeClaim:
          claimName: goliteflow-logs
```

### ConfigMap

Create `goliteflow-configmap.yaml`:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: goliteflow-workflows
data:
  workflows.yml: |
    version: "1.0"
    workflows:
      - name: daily_backup
        schedule: "0 2 * * *"
        tasks:
          - id: backup_data
            command: "tar -czf backup.tar.gz /data"
            retry: 3
```

### PersistentVolumeClaim

Create `goliteflow-pvc.yaml`:

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: goliteflow-reports
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: goliteflow-logs
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
```

### Service

Create `goliteflow-service.yaml`:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: goliteflow-service
spec:
  selector:
    app: goliteflow
  ports:
  - port: 8080
    targetPort: 8080
  type: ClusterIP
```

Deploy to Kubernetes:

```bash
kubectl apply -f goliteflow-configmap.yaml
kubectl apply -f goliteflow-pvc.yaml
kubectl apply -f goliteflow-deployment.yaml
kubectl apply -f goliteflow-service.yaml
```

## 🔧 Systemd Service

### Service File

Create `/etc/systemd/system/goliteflow.service`:

```ini
[Unit]
Description=GoliteFlow Workflow Scheduler
After=network.target

[Service]
Type=simple
User=goliteflow
Group=goliteflow
WorkingDirectory=/opt/goliteflow
ExecStart=/usr/local/bin/goliteflow run --config=/opt/goliteflow/workflows.yml --daemon
ExecReload=/bin/kill -HUP $MAINPID
Restart=always
RestartSec=10
StandardOutput=journal
StandardError=journal
SyslogIdentifier=goliteflow

# Security settings
NoNewPrivileges=true
PrivateTmp=true
ProtectSystem=strict
ProtectHome=true
ReadWritePaths=/opt/goliteflow/reports /opt/goliteflow/logs

[Install]
WantedBy=multi-user.target
```

### Setup Script

Create setup script:

```bash
#!/bin/bash

# Create user and directories
sudo useradd -r -s /bin/false goliteflow
sudo mkdir -p /opt/goliteflow/{reports,logs}
sudo chown -R goliteflow:goliteflow /opt/goliteflow

# Install binary
sudo cp goliteflow /usr/local/bin/
sudo chmod +x /usr/local/bin/goliteflow

# Copy configuration
sudo cp workflows.yml /opt/goliteflow/
sudo chown goliteflow:goliteflow /opt/goliteflow/workflows.yml

# Enable and start service
sudo systemctl daemon-reload
sudo systemctl enable goliteflow
sudo systemctl start goliteflow

# Check status
sudo systemctl status goliteflow
```

## ☁️ Cloud Deployment

### AWS ECS

Create `goliteflow-task-definition.json`:

```json
{
  "family": "goliteflow",
  "networkMode": "awsvpc",
  "requiresCompatibilities": ["FARGATE"],
  "cpu": "256",
  "memory": "512",
  "executionRoleArn": "arn:aws:iam::ACCOUNT:role/ecsTaskExecutionRole",
  "taskRoleArn": "arn:aws:iam::ACCOUNT:role/ecsTaskRole",
  "containerDefinitions": [
    {
      "name": "goliteflow",
      "image": "sintakaridina/goliteflow:latest",
      "command": ["goliteflow", "run", "--config=/app/workflows.yml", "--daemon"],
      "portMappings": [],
      "essential": true,
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-group": "/ecs/goliteflow",
          "awslogs-region": "us-west-2",
          "awslogs-stream-prefix": "ecs"
        }
      },
      "environment": [
        {
          "name": "TZ",
          "value": "UTC"
        }
      ],
      "mountPoints": [
        {
          "sourceVolume": "workflows",
          "containerPath": "/app/workflows.yml",
          "readOnly": true
        }
      ]
    }
  ],
  "volumes": [
    {
      "name": "workflows",
      "efsVolumeConfiguration": {
        "fileSystemId": "fs-12345678",
        "rootDirectory": "/workflows"
      }
    }
  ]
}
```

### Google Cloud Run

Create `cloud-run.yaml`:

```yaml
apiVersion: serving.knative.dev/v1
kind: Service
metadata:
  name: goliteflow
  annotations:
    run.googleapis.com/ingress: all
spec:
  template:
    metadata:
      annotations:
        run.googleapis.com/execution-environment: gen2
    spec:
      containerConcurrency: 1
      timeoutSeconds: 3600
      containers:
      - image: sintakaridina/goliteflow:latest
        command: ["goliteflow", "run", "--config=/app/workflows.yml", "--daemon"]
        env:
        - name: TZ
          value: "UTC"
        resources:
          limits:
            cpu: "1"
            memory: "512Mi"
        volumeMounts:
        - name: workflows
          mountPath: /app/workflows.yml
          subPath: workflows.yml
          readOnly: true
      volumes:
      - name: workflows
        secret:
          secretName: goliteflow-workflows
```

## 📊 Monitoring & Logging

### Prometheus Metrics

Create custom metrics collection:

```yaml
version: "1.0"
workflows:
  - name: metrics_collection
    schedule: "*/5 * * * *"  # Every 5 minutes
    tasks:
      - id: collect_metrics
        command: "python collect_metrics.py --output=metrics.prom"
        retry: 2
        timeout: "2m"
      
      - id: push_metrics
        depends_on: ["collect_metrics"]
        command: "curl -X POST http://prometheus-pushgateway:9091/metrics/job/goliteflow --data-binary @metrics.prom"
        retry: 2
        timeout: "30s"
```

### Log Aggregation

Configure log forwarding:

```yaml
version: "1.0"
workflows:
  - name: log_processing
    schedule: "0 */1 * * *"  # Every hour
    tasks:
      - id: collect_logs
        command: "find /var/log/goliteflow -name '*.log' -mtime -1 -exec cat {} \\;"
        retry: 1
        timeout: "5m"
      
      - id: send_to_elasticsearch
        depends_on: ["collect_logs"]
        command: "python send_logs_to_elasticsearch.py --input=logs.json"
        retry: 2
        timeout: "10m"
```

### Health Checks

```yaml
version: "1.0"
workflows:
  - name: health_monitoring
    schedule: "*/2 * * * *"  # Every 2 minutes
    tasks:
      - id: check_goliteflow_health
        command: "goliteflow validate --config=workflows.yml"
        retry: 1
        timeout: "30s"
      
      - id: check_disk_space
        command: "df -h | awk '$5 > 80 {print $0}'"
        retry: 1
        timeout: "10s"
      
      - id: check_memory_usage
        command: "free | awk 'NR==2{printf \"%.2f%%\", $3*100/$2 }'"
        retry: 1
        timeout: "10s"
      
      - id: send_health_status
        depends_on: ["check_goliteflow_health", "check_disk_space", "check_memory_usage"]
        command: "python send_health_status.py"
        retry: 1
        timeout: "30s"
```

## 🔒 Security Considerations

### User Permissions

```bash
# Create dedicated user
sudo useradd -r -s /bin/false goliteflow

# Set proper permissions
sudo chown -R goliteflow:goliteflow /opt/goliteflow
sudo chmod 755 /opt/goliteflow
sudo chmod 644 /opt/goliteflow/workflows.yml
```

### Network Security

```yaml
# Firewall rules (iptables)
-A INPUT -p tcp --dport 8080 -s 10.0.0.0/8 -j ACCEPT
-A INPUT -p tcp --dport 8080 -j DROP
```

### Secrets Management

Use environment variables or secret management:

```yaml
version: "1.0"
workflows:
  - name: secure_workflow
    schedule: "0 2 * * *"
    tasks:
      - id: secure_task
        command: "python secure_script.py --api-key=${API_KEY} --db-password=${DB_PASSWORD}"
        retry: 2
        timeout: "30m"
```

### Container Security

```dockerfile
FROM sintakaridina/goliteflow:latest

# Create non-root user
RUN addgroup -g 1001 -S goliteflow && \
    adduser -u 1001 -S goliteflow -G goliteflow

# Set proper permissions
RUN chown -R goliteflow:goliteflow /app

# Switch to non-root user
USER goliteflow

# Security options
RUN apk --no-cache add ca-certificates
```

## 🚀 Production Checklist

### Pre-deployment

- [ ] Test workflows in staging environment
- [ ] Validate configuration files
- [ ] Set up monitoring and alerting
- [ ] Configure log aggregation
- [ ] Set up backup procedures
- [ ] Test disaster recovery procedures

### Deployment

- [ ] Deploy to production environment
- [ ] Verify service is running
- [ ] Check health endpoints
- [ ] Monitor initial executions
- [ ] Verify report generation
- [ ] Test alerting systems

### Post-deployment

- [ ] Monitor performance metrics
- [ ] Review execution logs
- [ ] Check resource utilization
- [ ] Verify backup procedures
- [ ] Update documentation
- [ ] Train operations team

## 📚 Related Documentation

- [Getting Started](/getting-started) - Quick setup guide
- [Configuration Reference](/configuration) - YAML configuration details
- [CLI Reference](/cli-reference) - Command-line interface
- [Library API](/api) - Go library documentation
- [Examples](/examples) - Real-world use cases

---

<div class="next-steps">
  <h3>Ready to deploy?</h3>
  <p>Choose your deployment method and follow the guide.</p>
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
