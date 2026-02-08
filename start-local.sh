#!/bin/bash

# Overss Local-Only Start Script

echo "🚀 Starting Overss RSS Feed Server (Local Network Only)..."
echo ""

# Check if binary exists
if [ ! -f "./overss" ]; then
    echo "📦 Building Overss..."
    go build -o overss
    if [ $? -ne 0 ]; then
        echo "❌ Build failed. Please check for errors."
        exit 1
    fi
    echo "✅ Build successful!"
    echo ""
fi

# Create audiobooks directory if it doesn't exist
if [ ! -d "./audiobooks" ]; then
    echo "📁 Creating audiobooks directory..."
    mkdir -p audiobooks
fi

echo "🌐 The server will display all available network addresses on startup"
echo "📡 Access from other devices using the Network URLs shown"
echo ""

# Start the server (it will display all access URLs)
./overss
