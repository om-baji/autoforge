#!/usr/bin/env bash

set -euo pipefail

log() {
  printf '[%s] %s\n' "$(date '+%Y-%m-%d %H:%M:%S')" "$1"
}

tf=$(terraform --version | grep Terraform | wc -l)

if [ $tf -eq 1 ]; then
    log "Terraform Ready!"
else {
    log "Terraform CLI not installed!"
    exit 1
}
fi
