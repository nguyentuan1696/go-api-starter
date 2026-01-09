# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o anawim-backend ./cmd/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Install necessary runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Set timezone to Asia/Ho_Chi_Minh
ENV TZ=Asia/Ho_Chi_Minh

# Copy the binary from builder
COPY --from=builder /app/anawim-backend .

# Expose the application port (adjust if your config uses a different port)
EXPOSE 8080

# Command to run the application
CMD ["./anawim-backend", "serve"]
