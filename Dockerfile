# Build stage
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o receipt-processor-api ./cmd/server

# Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/receipt-processor-api .
CMD ["./receipt-processor-api"]
