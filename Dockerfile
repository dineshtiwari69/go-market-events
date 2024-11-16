# Build stage
FROM golang:1.21.1 AS builder

WORKDIR /app

# Set environment variables to bypass proxy for specific packages
ENV GOPRIVATE=github.com/consensys/*

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN make fmt && make lint && make test && CGO_ENABLED=0 GOOS=linux go build -o /app/market-events

# Final stage
FROM alpine:3.18

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/market-events .

# Run the application
CMD ["./market-events"]