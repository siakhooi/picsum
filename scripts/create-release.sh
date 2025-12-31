#!/bin/bash
set -e

# shellcheck disable=SC1091
. ./release.env

RELEASE_NOTE="$RELEASE_TITLE"

go_tag="v${RELEASE_VERSION}"
set -x
gh release create "$go_tag" --title "$RELEASE_TITLE" --notes "${RELEASE_NOTE}" --latest
