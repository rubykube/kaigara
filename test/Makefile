NAME = kaigara/test
VERSION = 0.0.1

.PHONY: build test clean

build: kaigara
	docker build -t $(NAME):$(VERSION) --rm .
	docker tag $(NAME):$(VERSION) $(NAME):latest

test: clean build
	./scripts/start.sh

kaigara:
	cp ../kaigara .

clean:
	rm -f kaigara
	-docker rmi -f $(NAME):$(VERSION)
