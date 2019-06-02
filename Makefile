.PHONY: mod vendor build build-linux dep


build:
	go build .

build-linux:
	GOOS=linux GOARCH=amd64 go build .

dep:
	go mod tidy
	go mod vendor
