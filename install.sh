#!/bin/bash
# Installation script for GCM

set -e

echo " Installing GCM (Git Commit Manager)..."
echo ""

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo " Error: Go is not installed"
    echo "Please install Go from https://golang.org/dl/"
    exit 1
fi

echo "✓ Go is installed: $(go version)"
echo ""

# Build the binary
echo " Building gcm..."
go build -o gcm .

if [ ! -f "./gcm" ]; then
    echo " Build failed"
    exit 1
fi

echo "✓ Build successful"
echo ""

# Offer to install globally
read -p "Install gcm globally to /usr/local/bin? (y/n) " -n 1 -r
echo ""

if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo " Installing globally (may require password)..."
    sudo cp gcm /usr/local/bin/gcm
    sudo chmod +x /usr/local/bin/gcm
    echo "✓ Installed to /usr/local/bin/gcm"
    echo ""
    echo "You can now run 'gcm' from any directory!"
else
    echo "Skipped global installation."
    echo ""
    echo "To use gcm from this directory:"
    echo "  ./gcm"
    echo ""
    echo "Or add to your PATH:"
    echo "  export PATH=\"\$PATH:$(pwd)\""
fi

echo ""
echo " Installation complete!"
echo ""
echo " Quick start:"
echo "  1. Navigate to a git repository"
echo "  2. Make some changes"
echo "  3. Run: gcm"
echo ""
echo " Documentation:"
echo "  - README.md      - Full documentation"
echo "  - QUICKSTART.md  - Getting started guide"
echo ""
echo "Happy committing! "

