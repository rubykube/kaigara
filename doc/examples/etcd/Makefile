NAME = etcd
VERSION = 0.0.1

.PHONY: build clean

build: kaigara
	docker build -t $(NAME):$(VERSION) --rm .
	docker tag $(NAME):$(VERSION) $(NAME):latest

kaigara:
	cp ../../../bin/kaigara .

clean:
	rm -f $(NAME)
	rm -f kaigara
	docker rmi -f $(NAME):$(VERSION)
