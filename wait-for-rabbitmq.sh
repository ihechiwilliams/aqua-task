#!/bin/bash

set -e

host=${RABBITMQ_HOST:-rabbitmq}
port=${RABBITMQ_PORT:-5672}

echo "Waiting for RabbitMQ to be ready on ${host}:${port}..."

while ! nc -z "$host" "$port"; do
  sleep 1
done

echo "RabbitMQ is ready. Starting the application..."
exec "$@"
