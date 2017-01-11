PACKAGES := $(shell glide novendor)

all: test

test:
	@go test -v ${PACKAGES}

.PHONY: all test
