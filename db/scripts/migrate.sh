#!/bin/bash
echo "Running migration on: ${DATABASE_URL}"
migrate -path ./db/migrations \
  -database "${DATABASE_URL}"\
  up