all: clean build test

clean:
	go clean
	rm -rf bin/ dist/ *.snap

dependencies:
	go get -v -t -d ./...
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

build: dependencies
	go build -o bin/fuzzy-repo-finder cmd/fuzzy-repo-finder/main.go

test:
	go test -v ./...

run: clean build
	./bin/fuzzy-repo-finder $(arg)

install: clean build
	go install -v ./...

gituh-release: dependencies
	curl -sL https://git.io/goreleaser | bash

github-release-dry-run: dependencies
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
