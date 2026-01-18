.PHONY: all build test bench lint generate clean

all: generate build test

build:
	go build ./...

test:
	go test -v -race ./...

bench:
	go test -bench=. -benchmem ./...

lint:
	golangci-lint run ./...

generate:
	go run scripts/generate_data.go

clean:
	rm -f pkg/useragent/realdata.go
