#!/bin/bash

# LeetCode quick test script
# Usage: ./lc-test.sh <problem_number>
# Example: ./lc-test.sh 0001

set -e

if [ -z "$1" ]; then
    echo "Usage: $0 <problem_number>"
    echo "Example: $0 0001"
    exit 1
fi

# Get script directory
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
LEETCODE_DIR="$SCRIPT_DIR/leetcode"

# Pad problem number to 4 digits
NUM=$(printf "%04d" "${1#0}")

# Calculate range directory (0001-0100, 0101-0200, ...)
START=$(( ((10#$NUM - 1) / 100) * 100 + 1 ))
END=$(( START + 99 ))
RANGE_DIR=$(printf "%04d-%04d" $START $END)

# Find matching problem directory
PROBLEM_DIR=$(find "$LEETCODE_DIR/$RANGE_DIR" -maxdepth 1 -type d -name "${NUM}_*" 2>/dev/null | head -1)

if [ -z "$PROBLEM_DIR" ]; then
    echo "Error: Problem $NUM not found"
    echo "Search path: $LEETCODE_DIR/$RANGE_DIR/${NUM}_*"
    exit 1
fi

echo "Running test: $PROBLEM_DIR"
go test -v "$PROBLEM_DIR/..."
