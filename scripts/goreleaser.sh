#!/usr/bin/env bash

set -euo pipefail

if [[ $# -ne 1 ]]; then
  echo "usage: scripts/goreleaser.sh snapshot|release" >&2
  exit 1
fi

mode=$1

case "${mode}" in
snapshot)
  goreleaser release --snapshot --clean
  ;;
release)
  goreleaser release --clean
  ;;
*)
  echo "usage: scripts/goreleaser.sh snapshot|release" >&2
  exit 1
  ;;
esac
