
NAME=kaigara

TAG := $(shell git describe --tags --abbrev=0 2>/dev/null)

LDFLAGS += -X github.com/mod/kaigara/pkg/version.Version=${TAG}

all: clean test $(NAME)

.PHONY: test build release dist clean

$(NAME):
	go build -v -ldflags '$(LDFLAGS)'

build:
		GOBIN=$(BINDIR) $(GO) install $(GOFLAGS) -tags '$(TAGS)' k8s.io/helm/cmd/...

test:
	go test ./pkg/...

dist: clean
	mkdir -p ./bin
	mkdir -p dist/linux/amd64		&& CGO_ENABLED=0 GOOS=linux GOARCH=amd64	go build -ldflags '$(LDFLAGS)' -o dist/linux/amd64/$(NAME)
	mkdir -p dist/linux/i386		&& CGO_ENABLED=0 GOOS=linux GOARCH=386		go build -ldflags '$(LDFLAGS)' -o dist/linux/i386/$(NAME)
	mkdir -p dist/darwin/amd64	&& CGO_ENABLED=0 GOOS=darwin GOARCH=amd64	go build -ldflags '$(LDFLAGS)' -o dist/darwin/amd64/$(NAME)
	mkdir -p dist/darwin/i386		&& CGO_ENABLED=0 GOOS=darwin GOARCH=386		go build -ldflags '$(LDFLAGS)' -o dist/darwin/i386/$(NAME)
	ln -sf ../dist/linux/amd64/$(NAME) ./bin/

release: dist
	mkdir releases
	tar -cvzf releases/$(NAME)-linux-amd64-$(TAG).tar.gz -C dist/linux/amd64 $(NAME)
	tar -cvzf releases/$(NAME)-linux-i386-$(TAG).tar.gz -C dist/linux/i386 $(NAME)
	tar -cvzf releases/$(NAME)-darwin-amd64-$(TAG).tar.gz -C dist/darwin/amd64 $(NAME)
	tar -cvzf releases/$(NAME)-darwin-i386-$(TAG).tar.gz -C dist/darwin/i386 $(NAME)

clean:
	rm -rf $(NAME)
	rm -rf dist
	rm -rf releases
