#!/usr/bin/env bash
set -euo pipefail

PROFILE="${AWS_PROFILE:-default}"

command -v aws >/dev/null 2>&1 || {
  echo -e "AWS CLI not installed"
  exit 1
}

aws configure list-profiles | grep -qx "$PROFILE" || {
  echo -e "AWS profile '$PROFILE' not configured"
  exit 1
}

aws sts get-caller-identity >/dev/null 2>&1 || {
  echo -e "AWS credentials invalid or expired"
  exit 1
}

echo -e "AWS ready (profile: $PROFILE)"
