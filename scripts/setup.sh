#!/bin/sh

if [ ! -f ".env" ]; then
	cp ".env.example" ".env";
fi

if ! command -v go >/dev/null 2>&1; then
	go mod tidy
	go mod download
fi

if ! command -v air >/dev/null 2>&1; then
	go install github.com/air-verse/air@latest
fi
