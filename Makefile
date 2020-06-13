.PHONY: clean dependencies build test run install github-release github-release-dry-run goreleaser-dry-run snap-build snap-list snap-install snap-remove

REPO_NAME = github.com/hoto/fuzzy-repo-finder

clean:
	go clean
	rm -rf bin/ dist/ *.snap

dependencies:
	go mod download
	go mod tidy
	go mod verify

build: dependencies
	go build -ldflags="-X '${REPO_NAME}/pkg/config.Version=0.0.0' -X '${REPO_NAME}/pkg/config.ShortCommit=HASH' -X '${REPO_NAME}/pkg/config.BuildDate=DATE'" -o bin/fuzzy-repo-finder ./cmd/fuzzy-repo-finder/main.go

test:
	go test -v ./...

run: clean build
	./bin/fuzzy-repo-finder $(arg)

install: clean build
	go install -v ./...

github-release: clean dependencies
	curl -sL https://git.io/goreleaser | VERSION=v0.137.0 bash

github-release-dry-run: clean dependencies
	curl -sL https://git.io/goreleaser | VERSION=v0.137.0 bash -s -- --skip-publish --snapshot --rm-dist

goreleaser-dry-run: dependencies
	goreleaser release --skip-publish --snapshot --rm-dist

snap-build:
	snapcraft

snap-list:
	unsquashfs -l *.snap
	snap list

snap-install:
	sudo snap install --dangerous *.snap

snap-remove:
	sudo snap remove fuzzy-repo-finder
