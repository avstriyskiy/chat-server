#!/bin/bash
source .env

export MIGRATION_DSN="host=$DB_HOST port=5432 dbname=$CHAT_SERVER_APP_DB user=$POSTGRES_USER password=$POSTGRES_PASSWORD sslmode=disable"

sleep 2 && goose -dir migrations postgres "${MIGRATION_DSN}" up -v
