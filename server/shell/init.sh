#!/usr/bin/env bash
set -euo pipefail

log() {
  printf '[%s] %s\n' "$(date '+%Y-%m-%d %H:%M:%S')" "$1"
}

log "Checking AWS configuration..."
./aws.sh

log "Checking Terraform CLI..."
./terraform.sh

# log "Validating Terraform state bucket..."
# ./bucket.sh "$1"

log "All prerequisite checks completed successfully."

log "Starting Client.."

cd "$(pwd)/../../client"

pnpm dev
