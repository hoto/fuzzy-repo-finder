all: clean build test

clean:
	go clean
	rm -rf bin/

build:
	go build -o bin/fuzzy-repo-finder cmd/fuzzy-repo-finder/main.go

test:
	go test -v ./...

run: clean build
	./bin/fuzzy-repo-finder

download:
	go get -v -t -d ./...
