.PHONY: build run test clean

build:
	go build -o bin/ssh_tunnel main.go

run: build
	./bin/ssh_tunnel

test:
	go test ./...

clean:
	rm -rf bin
