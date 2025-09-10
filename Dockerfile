# =============================================================================
# Build stage
# =============================================================================
FROM golang:1.24.4-alpine AS builder

# Install git and ca-certificates (needed for fetching dependencies and HTTPS)
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# =============================================================================
# Final stage
# =============================================================================
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates tzdata

# Create a non-root user (có thể override UID/GID khi build)
ARG APP_UID=1000
ARG APP_GID=1000
RUN addgroup -g ${APP_GID} -S appgroup && \
    adduser  -u ${APP_UID} -S appuser -G appgroup

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/main .

# Tạo thư mục log chuẩn tại /var/log/go-api-starter
RUN mkdir -p /app/logs \
    && chown -R ${APP_UID}:${APP_GID} /app/logs \
    && chmod -R 775 /app/logs

USER appuser

# Expose port
EXPOSE 7070

# Environment variables
ENV DOCKER_CONTAINER=true

# Run the binary
CMD ["./main"]