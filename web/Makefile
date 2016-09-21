NAME = kaiweb
VERSION = 0.0.1

.PHONY: build clean

build:
	docker build -t $(NAME):$(VERSION) --rm .
	docker tag $(NAME):$(VERSION) $(NAME):latest

clean:
	docker rmi -f $(NAME):$(VERSION)
