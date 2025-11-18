# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview
This is a Go project called `fontgen`.
See @project.md for details.

## Development Commands

### Basic Go Operations
```bash
# Build the project
go build

# Run the main application
go run .

# Run tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run a specific test
go test -run TestName ./...

# Format code
go fmt ./...

# Vet code for potential issues
go vet ./...

# Download dependencies
go mod download

# Tidy up dependencies
go mod tidy

# Add a new dependency
go get package-name

# Update all dependencies
go get -u ./...
```

## Development Workflow
- Use `go mod tidy` after adding new dependencies
- Run `go fmt` before committing to maintain consistent formatting
- Use `go test` to run the test suite

## Notes
- Standard Go toolchain commands apply
