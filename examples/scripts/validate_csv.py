#!/usr/bin/env python3
"""
Validate CSV file structure and data quality
"""

import csv
import sys
from datetime import datetime

def validate_csv(filename):
    """Validate CSV file structure and data"""
    
    if not filename:
        print("❌ Error: No filename provided")
        sys.exit(1)
    
    try:
        with open(filename, 'r', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            
            # Expected columns
            expected_columns = ['id', 'date', 'customer', 'product', 'category', 'quantity', 'price', 'total']
            
            # Check columns
            missing_columns = set(expected_columns) - set(reader.fieldnames)
            if missing_columns:
                print(f"❌ Missing columns: {missing_columns}")
                sys.exit(1)
            
            # Validate data
            row_count = 0
            errors = []
            
            for row_num, row in enumerate(reader, start=2):  # Start at 2 (header is row 1)
                row_count += 1
                
                # Validate ID
                try:
                    int(row['id'])
                except ValueError:
                    errors.append(f"Row {row_num}: Invalid ID '{row['id']}'")
                
                # Validate date
                try:
                    datetime.strptime(row['date'], '%Y-%m-%d')
                except ValueError:
                    errors.append(f"Row {row_num}: Invalid date '{row['date']}'")
                
                # Validate numeric fields
                try:
                    quantity = int(row['quantity'])
                    if quantity <= 0:
                        errors.append(f"Row {row_num}: Quantity must be positive")
                except ValueError:
                    errors.append(f"Row {row_num}: Invalid quantity '{row['quantity']}'")
                
                try:
                    price = float(row['price'])
                    if price <= 0:
                        errors.append(f"Row {row_num}: Price must be positive")
                except ValueError:
                    errors.append(f"Row {row_num}: Invalid price '{row['price']}'")
                
                try:
                    float(row['total'])
                except ValueError:
                    errors.append(f"Row {row_num}: Invalid total '{row['total']}'")
                
                # Check required text fields
                if not row['customer'].strip():
                    errors.append(f"Row {row_num}: Customer name is empty")
                if not row['product'].strip():
                    errors.append(f"Row {row_num}: Product name is empty")
            
            if errors:
                print(f"❌ Validation failed with {len(errors)} errors:")
                for error in errors[:10]:  # Show first 10 errors
                    print(f"  • {error}")
                if len(errors) > 10:
                    print(f"  ... and {len(errors) - 10} more errors")
                sys.exit(1)
            
            print(f"✓ CSV validation passed")
            print(f"  File: {filename}")
            print(f"  Rows: {row_count}")
            print(f"  Columns: {len(reader.fieldnames)}")
            
    except FileNotFoundError:
        print(f"❌ File not found: {filename}")
        sys.exit(1)
    except Exception as e:
        print(f"❌ Validation error: {e}")
        sys.exit(1)

if __name__ == '__main__':
    if len(sys.argv) != 2:
        print("Usage: python validate_csv.py <filename.csv>")
        sys.exit(1)
    
    validate_csv(sys.argv[1])