#!/usr/bin/env bash

set -euo pipefail

golangci-lint run
scripts/test.sh
