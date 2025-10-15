# Screenshots

This directory contains screenshots for the GoliteFlow documentation.

## Required Screenshots

### HTML Report Dashboard
- **File**: `html-report.png`
- **Description**: Screenshot of the HTML report dashboard showing workflow execution history, task details, and statistics
- **Size**: 1200x800px recommended

### CLI Usage
- **File**: `cli-usage.png`
- **Description**: Screenshot of CLI commands in action showing workflow validation, execution, and report generation
- **Size**: 800x600px recommended

## How to Generate Screenshots

1. **HTML Report Screenshot**:
   ```bash
   # Run a sample workflow
   goliteflow run --config=examples/sample-workflow.yml
   
   # Generate report
   goliteflow report --output=report.html
   
   # Open report.html in browser and take screenshot
   ```

2. **CLI Usage Screenshot**:
   ```bash
   # Take screenshot of terminal showing:
   goliteflow validate --config=workflows.yml
   goliteflow run --config=workflows.yml
   goliteflow report --output=report.html
   ```

## Placeholder Images

Until real screenshots are available, you can use placeholder images:
- [Placeholder.com](https://via.placeholder.com/1200x800/3182ce/ffffff?text=HTML+Report+Dashboard)
- [Placeholder.com](https://via.placeholder.com/800x600/2d3748/ffffff?text=CLI+Usage)
