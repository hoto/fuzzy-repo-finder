.PHONY: clean dependencies build test vet run install goreleaser-release goreleaser-dry-run snap-build snap-list snap-install snap-remove

REPO_NAME    = github.com/hoto/fuzzy-repo-finder
GORELEASER   = go run github.com/goreleaser/goreleaser/v2@latest
VERSION      ?= 0.0.0
SHORT_COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)
BUILD_DATE   ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS      = -X '$(REPO_NAME)/pkg/config.Version=$(VERSION)' \
               -X '$(REPO_NAME)/pkg/config.ShortCommit=$(SHORT_COMMIT)' \
               -X '$(REPO_NAME)/pkg/config.BuildDate=$(BUILD_DATE)'

clean:
	go clean
	rm -rf bin/ dist/ *.snap

dependencies:
	go mod download
	go mod tidy
	go mod verify

build: dependencies
	go build -trimpath -ldflags="$(LDFLAGS)" -o bin/fuzzy-repo-finder ./cmd/fuzzy-repo-finder

test:
	go test -v ./...

vet:
	go vet ./...

run: clean build
	./bin/fuzzy-repo-finder $(arg)

install: clean build
	go install -v ./...

goreleaser-release: clean dependencies
	$(GORELEASER) release --clean

goreleaser-dry-run: clean dependencies
	$(GORELEASER) release --snapshot --skip=publish --clean

snap-build:
	snapcraft

snap-list:
	unsquashfs -l *.snap
	snap list

snap-install:
	sudo snap install --dangerous *.snap

snap-remove:
	sudo snap remove fuzzy-repo-finder
