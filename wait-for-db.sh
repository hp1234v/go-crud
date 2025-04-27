#!/bin/sh
# wait-for-db.sh

echo "Waiting for Postgres to be ready..."

until pg_isready -h db -p 5432; do
  sleep 1
done

echo "Postgres is ready. Starting the app..."
exec "$@"