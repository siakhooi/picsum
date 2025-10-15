#!/bin/bash
set -e

# shellcheck disable=SC1091
. ./release.env

sed -i 'internal/version/version.go' -e 's@const picSumVersion = ".*"@const picSumVersion = "'"$RELEASE_VERSION"'"@g'
sed -i 'internal/version/version_test.go' -e 's@const expectedPicsumVersion = ".*"@const expectedPicsumVersion = "'"$RELEASE_VERSION"'"@g'
