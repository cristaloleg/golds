.PHONY: all build bench test fmt race cpuprof memprof 

TIMESTAMP=$(shell date +"%Y-%m-%d_%H-%M-%S")

all: test install

build:
	go build .

bench:
	go test -bench=. -benchmem

test:
	go test -cover ./...

fmt:
	gofmt -l -w *.go

race:
	go test -race ./...

cpuprof:
	go test -cpuprofile cpu-${TIMESTAMP}.prof

memprof:
	go test -memprofile mem-${TIMESTAMP}.prof