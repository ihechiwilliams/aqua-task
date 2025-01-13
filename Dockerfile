# Stage 1: Builder
FROM golang:1.23.4-alpine3.21 AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux

# Precompile the standard library for faster builds
RUN go install -v -installsuffix cgo -a std

# Install necessary tools
RUN apk update && apk add --no-cache git make

# Install migrate for database migrations
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.0
RUN ls -l /go/bin/migrate  # Verify installation

# Set the working directory
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application code
COPY . .

# Build the binaries
RUN go build -o main ./cmd/server/*.go
RUN go build -o notification ./cmd/notification/*.go


# Stage 2: Runtime Image
FROM alpine:3.21 AS runtime

# Install necessary tools (e.g., CA certificates)
RUN apk add --no-cache ca-certificates bash

# Copy binaries and scripts from builder
COPY --from=builder $GOROOT/go/bin/migrate /usr/local/bin/
COPY --from=builder /usr/bin/make /usr/local/bin/
COPY --from=builder /app/wait-for-rabbitmq.sh /usr/local/bin/
COPY --from=builder /app/db/migrations /app/db/migrations
COPY --from=builder /app/db/scripts/migrate.sh /app/db/scripts/migrate.sh
COPY --from=builder /app/main /usr/local/bin/main
COPY --from=builder /app/notification /usr/local/bin/notification

# Grant execution permissions for scripts
RUN chmod +x /usr/local/bin/wait-for-rabbitmq.sh
RUN chmod +x /app/db/scripts/migrate.sh

# Set working directory
WORKDIR /app

# Expose application ports
EXPOSE 8080 9090 50051

# Set entrypoint to wait for RabbitMQ and run the binary
ENTRYPOINT ["/usr/local/bin/wait-for-rabbitmq.sh"]
CMD ["main"]
