#!/usr/bin/env python3
"""
Cleanup temporary files and organize outputs
"""

import os
import shutil
from datetime import datetime

def cleanup():
    """Clean up temporary files and organize outputs"""
    
    print("🧹 Starting cleanup process...")
    
    # Create outputs directory with timestamp
    timestamp = datetime.now().strftime('%Y%m%d_%H%M%S')
    output_dir = f"outputs/run_{timestamp}"
    os.makedirs(output_dir, exist_ok=True)
    
    files_moved = 0
    files_removed = 0
    
    # Move important outputs to timestamped directory
    important_files = [
        'data/processed_data.csv',
        'data/summary_report.txt'
    ]
    
    for file in important_files:
        if os.path.exists(file):
            filename = os.path.basename(file)
            dest = os.path.join(output_dir, filename)
            shutil.copy2(file, dest)
            print(f"  📁 Archived: {file} → {dest}")
            files_moved += 1
    
    # Clean up temporary files (keep originals)
    temp_patterns = [
        'data/*.tmp',
        'logs/*.tmp', 
        '*.pyc',
        '__pycache__'
    ]
    
    # Remove __pycache__ directories
    for root, dirs, files in os.walk('.'):
        if '__pycache__' in dirs:
            cache_dir = os.path.join(root, '__pycache__')
            shutil.rmtree(cache_dir)
            print(f"  🗑️  Removed: {cache_dir}")
            files_removed += 1
    
    # Remove .pyc files
    for root, dirs, files in os.walk('.'):
        for file in files:
            if file.endswith('.pyc'):
                pyc_file = os.path.join(root, file)
                os.remove(pyc_file)
                print(f"  🗑️  Removed: {pyc_file}")
                files_removed += 1
    
    # Create cleanup summary
    summary_file = os.path.join(output_dir, 'cleanup_summary.txt')
    with open(summary_file, 'w') as f:
        f.write("GOLITEFLOW CLEANUP SUMMARY\n")
        f.write("=" * 40 + "\n")
        f.write(f"Timestamp: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}\n")
        f.write(f"Files archived: {files_moved}\n")
        f.write(f"Files removed: {files_removed}\n")
        f.write(f"Output directory: {output_dir}\n")
        f.write("\nArchived files:\n")
        for file in important_files:
            if os.path.exists(file):
                f.write(f"  - {os.path.basename(file)}\n")
    
    print(f"✓ Cleanup completed")
    print(f"  Files archived: {files_moved}")
    print(f"  Files removed: {files_removed}")
    print(f"  Output directory: {output_dir}")
    print(f"  Summary: {summary_file}")

if __name__ == '__main__':
    # Create outputs directory if it doesn't exist
    os.makedirs('outputs', exist_ok=True)
    cleanup()