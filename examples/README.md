# GoliteFlow Examples

This directory contains real-world examples that demonstrate GoliteFlow's capabilities. Each example is fully functional and can be run immediately to see actual results and HTML reports.

## 📁 Available Examples

### 1. **File Processing Workflow** (`file-processing-workflow.yml`)

A complete data processing pipeline that:

- ✅ Generates sample e-commerce data (CSV)
- ✅ Validates data quality and structure
- ✅ Processes data with analytics (profit, rankings, categorization)
- ✅ Generates comprehensive summary report
- ✅ Cleans up temporary files and archives results

### 2. **API Monitoring Workflow** (`api-monitoring-workflow.yml`)

Real-time API health monitoring that:

- ✅ Makes HTTP requests to test APIs
- ✅ Validates JSON response structure
- ✅ Logs API status and response times
- ✅ Tracks uptime statistics
- ✅ Generates alerts for failures

### 3. **Backup & Cleanup Workflow** (`backup-cleanup-workflow.yml`)

Automated maintenance operations:

- ✅ Creates data backups
- ✅ Compresses and verifies backups
- ✅ Cleans up old files
- ✅ Monitors disk usage
- ✅ Optimizes system resources

## 🚀 Quick Start

### Prerequisites

- **GoliteFlow binary** installed (see main README)
- **Python 3.6+** for support scripts
- **curl** for API monitoring examples

### Run Your First Example

```bash
# 1. Navigate to examples directory
cd examples/

# 2. Run the file processing workflow
goliteflow run --config=file-processing-workflow.yml

# 3. Generate HTML report
goliteflow report --output=processing_report.html

# 4. View results
open processing_report.html  # macOS
# or
start processing_report.html  # Windows
# or
xdg-open processing_report.html  # Linux
```

## 📊 Expected Results

### File Processing Example

After running, you'll see:

```
data/
├── sample_data.csv          # 100 sample e-commerce records
├── processed_data.csv       # Enhanced data with analytics
└── summary_report.txt       # Business intelligence report

outputs/
└── run_20231015_143022/     # Timestamped archive
    ├── processed_data.csv
    ├── summary_report.txt
    └── cleanup_summary.txt
```

### Sample Output from Summary Report:

```
📊 OVERALL STATISTICS
----------------------------------------
Total Orders: 100
Total Revenue: $52,847.32
Total Profit: $15,854.20
Average Order Value: $528.47
Average Profit Margin: 30.0%

🏆 TOP 5 PRODUCTS BY REVENUE
----------------------------------------
1. Laptop          $15,234.56 ( 23 units sold)
2. Monitor          $12,487.90 ( 19 units sold)
3. Phone            $10,928.34 ( 31 units sold)
```

### API Monitoring Example

Continuous monitoring with logs:

```
logs/
├── api_response.json        # Latest API response
├── api_validation.log       # Validation history
├── api_status.log          # Status timeline
└── api_status_detail.json  # Detailed metrics
```

## 🎯 Example Workflows in Detail

### File Processing Pipeline

This example demonstrates a typical **ETL (Extract, Transform, Load)** workflow:

1. **Extract**: Generate sample sales data
2. **Transform**: Add calculated fields (profit, rankings, categories)
3. **Load**: Create reports and archive results

**Business Value:**

- Data quality validation
- Automated analytics
- Business intelligence reporting
- Clean data archival

### API Health Monitoring

Real-world **DevOps monitoring** scenario:

1. **Health Checks**: HTTP requests to API endpoints
2. **Validation**: JSON structure and data quality
3. **Alerting**: Status logging and failure detection
4. **Reporting**: Uptime statistics and trends

**Business Value:**

- Early problem detection
- SLA monitoring
- Performance tracking
- Incident response

### Backup & Maintenance

Enterprise **data protection** workflow:

1. **Backup Creation**: Copy critical data files
2. **Compression**: Optimize storage space
3. **Verification**: Ensure backup integrity
4. **Cleanup**: Remove old files and optimize resources

**Business Value:**

- Data protection
- Storage optimization
- Automated maintenance
- Compliance requirements

## 🔧 Customizing Examples

### Modify Schedules

Change cron expressions in workflow files:

```yaml
schedule: "0 9 * * *"     # Daily at 9 AM
schedule: "*/15 * * * *"  # Every 15 minutes
schedule: "0 2 * * 0"     # Weekly on Sunday at 2 AM
```

### Add Your Own Tasks

Extend workflows with custom commands:

```yaml
- id: custom_task
  depends_on: ["previous_task"]
  command: "python my_custom_script.py"
  retry: 3
  timeout: "60s"
```

### Environment Variables

Configure scripts with environment variables:

```yaml
- id: upload_to_s3
  command: "aws s3 cp backup.zip s3://my-bucket/"
  env:
    AWS_ACCESS_KEY_ID: "your-key"
    AWS_SECRET_ACCESS_KEY: "your-secret"
```

## 📈 Viewing Results

### HTML Reports

GoliteFlow generates beautiful HTML reports showing:

- ✅ Task execution timeline
- ✅ Success/failure status
- ✅ Execution logs and output
- ✅ Performance metrics
- ✅ Retry attempts and errors

### Log Files

Each example generates structured logs:

- **Execution logs**: Task output and errors
- **Status logs**: Timeline of operations
- **Validation logs**: Data quality checks
- **Performance logs**: Timing and metrics

## 🎛️ Running Examples

### One-Time Execution

```bash
# Run once and exit
goliteflow run --config=file-processing-workflow.yml
```

### Daemon Mode (Scheduled)

```bash
# Run as daemon with cron scheduling
goliteflow run --config=api-monitoring-workflow.yml --daemon
```

### Validate Before Running

```bash
# Check configuration syntax
goliteflow validate --config=backup-cleanup-workflow.yml
```

### Custom Report Output

```bash
# Generate report with custom filename
goliteflow report --output=my_custom_report.html
```

## 🔍 Troubleshooting

### Common Issues

**Python not found:**

```bash
# Make sure Python is in PATH
python --version
python3 --version
```

**Permission errors:**

```bash
# Make scripts executable (Linux/macOS)
chmod +x scripts/*.py
```

**Missing directories:**
The scripts automatically create required directories, but you can create them manually:

```bash
mkdir -p data logs outputs backups
```

### Debug Mode

Add verbose logging to see detailed execution:

```bash
goliteflow run --config=workflow.yml --verbose
```

## 💡 Next Steps

1. **Start with file processing** - Easy to understand and see results
2. **Try API monitoring** - See real-time monitoring in action
3. **Customize for your needs** - Modify scripts and schedules
4. **Build your own workflows** - Use examples as templates
5. **Share your examples** - Contribute back to the community

## 🤝 Contributing Examples

Have a great workflow example? We'd love to include it! Please:

1. Create a complete, working example
2. Include supporting scripts and documentation
3. Test thoroughly on multiple platforms
4. Submit a pull request with:
   - Workflow YAML file
   - Supporting Python scripts
   - Sample data (if needed)
   - Documentation updates

---

**Happy workflow automation with GoliteFlow!** 🚀

For more information, see the main [GoliteFlow documentation](../README.md).
