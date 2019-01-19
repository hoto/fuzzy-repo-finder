all: clean build test

clean:
	go clean
	rm -rf bin/
	rm -rf dist/

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

release: dependencies
	curl -sL https://git.io/goreleaser | bash

release_dry_run: dependencies
	goreleaser release --skip-publish --snapshot --rm-dist
