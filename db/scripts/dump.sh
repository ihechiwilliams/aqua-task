#!/bin/bash

pg_dump "${DATABASE_URL}" \
  --schema-only
