NAME = kaigara/web
VERSION = 0.0.4

.PHONY: build clean

build:	module
	docker build -t $(NAME):$(VERSION) --rm .
	docker tag $(NAME):$(VERSION) $(NAME):latest

release: build
	docker push $(NAME)

deploy: release
	kubectl apply -f kube

module:
	git submodule init
	git submodule update --remote

clean:
	docker rmi -f $(NAME):$(VERSION)
