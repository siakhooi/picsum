#!/usr/bin/env bash

set -euo pipefail

# shellcheck disable=SC1091
. ./release.env

go_tag="v${RELEASE_VERSION}"
set -x
gh release create "${go_tag}" --title "${RELEASE_TITLE}" --notes "${RELEASE_NOTE}" --latest
