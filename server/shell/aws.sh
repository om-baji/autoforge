#!/usr/bin/env bash
set -euo pipefail

log() {
  printf '[%s] %s\n' "$(date '+%Y-%m-%d %H:%M:%S')" "$1"
}

PROFILE="${AWS_PROFILE:-default}"

command -v aws >/dev/null 2>&1 || {
  log "AWS CLI not installed"
  exit 1
}

aws configure list-profiles | grep -qx "$PROFILE" || {
  log "AWS profile '$PROFILE' not configured"
  exit 1
}

aws sts get-caller-identity >/dev/null 2>&1 || {
  log "AWS credentials invalid or expired"
  exit 1
}

log "AWS ready (profile: $PROFILE)"
