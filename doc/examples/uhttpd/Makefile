NAME = uhttpd
VERSION = 0.0.1

.PHONY: all build clean

all: build

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build uhttpd.go
	docker build -t $(NAME):$(VERSION) --rm .
	docker tag $(NAME):$(VERSION) $(NAME):latest

clean:
	rm -f $(NAME)
	docker rmi -f $(NAME):$(VERSION)
