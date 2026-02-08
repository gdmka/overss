#!/bin/bash

# Overss Quick Start Script

echo "🚀 Starting Overss RSS Feed Server..."
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

echo "🌐 Starting local server and ngrok tunnel..."
echo ""

# Start the server in the background
./overss &
SERVER_PID=$!

# Wait a moment for server to start
sleep 2

# Check if ngrok is installed
if command -v ngrok &> /dev/null; then
    echo "🌍 Starting ngrok tunnel for internet access..."
    echo ""
    ngrok http 8083
else
    echo "⚠️  ngrok not found. Server running locally only."
    echo "📥 Install ngrok from: https://ngrok.com/download"
    echo ""
    echo "🌐 Local access URLs:"
    echo "  http://localhost:8083"
    echo ""
    echo "Press Ctrl+C to stop the server"

    # Wait for server process
    wait $SERVER_PID
fi
