all: clean build test

clean:
	go clean
	rm -rf bin/
	rm -rf dist/

build:
	go build -o bin/fuzzy-repo-finder cmd/fuzzy-repo-finder/main.go

test:
	go test -v ./...

run: clean build
	./bin/fuzzy-repo-finder $(arg)

download:
	go get -v -t -d ./...

install: clean build
	go install -v ./...
