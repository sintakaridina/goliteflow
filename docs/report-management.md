# GoliteFlow Report Management

This document describes the advanced report management features in GoliteFlow that address scalability concerns for production deployments.

## Problem Statement

In production environments where workflows run continuously over months or years, HTML reports can grow unboundedly in size, causing:

- Large file sizes (potentially GBs of data)
- Slow loading times
- High memory usage
- Poor user experience

## Solution: Enhanced Report Management

GoliteFlow provides a comprehensive report management system with:

### 1. Enhanced HTML Reporter

The enhanced reporter (`report-enhanced` command) provides:

- **Automatic Rotation**: Limits the number of executions shown in the main report
- **Intelligent Archival**: Archives old execution data by month
- **Configurable Retention**: Customizable policies for data retention
- **Performance Optimization**: Fast loading with paginated data

#### Usage

```bash
# Basic enhanced report
goliteflow report-enhanced -o report.html

# Custom configuration
goliteflow report-enhanced \
  --max-executions 100 \
  --archive-after 15 \
  --cleanup-after 60 \
  --output enhanced_report.html
```

#### Configuration Options

- `--max-executions`: Maximum executions in main report (default: 50)
- `--archive-after`: Archive reports after N days (default: 30)
- `--cleanup-after`: Delete archived reports after N days (default: 90)
- `--page-size`: Executions per page for pagination (default: 20)
- `--pagination`: Enable/disable pagination (default: true)
- `--report-dir`: Directory for active reports (default: "reports")
- `--archive-dir`: Directory for archived reports (default: "reports/archive")

### 2. Report Management Commands

#### Archive Management

```bash
# Archive old reports manually
goliteflow report-manage archive --days 30

# Archive with custom settings
goliteflow report-manage archive \
  --days 15 \
  --report-dir "custom/reports" \
  --archive-dir "custom/archive"
```

#### Cleanup Operations

```bash
# Cleanup archived reports older than default (90 days)
goliteflow report-manage cleanup

# Custom cleanup threshold
goliteflow report-manage cleanup --days 60
```

#### Statistics and Monitoring

```bash
# View report statistics
goliteflow report-manage stats

# Example output:
# 📊 GoliteFlow Report Statistics
# ================================
# Total Executions: 1,250
# Completed: 1,180
# Failed: 70
# Recent (7 days): 45
# Success Rate: 94.4%
#
# 📁 Storage Information
# Report Directory: reports
# Archive Directory: reports/archive
```

### 3. Architecture Overview

#### Data Flow

1. **Active Reports**: Recent executions stored in `reports/` directory
2. **Archival Process**: Executions older than threshold moved to `reports/archive/{YYYY-MM}/`
3. **Cleanup Process**: Archive directories older than cleanup threshold are deleted
4. **Report Generation**: Main report shows only recent executions with links to archives

#### File Structure

```
reports/
├── index.json              # Main execution index
├── executions/
│   ├── exec_001.json       # Individual execution files
│   ├── exec_002.json
│   └── ...
├── archive/
│   ├── 2024-01/            # Archived by month
│   │   ├── index.json
│   │   └── executions/
│   ├── 2024-02/
│   └── ...
└── enhanced_report.html    # Generated report
```

#### Performance Benefits

- **Constant Time Loading**: Main report always loads ~50 executions regardless of total history
- **Memory Efficient**: Archived data not loaded unless specifically requested
- **Scalable Storage**: Archived data organized by month for efficient access
- **Automatic Cleanup**: Old archives automatically removed to prevent disk space issues

### 4. Production Deployment

#### Automated Archival

Set up automated archival with cron jobs or scheduled tasks:

```bash
# Daily archival (Linux/macOS cron)
0 2 * * * /path/to/goliteflow report-manage archive

# Weekly cleanup (Linux/macOS cron)
0 3 * * 0 /path/to/goliteflow report-manage cleanup
```

For Windows Task Scheduler:

```powershell
# Create daily task for archival
schtasks /create /tn "GoliteFlow Archive" /tr "C:\path\to\goliteflow.exe report-manage archive" /sc daily /st 02:00

# Create weekly task for cleanup
schtasks /create /tn "GoliteFlow Cleanup" /tr "C:\path\to\goliteflow.exe report-manage cleanup" /sc weekly /st 03:00
```

#### Configuration Recommendations

For different deployment scenarios:

**High-Volume Production (1000+ executions/day)**

```bash
goliteflow report-enhanced \
  --max-executions 25 \
  --archive-after 7 \
  --cleanup-after 30 \
  --page-size 10
```

**Medium-Volume Production (100-1000 executions/day)**

```bash
goliteflow report-enhanced \
  --max-executions 50 \
  --archive-after 30 \
  --cleanup-after 90 \
  --page-size 20
```

**Low-Volume Production (<100 executions/day)**

```bash
goliteflow report-enhanced \
  --max-executions 100 \
  --archive-after 60 \
  --cleanup-after 180 \
  --page-size 25
```

### 5. Monitoring and Maintenance

#### Health Checks

Monitor report system health:

```bash
# Regular statistics check
goliteflow report-manage stats

# Archive size monitoring (Linux/macOS)
du -sh reports/archive/

# Windows PowerShell
Get-ChildItem -Path "reports\archive" -Recurse | Measure-Object -Property Length -Sum
```

#### Troubleshooting

Common issues and solutions:

**Large Archive Sizes**

- Reduce `--cleanup-after` value
- Run cleanup more frequently
- Consider manual cleanup of specific months

**Slow Report Generation**

- Reduce `--max-executions` value
- Enable pagination with smaller `--page-size`
- Check for corrupted execution files

**Missing Historical Data**

- Check archive directories
- Verify cleanup settings
- Review archival logs

### 6. Migration from Basic Reports

To migrate from basic HTML reports to enhanced reports:

1. **Backup existing reports**:

   ```bash
   cp report.html report_backup.html
   ```

2. **Generate enhanced report**:

   ```bash
   goliteflow report-enhanced -o report.html
   ```

3. **Set up automated maintenance**:
   ```bash
   # Add to crontab or task scheduler
   goliteflow report-manage archive
   goliteflow report-manage cleanup
   ```

The enhanced report system is fully backward compatible and will import existing execution data automatically.

## Conclusion

The GoliteFlow enhanced report management system provides enterprise-grade scalability for production deployments, ensuring that reports remain fast and manageable regardless of execution volume or deployment duration.

Key benefits:

- ✅ Constant loading times regardless of historical data size
- ✅ Automatic data archival and cleanup
- ✅ Configurable retention policies
- ✅ Production-ready automation
- ✅ Backward compatibility with existing reports
