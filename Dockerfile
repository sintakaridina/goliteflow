# Build stage
FROM golang:1.22-alpine AS builder

# Install git and ca-certificates (needed for go mod download)
RUN apk update && apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o goliteflow ./cmd/goliteflow

# Final stage
FROM alpine:latest

# Install ca-certificates and timezone data
RUN apk --no-cache add ca-certificates tzdata

# Create non-root user
RUN addgroup -g 1001 -S goliteflow && \
    adduser -u 1001 -S goliteflow -G goliteflow

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/goliteflow .

# Copy example configuration
COPY --from=builder /app/examples/lite-workflows.yml ./examples/

# Change ownership to non-root user
RUN chown -R goliteflow:goliteflow /app

# Switch to non-root user
USER goliteflow

# Expose port (if needed for future web interface)
EXPOSE 8080

# Default command
ENTRYPOINT ["./goliteflow"]
CMD ["--help"]
