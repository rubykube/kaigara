NAME = uhttpd
VERSION = 0.0.1

.PHONY: build clean

build: uhttpd
	docker build -t $(NAME):$(VERSION) --rm .
	docker tag $(NAME):$(VERSION) $(NAME):latest

uhttpd: uhttpd.go
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build uhttpd.go

clean:
	rm -f $(NAME)
	docker rmi -f $(NAME):$(VERSION)
