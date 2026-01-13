#!/usr/bin/env bash
set -euo pipefail

log() {
  printf '[%s] %s\n' "$(date '+%Y-%m-%d %H:%M:%S')" "$1"
}

PROFILE="${AWS_PROFILE:-root}"
BUCKET="$1"

if aws --profile "$PROFILE" s3api head-bucket --bucket "$BUCKET" 2>/dev/null; then
  log "Bucket ready for state!"
else
  log "Bucket does not exist or access denied"
  exit 1
fi
