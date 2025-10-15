#!/usr/bin/env python3
"""
Validate API response and log results
"""

import json
import sys
from datetime import datetime

def validate_api_response(response_file):
    """Validate JSON API response"""
    
    try:
        with open(response_file, 'r') as f:
            response = json.load(f)
        
        # Basic validation
        required_fields = ['origin', 'url']  # httpbin.org/json returns these
        
        validation_results = {
            'timestamp': datetime.now().isoformat(),
            'file': response_file,
            'status': 'success',
            'errors': [],
            'warnings': [],
            'data_size': len(json.dumps(response))
        }
        
        # Check required fields
        for field in required_fields:
            if field not in response:
                validation_results['errors'].append(f"Missing required field: {field}")
        
        # Additional checks
        if 'origin' in response:
            if not response['origin'].replace('.', '').replace(':', '').replace(',', '').replace(' ', '').isdigit():
                validation_results['warnings'].append("Origin field doesn't look like IP address")
        
        if 'url' in response:
            if not response['url'].startswith('http'):
                validation_results['errors'].append("URL field doesn't start with http")
        
        # Check data freshness (should be recent)
        if validation_results['data_size'] < 50:
            validation_results['warnings'].append("Response seems unusually small")
        
        # Set final status
        if validation_results['errors']:
            validation_results['status'] = 'failed'
        elif validation_results['warnings']:
            validation_results['status'] = 'warning'
        
        # Log results
        log_file = 'logs/api_validation.log'
        with open(log_file, 'a') as f:
            f.write(f"{datetime.now().isoformat()}: {validation_results['status'].upper()} - ")
            f.write(f"Size: {validation_results['data_size']} bytes")
            if validation_results['errors']:
                f.write(f" - Errors: {len(validation_results['errors'])}")
            if validation_results['warnings']:
                f.write(f" - Warnings: {len(validation_results['warnings'])}")
            f.write("\n")
        
        # Output results
        if validation_results['status'] == 'failed':
            print("❌ API response validation failed")
            for error in validation_results['errors']:
                print(f"  • {error}")
            sys.exit(1)
        elif validation_results['status'] == 'warning':
            print("⚠️  API response has warnings")
            for warning in validation_results['warnings']:
                print(f"  • {warning}")
        else:
            print("✓ API response validation passed")
        
        print(f"  Response size: {validation_results['data_size']} bytes")
        print(f"  Log updated: {log_file}")
        
    except FileNotFoundError:
        print(f"❌ Response file not found: {response_file}")
        sys.exit(1)
    except json.JSONDecodeError as e:
        print(f"❌ Invalid JSON in response: {e}")
        sys.exit(1)
    except Exception as e:
        print(f"❌ Validation error: {e}")
        sys.exit(1)

if __name__ == '__main__':
    if len(sys.argv) != 2:
        print("Usage: python validate_api_response.py <response.json>")
        sys.exit(1)
    
    validate_api_response(sys.argv[1])