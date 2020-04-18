#!/bin/sh
set -euo pipefail

echo "Waiting for local db to be available"

timeout 60 sh -c '
    while ! mysqladmin ping -h ${MYSQL_HOST_URL} -u ${MASTER_USERNAME} -p${MASTER_PASSWORD} --silent; do
        echo "Failed to connect to Mysql at ${MYSQL_HOST_URL}, trying again in 3 seconds..."
        sleep 3
    done
    echo "MYSQL is running."
'
go test ./... -cover -failfast