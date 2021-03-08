#!/usr/bin/env bash

# AVOID INVOKING THIS SCRIPT DIRECTLY -- USE `make lint`

set -euxo pipefail

golangci-lint --out-format=github-actions run \
  ./cmd/... \
  ./pkg/...
