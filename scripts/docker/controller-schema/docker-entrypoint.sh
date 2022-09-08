#!/bin/sh

set -e

POSTGRES_HOSTPORT="${POSTGRES_DSN##*@}"
POSTGRES_HOSTPORT="${POSTGRES_HOSTPORT%%/*}"

POSTGRES_HOST="${POSTGRES_HOSTPORT%%:*}"
POSTGRES_PORT="${POSTGRES_HOSTPORT##*:}"

until nc -z -v -w30 "$POSTGRES_HOST" "$POSTGRES_PORT"
do
  echo "Waiting for database connection..."
  # wait for 5 seconds before check again
  sleep 5
done

migrate -source file:///migrations -database "$POSTGRES_DSN" "$@"
