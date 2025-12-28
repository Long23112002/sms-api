#!/bin/bash

# Script to remove BOM (Byte Order Mark) from .env files
# Usage: ./fix-env-bom.sh

echo "Fixing BOM in .env files..."

# Function to remove BOM from a file
remove_bom() {
    local file=$1
    if [ -f "$file" ]; then
        echo "Processing $file..."
        # Remove BOM using sed
        sed -i '1s/^\xEF\xBB\xBF//' "$file"
        # Alternative method using dos2unix if available
        # dos2unix "$file" 2>/dev/null || true
        echo "✓ Fixed $file"
    else
        echo "⚠ File not found: $file"
    fi
}

# Fix web/.env
remove_bom "./web/.env"

# Fix api/.env
remove_bom "./api/.env"

echo "Done! You can now run: docker-compose up -d"

