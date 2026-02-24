#!/bin/bash

program=picsum
source=./cmd/picsum

# shellcheck disable=SC1091
. ./release.env

build_date=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
git_commit=$(git rev-parse HEAD)

build(){
  local GOOS=$1
  local GOARCH=$2
  local extension=$3
  echo "Building for $GOOS/$GOARCH"
  go build \
  -ldflags "-X github.com/siakhooi/picsum/internal/versioninfo.Version=$RELEASE_VERSION -X github.com/siakhooi/picsum/internal/versioninfo.Date=$build_date -X github.com/siakhooi/picsum/internal/versioninfo.Commit=$git_commit" \
  -o bin/"${program}-${GOOS}-${GOARCH}${extension}" $source
}

build linux amd64 ""
build linux arm64 ""
build windows amd64 ".exe"
build windows 386 ".exe"
build darwin amd64 ""
build darwin arm64 ""
build freebsd amd64 ""
build freebsd arm64 ""
build netbsd amd64 ""
build netbsd arm64 ""
build openbsd amd64 ""
build openbsd arm64 ""
