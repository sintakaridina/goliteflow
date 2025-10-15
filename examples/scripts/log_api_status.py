#!/usr/bin/env python3
"""
Log API monitoring status and generate alerts
"""

import os
import json
from datetime import datetime

def log_api_status():
    """Log current API monitoring status"""
    
    timestamp = datetime.now()
    status_entry = {
        'timestamp': timestamp.isoformat(),
        'check_type': 'api_health',
        'status': 'healthy',
        'response_time': None,
        'notes': []
    }
    
    # Check if API response file exists and is recent
    response_file = 'logs/api_response.json'
    if os.path.exists(response_file):
        file_age = timestamp.timestamp() - os.path.getmtime(response_file)
        
        if file_age > 300:  # 5 minutes
            status_entry['status'] = 'stale'
            status_entry['notes'].append(f"Response file is {file_age:.0f} seconds old")
        else:
            status_entry['notes'].append(f"Response file is {file_age:.0f} seconds old")
            
        # Try to read response size
        try:
            with open(response_file, 'r') as f:
                response = json.load(f)
                response_size = len(json.dumps(response))
                status_entry['response_size'] = response_size
                status_entry['notes'].append(f"Response size: {response_size} bytes")
        except Exception as e:
            status_entry['status'] = 'error'
            status_entry['notes'].append(f"Failed to read response: {e}")
    else:
        status_entry['status'] = 'missing'
        status_entry['notes'].append("API response file not found")
    
    # Write to status log
    status_log = 'logs/api_status.log'
    with open(status_log, 'a') as f:
        log_line = f"{status_entry['timestamp']}: {status_entry['status'].upper()}"
        if status_entry['notes']:
            log_line += f" - {'; '.join(status_entry['notes'])}"
        f.write(log_line + "\n")
    
    # Write detailed JSON log
    json_log = 'logs/api_status_detail.json'
    if os.path.exists(json_log):
        with open(json_log, 'r') as f:
            try:
                logs = json.load(f)
            except json.JSONDecodeError:
                logs = []
    else:
        logs = []
    
    logs.append(status_entry)
    
    # Keep only last 100 entries
    if len(logs) > 100:
        logs = logs[-100:]
    
    with open(json_log, 'w') as f:
        json.dump(logs, f, indent=2)
    
    # Generate summary
    recent_logs = [log for log in logs if log['timestamp'] > (timestamp.timestamp() - 3600)]  # Last hour
    healthy_count = len([log for log in recent_logs if log['status'] == 'healthy'])
    total_recent = len(recent_logs)
    
    health_percentage = (healthy_count / total_recent * 100) if total_recent > 0 else 0
    
    # Output status
    status_icon = {
        'healthy': '✅',
        'stale': '⚠️ ',
        'error': '❌',
        'missing': '❌'
    }
    
    print(f"{status_icon.get(status_entry['status'], '❓')} API Status: {status_entry['status'].upper()}")
    
    if status_entry['notes']:
        for note in status_entry['notes']:
            print(f"  • {note}")
    
    if total_recent > 0:
        print(f"  • Last hour health: {health_percentage:.1f}% ({healthy_count}/{total_recent})")
    
    print(f"  • Status logged to: {status_log}")
    print(f"  • Detail logged to: {json_log}")
    
    # Exit with error code if status is not healthy
    if status_entry['status'] != 'healthy':
        exit(1)

if __name__ == '__main__':
    # Ensure logs directory exists
    os.makedirs('logs', exist_ok=True)
    log_api_status()