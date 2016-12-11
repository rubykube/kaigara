.PHONY: build

NAME=kaigara
VERSION=0.0.2
TAG=v$(VERSION)

all: clean test $(NAME)

$(NAME):
	echo "Building $(NAME)"
	go build -v

.PHONY: test
test:
	go test ./...

dist: dist-clean
	mkdir -p dist/linux/amd64		&& CGO_ENABLED=0 GOOS=linux GOARCH=amd64	go build -o dist/linux/amd64/$(NAME)
	mkdir -p dist/linux/i386		&& CGO_ENABLED=0 GOOS=linux GOARCH=386		go build -o dist/linux/i386/$(NAME)
	mkdir -p dist/darwin/amd64	&& CGO_ENABLED=0 GOOS=darwin GOARCH=amd64	go build -o dist/darwin/amd64/$(NAME)
	mkdir -p dist/darwin/i386		&& CGO_ENABLED=0 GOOS=darwin GOARCH=386		go build -o dist/darwin/i386/$(NAME)

default:
	mkdir -p ./bin
	ln -s ../dist/linux/amd64/$(NAME) ./bin/

release: dist
	mkdir releases
	tar -cvzf releases/$(NAME)-linux-amd64-$(TAG).tar.gz -C dist/linux/amd64 $(NAME)
	tar -cvzf releases/$(NAME)-linux-i386-$(TAG).tar.gz -C dist/linux/i386 $(NAME)
	tar -cvzf releases/$(NAME)-darwin-amd64-$(TAG).tar.gz -C dist/darwin/amd64 $(NAME)
	tar -cvzf releases/$(NAME)-darwin-i386-$(TAG).tar.gz -C dist/darwin/i386 $(NAME)

dist-clean:
	rm -rf dist
	rm -rf releases

clean:
	rm -rf $(NAME)
