#!/bin/bash
DB_URL="postgres://postgres:password@db:5432/contacthub?sslmode=disable"

for f in migrations/*.sql; do
    echo "Running $f"
    docker exec -i contacthub-db psql "$DB_URL" < "$f"
done
echo "All migrations done"
