# Receipt Processor Api

## Overview

This is a web service that processes receipts and calculates points based on given rules.

## Running the Application

### Locally

1. Ensure Go is installed.
2. Run `go run ./cmd/server` to start the application on `localhost:8080`.

### Using Docker

1. Build the Docker image:
   ```bash
   docker build -t receipt-processor .
   docker run -p 8080:8080 receipt-processor