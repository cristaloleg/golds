.PHONY: all build bench test fmt race cpuprof memprof 

TIMESTAMP=$(shell date +"%Y-%m-%d_%H-%M-%S")

all: install build test

install:
	go get github.com/golang/lint/golint
	go get github.com/mattn/goveralls
	go get golang.org/x/tools/cmd/cover

build:
	go build .

bench:
	go test -bench=. -benchmem

test:
	go test -v -cover ./...

fmt:
	gofmt -l -w *.go

lint:
	golint ./...
	go vet ./...

race:
	go test -race ./...

cpuprof:
	go test -cpuprofile cpu-${TIMESTAMP}.prof

memprof:
	go test -memprofile mem-${TIMESTAMP}.prof