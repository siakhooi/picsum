# List available recipes
default:
    @just --list

# Remove build outputs and downloaded images
clean:
    rm -rf bin dist *.jpg

# Cross-compile binaries and build a GoReleaser snapshot
build:
    scripts/build.sh
    goreleaser build --snapshot --clean

# Build a local GoReleaser snapshot release
go-release:
    goreleaser release --snapshot --clean

# Run tests and write coverage reports
test:
    scripts/test.sh

# Run golangci-lint
golangci-lint:
    golangci-lint run

# Clean, test, lint, and build
all: clean test golangci-lint build

# Create a GitHub release from release.env
release:
    scripts/create-release.sh

# Watch the current GitHub Actions run
commit-watch:
    gh run watch

# Create a GitHub release and watch the Actions run
release-watch: release
    gh run watch

picsum := "bin/picsum-linux-amd64"

# Show CLI help
run-help:
    {{ picsum }} -h

# Print build info
run-build:
    {{ picsum }} --build

# Exercise the common CLI paths
run:
    {{ picsum }} 200
    {{ picsum }} -q 200 300
    {{ picsum }} -g 200 300
    {{ picsum }} -i 237 200
    {{ picsum }} -i 237 200 300
    {{ picsum }} -g -i 237 200
    {{ picsum }} -s picsumabc 200
    {{ picsum }} -s 'hello% hello' 200 300
    {{ picsum }} -g -s hellohello 200 300
    {{ picsum }} -b -s hellohello 200 300
    {{ picsum }} -q -g -b -s hellohello -o hello.jpg 200 300
    {{ picsum }} -q -g -b -s hellohello -o hello.jpg 200 300
    {{ picsum }} -q -g -b -s hellohello -o hello.jpg -f 200 300

# Exercise invalid CLI paths
run-i:
    {{ picsum }}
    {{ picsum }} 200 300 400
    {{ picsum }} -i 237 -s hellohello 200 300
