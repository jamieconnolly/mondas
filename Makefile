DIRS := $(shell ls -d */ | grep -v vendor)
PACKAGES := $(shell glide novendor)
SOURCES := $(wildcard $(addsuffix *.go, $(DIRS)) *.go)

all: test

check: fmt vet lint

fmt:
	@gofmt -s -l $(SOURCES) | awk '{print $$1 ": file is not formatted correctly"} END{if(NR>0) {exit 1}}' 2>&1; \
	if [ $$? -eq 1 ]; then \
		echo "!!! ERROR: Gofmt found unformatted files"; \
		exit 1; \
	fi

lint:
	@echo $(PACKAGES) | xargs -n 1 golint

test: check
	@go test -v $(PACKAGES)

vet:
	@go tool vet $(SOURCES) 2>&1; \
	if [ $$? -eq 1 ]; then \
		echo "!!! ERROR: Vet found suspicious constructs"; \
		exit 1; \
	fi

.PHONY: all check fmt lint test vet
