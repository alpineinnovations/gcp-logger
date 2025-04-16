# GCP Logger

A Go logging library that provides structured logging in Google Cloud Platform (GCP) compatible format, along with HTTP
middleware for request/response logging.

## Features

- GCP-compatible structured logging format
- HTTP middleware for automatic request/response logging
- Context-based logger management
- Support for different log levels (DEBUG, INFO, WARNING, ERROR)
- Request with latency measurements
- JSON-based logging output

## Installation

For our private repository, we need to add the following line to our git config:

```bash
git config --global url."git@github.com:".insteadOf "https://github.com/"
```

Set `GOPRIVATE` Environment Variable

```bash
export GOPRIVATE=github.com/alpineinnovations/*
```

```bash
go get github.com/alpineinnovations/gcp-logger
```

## Usage

### Basic Logging

```go
import (
"github.com/alpineinnovations/gcp-logger/logger"
)

// Create a new GCP handler with INFO level
handler := logger.NewGCPHandler("INFO")
```

### HTTP Middleware

```go
import (
"github.com/alpineinnovations/gcp-logger/web/middlewares"
)

// Add the logging middleware to your HTTP server
http.Handle(url, middlewares.LogMiddleware(...))
```

The middleware automatically logs:

- Incoming requests with method, URL, IP, user agent, and other HTTP details
- Outgoing responses with status code and latency
- All logs are formatted in GCP-compatible JSON structure

### Context-Based Logging

```go
// Add logger to context with user ID
ctx = logger.LoggerContext(ctx, userID)

// Get logger from context
logger := logger.FromCtx(ctx)
logger.Info("Your log message", "key", "value")
```

## Log Format

The logger outputs JSON-structured logs compatible with GCP logging format, including:

- Severity level
- Timestamp
- Source location
- Message
- HTTP request details (when using middleware)
- Custom attributes

## Requirements

- Go 1.24.1 or higher
