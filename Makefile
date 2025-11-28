clean:
	rm -rf bin dist  *.jpg
build:
	scripts/build.sh
	goreleaser build --snapshot --clean
test:
	scripts/test.sh
golangci-lint:
	golangci-lint run

set-version:
	scripts/set-version.sh
all: clean set-version test golangci-lint build

commit:
	scripts/git-commit-and-push.sh

release:
	scripts/create-release.sh

commit-watch: commit
	gh run watch

release-watch: release
	gh run watch

run:
	bin/picsum-linux-amd64 200
	bin/picsum-linux-amd64 200 300
