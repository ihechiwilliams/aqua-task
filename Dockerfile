# Use the official Go image
FROM golang:1.23.4-alpine3.21 AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux

# Precompile the standard library for faster builds
RUN go install -v -installsuffix cgo -a std

RUN apk update && apk add git make --no-cache

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.0
# Verify installation
RUN ls -l /go/bin/migrate

# Set the working directory
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN go build -o main ./cmd/server/*.go

# Use a minimal image for running the application
FROM alpine:3.21


# Install necessary tools (e.g., CA certificates)
RUN apk add --no-cache ca-certificates bash

COPY --from=builder $GOROOT/go/bin/migrate /usr/local/bin/
COPY --from=builder /usr/bin/make /usr/local/bin/

# Set working directory
WORKDIR /app

# Copy the compiled binary from the builder
COPY --from=builder /app/Makefile ./

COPY --from=builder /app/db/migrations /app/db/migrations
COPY --from=builder /app/db/scripts/migrate.sh /app/db/scripts/migrate.sh
RUN chmod +x /app/db/scripts/migrate.sh
COPY --from=builder /app/wait-for-rabbitmq.sh /usr/local/bin/


COPY --from=builder /app/main .

# Expose the application port
EXPOSE 8080

# Start the application
CMD ["./main"]