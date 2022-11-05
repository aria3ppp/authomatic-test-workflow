#!/bin/bash -e

echo "[`date`] Running entrypoint script..."

# export APP_DSN env to server
if [[ -z ${DSN} ]]; then
  export DSN="postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=disable"
fi

echo "[`date`] Running DB migrations..."
migrate -database "${DSN}" -path ./migrations up

echo "[`date`] Starting server..."
./server
