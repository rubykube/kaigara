NAME = kaigara/web
VERSION = 0.0.3

.PHONY: build clean

build:
	git submodule update --remote
	docker build -t $(NAME):$(VERSION) --rm .
	docker tag $(NAME):$(VERSION) $(NAME):latest

release: build
	docker push $(NAME)

deploy: release
	kubectl apply -f kube

clean:
	docker rmi -f $(NAME):$(VERSION)
