GOARCH ?= $(shell go env GOARCH)
GOOS ?= $(shell go env GOOS)

bin/$(NAME):
	@echo "==> Building $@…"
	@mkdir -p bin
	@GOARCH=$(GOARCH) GOOS=$(GOOS) go build -o "$@"

.PHONY: build
build: clean bin/$(NAME)
