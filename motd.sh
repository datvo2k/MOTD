#!/bin/bash

# Check if script is run with sudo/root privileges
if [ "$EUID" -ne 0 ]; then 
    echo "Please run as root (use sudo)"
    exit 1
fi

# Go to the directory containing Makefile
cd "$(dirname "$0")" || exit 1

# Run make
echo "Building project..."
make

# Check if make was successful
if [ $? -ne 0 ]; then
    echo "Build failed"
    exit 1
fi

# Assuming your binary is named 'motd' and is in current directory
# Copy binary to /usr/bin
echo "Copying binary to /usr/bin..."
cp motd /usr/bin/
cp config.json /etc/
chmod 755 /usr/bin/motd

