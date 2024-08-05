#!/bin/sh

# Setup environment variables
if [ ! -f ".env" ]; then
	cp ".env.example" ".env";
fi

# Download Go modules
if ! command -v go >/dev/null 2>&1; then
	go mod download
fi

# Install Air
if ! command -v air >/dev/null 2>&1; then
	go install github.com/air-verse/air@latest
fi
