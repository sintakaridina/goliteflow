#!/usr/bin/env python3
"""
Process CSV data - add calculated fields and filtering
"""

import csv
import sys
from datetime import datetime
from collections import defaultdict

def process_data(input_file, output_file):
    """Process CSV data and add analytics"""
    
    if not input_file or not output_file:
        print("❌ Error: Input and output filenames required")
        sys.exit(1)
    
    try:
        processed_data = []
        customer_totals = defaultdict(float)
        product_counts = defaultdict(int)
        
        # Read and process data
        with open(input_file, 'r', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            
            for row in reader:
                # Convert data types
                row['quantity'] = int(row['quantity'])
                row['price'] = float(row['price'])
                row['total'] = float(row['total'])
                
                # Add processing date
                row['processed_date'] = datetime.now().strftime('%Y-%m-%d %H:%M:%S')
                
                # Add profit margin (assume 30% markup)
                cost = row['price'] * 0.7
                profit = row['total'] - (cost * row['quantity'])
                row['profit'] = round(profit, 2)
                row['profit_margin'] = round((profit / row['total']) * 100, 2) if row['total'] > 0 else 0
                
                # Add categorization
                if row['total'] > 500:
                    row['order_size'] = 'Large'
                elif row['total'] > 100:
                    row['order_size'] = 'Medium'
                else:
                    row['order_size'] = 'Small'
                
                # Update aggregates
                customer_totals[row['customer']] += row['total']
                product_counts[row['product']] += row['quantity']
                
                processed_data.append(row)
        
        # Add customer ranking
        sorted_customers = sorted(customer_totals.items(), key=lambda x: x[1], reverse=True)
        customer_rankings = {customer: rank + 1 for rank, (customer, _) in enumerate(sorted_customers)}
        
        # Add rankings to data
        for row in processed_data:
            row['customer_rank'] = customer_rankings[row['customer']]
            row['customer_total'] = round(customer_totals[row['customer']], 2)
        
        # Write processed data
        fieldnames = ['id', 'date', 'customer', 'product', 'category', 'quantity', 'price', 'total', 
                     'profit', 'profit_margin', 'order_size', 'customer_rank', 'customer_total', 'processed_date']
        
        with open(output_file, 'w', newline='', encoding='utf-8') as csvfile:
            writer = csv.DictWriter(csvfile, fieldnames=fieldnames)
            writer.writeheader()
            for row in processed_data:
                writer.writerow(row)
        
        print(f"✓ Data processed successfully")
        print(f"  Input: {input_file}")
        print(f"  Output: {output_file}")
        print(f"  Processed {len(processed_data)} records")
        print(f"  Total revenue: ${sum(row['total'] for row in processed_data):,.2f}")
        print(f"  Total profit: ${sum(row['profit'] for row in processed_data):,.2f}")
        print(f"  Top customer: {sorted_customers[0][0]} (${sorted_customers[0][1]:,.2f})")
        
    except FileNotFoundError as e:
        print(f"❌ File not found: {e}")
        sys.exit(1)
    except Exception as e:
        print(f"❌ Processing error: {e}")
        sys.exit(1)

if __name__ == '__main__':
    if len(sys.argv) != 3:
        print("Usage: python process_data.py <input.csv> <output.csv>")
        sys.exit(1)
    
    process_data(sys.argv[1], sys.argv[2])