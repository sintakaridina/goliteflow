#!/usr/bin/env python3
"""
Generate sample CSV data for testing GoliteFlow file processing workflow
"""

import csv
import random
from datetime import datetime, timedelta
import os

def generate_sample_data():
    """Generate sample e-commerce data"""
    
    # Ensure data directory exists
    os.makedirs('data', exist_ok=True)
    
    # Sample data
    products = ['Laptop', 'Phone', 'Tablet', 'Monitor', 'Keyboard', 'Mouse']
    categories = ['Electronics', 'Computers', 'Accessories']
    customers = ['John Doe', 'Jane Smith', 'Bob Wilson', 'Alice Brown', 'Charlie Davis']
    
    data = []
    
    # Generate 100 sample records
    for i in range(1, 101):
        # Random date within last 30 days
        date = datetime.now() - timedelta(days=random.randint(0, 30))
        
        record = {
            'id': i,
            'date': date.strftime('%Y-%m-%d'),
            'customer': random.choice(customers),
            'product': random.choice(products),
            'category': random.choice(categories),
            'quantity': random.randint(1, 5),
            'price': round(random.uniform(10.99, 999.99), 2),
            'total': 0  # Will be calculated
        }
        
        # Calculate total
        record['total'] = round(record['quantity'] * record['price'], 2)
        data.append(record)
    
    # Write to CSV
    filename = 'data/sample_data.csv'
    with open(filename, 'w', newline='', encoding='utf-8') as csvfile:
        fieldnames = ['id', 'date', 'customer', 'product', 'category', 'quantity', 'price', 'total']
        writer = csv.DictWriter(csvfile, fieldnames=fieldnames)
        
        writer.writeheader()
        for record in data:
            writer.writerow(record)
    
    print(f"✓ Generated {len(data)} sample records in {filename}")
    print(f"  Total sales: ${sum(record['total'] for record in data):,.2f}")
    return filename

if __name__ == '__main__':
    generate_sample_data()