SERVICE  = kaigara
NAME    ?= rubykube/$(SERVICE)
VERSION ?= $(shell cat VERSION)

.PHONY: build clean

build: docs module
	docker build -t $(NAME):$(VERSION) --rm .
	docker tag $(NAME):$(VERSION) $(NAME):latest

release: build
	docker push $(NAME):latest
	docker push $(NAME):$(VERSION)

deploy:
	helm install --name $(SERVICE) ./chart --set="image.tag=$(VERSION)"

module:
	git submodule init
	git submodule update --remote

clean:
	docker rmi -f $(NAME):$(VERSION)

docs:
	bin/docsgen.sh
