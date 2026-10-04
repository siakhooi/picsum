#!/usr/bin/env bash

set -euo pipefail

program=picsum
cmd=./cmd/picsum

# shellcheck disable=SC1091
. ./release.env

build_date=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
git_commit=$(git rev-parse HEAD)

build() {
  local goos=$1
  local goarch=$2
  local extension=${3:-}
  echo "Building for ${goos}/${goarch}"
  CGO_ENABLED=0 GOOS="${goos}" GOARCH="${goarch}" go build \
    -ldflags "-X github.com/siakhooi/picsum/internal/versioninfo.Version=${RELEASE_VERSION} -X github.com/siakhooi/picsum/internal/versioninfo.Date=${build_date} -X github.com/siakhooi/picsum/internal/versioninfo.Commit=${git_commit}" \
    -o "bin/${program}-${goos}-${goarch}${extension}" "${cmd}"
}

while read -r goos goarch extension; do
  build "${goos}" "${goarch}" "${extension}"
done <<'EOF'
linux amd64
linux arm64
windows amd64 .exe
windows 386 .exe
darwin amd64
darwin arm64
freebsd amd64
freebsd arm64
netbsd amd64
netbsd arm64
openbsd amd64
openbsd arm64
EOF
