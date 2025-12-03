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
	bin/picsum-linux-amd64 -h
	bin/picsum-linux-amd64 200
	bin/picsum-linux-amd64 200 300
	bin/picsum-linux-amd64 -g 200 300
	bin/picsum-linux-amd64 -i 237 200
	bin/picsum-linux-amd64 -i 237 200 300
	bin/picsum-linux-amd64 -g -i 237 200
	bin/picsum-linux-amd64 -s picsumabc 200
	bin/picsum-linux-amd64 -s hellohello 200 300
	bin/picsum-linux-amd64 -g -s hellohello 200 300
run-i:
	bin/picsum-linux-amd64
	bin/picsum-linux-amd64 200 300 400
	bin/picsum-linux-amd64 -i 237 -s hellohello 200 300
