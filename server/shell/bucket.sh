#!/usr/bin/env bash
set -euo pipefail

PROFILE="${AWS_PROFILE:-root}"
BUCKET="$1"

if aws --profile "$PROFILE" s3api head-bucket --bucket "$BUCKET" 2>/dev/null; then
  echo "Bucket ready for state!"
else
  echo "Bucket does not exist or access denied"
  exit 1
fi
