NAME = uhttpd
VERSION = 0.0.1

.PHONY: build clean

build: kaigara uhttpd
	docker build -t $(NAME):$(VERSION) --rm .
	docker tag $(NAME):$(VERSION) $(NAME):latest

uhttpd:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build uhttpd.go

kaigara:
	cp ../../../kaigara .

clean:
	rm -f $(NAME)
	rm -f kaigara
	docker rmi -f $(NAME):$(VERSION)
