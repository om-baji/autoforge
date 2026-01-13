#!/usr/bin/env bash

set -euo pipefail

tf=$(terraform --version | grep Terraform | wc -l)

if [ $tf -eq 1 ]; then
    echo -e "Terraform Ready!"
else {
    echo -e "Terraform CLI not installed!"
    exit 1
}
fi
