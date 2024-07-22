#!/bin/bash

MIGRATION_DSN="host=${PG_HOST} port=${PG_PORT} dbname=${PG_DB} user=${PG_USER} password=${PG_PASSWORD} sslmode=${PG_SSL_MODE}"

sleep 2 && goose -dir "${MIGRATION_DIR}" postgres "$MIGRATION_DSN" up -v