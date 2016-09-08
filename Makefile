.PHONY: build

BIN_NAME=kaigara
BIN_VERSION=0.0.1

build: dist
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o ./bin/linux/amd64/$(BIN_NAME)
	ln -s linux/amd64/$(BIN_NAME) ./bin/$(BIN_NAME)

dist: dist-clean
	mkdir -p bin/linux/amd64

dist-clean:
	rm -rf bin/*
