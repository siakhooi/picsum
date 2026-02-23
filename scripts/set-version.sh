#!/bin/bash
set -e

# shellcheck disable=SC1091
. ./release.env

sed -i 'internal/versioninfo/versioninfo.go' -e 's@const picSumVersion = ".*"@const picSumVersion = "'"$RELEASE_VERSION"'"@g'
sed -i 'internal/versioninfo/versioninfo_test.go' -e 's@const expectedPicsumVersion = ".*"@const expectedPicsumVersion = "'"$RELEASE_VERSION"'"@g'
