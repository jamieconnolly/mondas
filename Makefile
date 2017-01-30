DIRS := $(shell ls -d */ | grep -v vendor)
PACKAGES := $(shell glide novendor)
SOURCES := $(wildcard $(addsuffix *.go, $(DIRS)) *.go)

all: clean test

check: fmt vet lint

clean:
	@rm -f .coverprofile */.coverprofile gover.coverprofile

cover: clean test
	@gover .
	@go tool cover -html=gover.coverprofile -o coverage.html

fmt:
	@gofmt -s -l $(SOURCES) | awk '{print $$1 ": file is not formatted correctly"} END{if(NR>0) {exit 1}}' 2>&1; \
	if [ $$? -eq 1 ]; then \
		echo "!!! ERROR: Gofmt found unformatted files"; \
		exit 1; \
	fi

lint:
	@echo $(PACKAGES) | xargs -n 1 golint | awk '{print} END{if(NR>0) {exit 1}}' 2>&1; \
	if [ $$? -eq 1 ]; then \
	  echo "!!! ERROR: Golint found stylistic issues"; \
	  exit 1; \
	fi

test: check
	@go list -f '"go test -v -coverprofile={{.Dir}}/.coverprofile {{.ImportPath}}"' $(PACKAGES) | xargs -n 1 bash -c

vet:
	@go tool vet $(SOURCES) 2>&1; \
	if [ $$? -eq 1 ]; then \
		echo "!!! ERROR: Vet found suspicious constructs"; \
		exit 1; \
	fi

.PHONY: all check clean cover fmt lint test vet
