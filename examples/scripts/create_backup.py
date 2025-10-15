#!/usr/bin/env python3
"""
Create backup of specified directory
"""

import os
import shutil
import sys
from datetime import datetime

def create_backup(source_dir, backup_dir):
    """Create backup of source directory"""
    
    if not source_dir or not backup_dir:
        print("❌ Error: Source and backup directories required")
        sys.exit(1)
    
    try:
        # Ensure backup directory exists
        os.makedirs(backup_dir, exist_ok=True)
        
        # Create timestamped backup name
        timestamp = datetime.now().strftime('%Y%m%d_%H%M%S')
        backup_name = f"backup_{timestamp}"
        backup_path = os.path.join(backup_dir, backup_name)
        
        if not os.path.exists(source_dir):
            print(f"❌ Source directory not found: {source_dir}")
            sys.exit(1)
        
        # Count files to backup
        total_files = 0
        total_size = 0
        
        for root, dirs, files in os.walk(source_dir):
            for file in files:
                file_path = os.path.join(root, file)
                if os.path.exists(file_path):
                    total_files += 1
                    total_size += os.path.getsize(file_path)
        
        if total_files == 0:
            print(f"⚠️  No files found in source directory: {source_dir}")
            # Create empty backup directory anyway
            os.makedirs(backup_path, exist_ok=True)
        else:
            # Copy directory tree
            shutil.copytree(source_dir, backup_path, dirs_exist_ok=True)
        
        # Create backup manifest
        manifest_file = os.path.join(backup_path, 'backup_manifest.txt')
        with open(manifest_file, 'w') as f:
            f.write("GOLITEFLOW BACKUP MANIFEST\n")
            f.write("=" * 40 + "\n")
            f.write(f"Created: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}\n")
            f.write(f"Source: {os.path.abspath(source_dir)}\n")
            f.write(f"Backup: {os.path.abspath(backup_path)}\n")
            f.write(f"Files: {total_files}\n")
            f.write(f"Size: {total_size:,} bytes ({total_size / 1024 / 1024:.2f} MB)\n")
            f.write("\nFile listing:\n")
            
            if total_files > 0:
                for root, dirs, files in os.walk(backup_path):
                    for file in files:
                        if file != 'backup_manifest.txt':
                            rel_path = os.path.relpath(os.path.join(root, file), backup_path)
                            file_size = os.path.getsize(os.path.join(root, file))
                            f.write(f"  {rel_path} ({file_size:,} bytes)\n")
        
        print(f"✓ Backup created successfully")
        print(f"  Source: {source_dir}")
        print(f"  Backup: {backup_path}")
        print(f"  Files: {total_files:,}")
        print(f"  Size: {total_size / 1024 / 1024:.2f} MB")
        print(f"  Manifest: {manifest_file}")
        
        return backup_path
        
    except PermissionError as e:
        print(f"❌ Permission error: {e}")
        sys.exit(1)
    except Exception as e:
        print(f"❌ Backup error: {e}")
        sys.exit(1)

if __name__ == '__main__':
    if len(sys.argv) != 3:
        print("Usage: python create_backup.py <source_dir> <backup_dir>")
        print("Example: python create_backup.py data/ backups/")
        sys.exit(1)
    
    create_backup(sys.argv[1], sys.argv[2])