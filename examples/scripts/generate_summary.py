#!/usr/bin/env python3
"""
Generate summary report from processed CSV data
"""

import csv
import sys
from datetime import datetime
from collections import defaultdict, Counter

def generate_summary(input_file, output_file):
    """Generate comprehensive summary report"""
    
    try:
        data = []
        
        # Read processed data
        with open(input_file, 'r', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                # Convert numeric fields
                row['quantity'] = int(row['quantity'])
                row['price'] = float(row['price'])
                row['total'] = float(row['total'])
                row['profit'] = float(row['profit'])
                row['profit_margin'] = float(row['profit_margin'])
                data.append(row)
        
        if not data:
            print("❌ No data found in input file")
            sys.exit(1)
        
        # Calculate statistics
        total_orders = len(data)
        total_revenue = sum(row['total'] for row in data)
        total_profit = sum(row['profit'] for row in data)
        avg_order_value = total_revenue / total_orders
        avg_profit_margin = sum(row['profit_margin'] for row in data) / total_orders
        
        # Category analysis
        category_sales = defaultdict(float)
        category_counts = defaultdict(int)
        for row in data:
            category_sales[row['category']] += row['total']
            category_counts[row['category']] += 1
        
        # Product analysis
        product_sales = defaultdict(float)
        product_quantities = defaultdict(int)
        for row in data:
            product_sales[row['product']] += row['total']
            product_quantities[row['product']] += row['quantity']
        
        # Customer analysis
        customer_orders = Counter(row['customer'] for row in data)
        customer_sales = defaultdict(float)
        for row in data:
            customer_sales[row['customer']] += row['total']
        
        # Order size analysis
        order_sizes = Counter(row['order_size'] for row in data)
        
        # Generate report
        report_lines = []
        report_lines.append("=" * 80)
        report_lines.append("GOLITEFLOW SALES ANALYTICS REPORT")
        report_lines.append("=" * 80)
        report_lines.append(f"Generated: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")
        report_lines.append(f"Data source: {input_file}")
        report_lines.append("")
        
        # Overall Statistics
        report_lines.append("📊 OVERALL STATISTICS")
        report_lines.append("-" * 40)
        report_lines.append(f"Total Orders: {total_orders:,}")
        report_lines.append(f"Total Revenue: ${total_revenue:,.2f}")
        report_lines.append(f"Total Profit: ${total_profit:,.2f}")
        report_lines.append(f"Average Order Value: ${avg_order_value:.2f}")
        report_lines.append(f"Average Profit Margin: {avg_profit_margin:.1f}%")
        report_lines.append("")
        
        # Category Performance
        report_lines.append("📈 CATEGORY PERFORMANCE")
        report_lines.append("-" * 40)
        sorted_categories = sorted(category_sales.items(), key=lambda x: x[1], reverse=True)
        for category, sales in sorted_categories:
            count = category_counts[category]
            avg_per_order = sales / count
            report_lines.append(f"{category:15} ${sales:>10,.2f} ({count:>3} orders, avg: ${avg_per_order:>6.2f})")
        report_lines.append("")
        
        # Top Products
        report_lines.append("🏆 TOP 5 PRODUCTS BY REVENUE")
        report_lines.append("-" * 40)
        sorted_products = sorted(product_sales.items(), key=lambda x: x[1], reverse=True)[:5]
        for i, (product, sales) in enumerate(sorted_products, 1):
            qty = product_quantities[product]
            report_lines.append(f"{i}. {product:15} ${sales:>10,.2f} ({qty:>3} units sold)")
        report_lines.append("")
        
        # Top Customers
        report_lines.append("👑 TOP 5 CUSTOMERS BY REVENUE")
        report_lines.append("-" * 40)
        sorted_customers = sorted(customer_sales.items(), key=lambda x: x[1], reverse=True)[:5]
        for i, (customer, sales) in enumerate(sorted_customers, 1):
            orders = customer_orders[customer]
            avg_per_order = sales / orders
            report_lines.append(f"{i}. {customer:15} ${sales:>10,.2f} ({orders:>2} orders, avg: ${avg_per_order:>6.2f})")
        report_lines.append("")
        
        # Order Size Distribution
        report_lines.append("💰 ORDER SIZE DISTRIBUTION")
        report_lines.append("-" * 40)
        for size in ['Large', 'Medium', 'Small']:
            count = order_sizes[size]
            percentage = (count / total_orders) * 100
            report_lines.append(f"{size:8} orders: {count:>4} ({percentage:>5.1f}%)")
        report_lines.append("")
        
        # Date Range Analysis
        dates = [row['date'] for row in data]
        date_range = f"{min(dates)} to {max(dates)}"
        report_lines.append("📅 DATA SUMMARY")
        report_lines.append("-" * 40)
        report_lines.append(f"Date Range: {date_range}")
        report_lines.append(f"Unique Customers: {len(customer_sales)}")
        report_lines.append(f"Unique Products: {len(product_sales)}")
        report_lines.append(f"Total Categories: {len(category_sales)}")
        report_lines.append("")
        
        report_lines.append("=" * 80)
        report_lines.append("Report generated by GoliteFlow Data Processing Pipeline")
        report_lines.append("=" * 80)
        
        # Write report
        with open(output_file, 'w', encoding='utf-8') as f:
            f.write('\n'.join(report_lines))
        
        print(f"✓ Summary report generated")
        print(f"  Input: {input_file}")
        print(f"  Output: {output_file}")
        print(f"  Total orders analyzed: {total_orders:,}")
        print(f"  Total revenue: ${total_revenue:,.2f}")
        
    except FileNotFoundError as e:
        print(f"❌ File not found: {e}")
        sys.exit(1)
    except Exception as e:
        print(f"❌ Report generation error: {e}")
        sys.exit(1)

if __name__ == '__main__':
    if len(sys.argv) != 3:
        print("Usage: python generate_summary.py <input.csv> <output.txt>")
        sys.exit(1)
    
    generate_summary(sys.argv[1], sys.argv[2])