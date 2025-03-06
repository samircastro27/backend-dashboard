#!/bin/bash

# Entrypoint for run in container and load
# connections from env var and replace values into the components yaml files.

# NOT USE LOCALLY

check_env_vars() {
  local vars=("POSTGRES_URI")
  
  for var in "${vars[@]}"; do
    if [ -z "${!var}" ]; then
      echo "Error: $var is required"
      exit 1
    fi
  done
}

# check env vars
check_env_vars

# create variables for run with sed command
ESCAPED_POSTGRES_URI=$(echo "$POSTGRES_URI" | sed 's/[\/&]/\\&/g')


# replaces values in files
sed -i "s/POSTGRES_URI/${ESCAPED_POSTGRES_URI}/g" components/postgres.yaml

# Ejecutar migraciones antes de iniciar Dapr
go run ./cmd/migrations/migration.go

# run dapr services
dapr run -f ./dapr.yaml
