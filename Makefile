GOFILES=$(shell find . -type f -name '*.go' | sort)

.PHONY: all
all: deps fmt vet lint

.PHONY: deps
deps:
	go get github.com/golang/lint/golint
	go get github.com/gorilla/mux
	go get github.com/mattn/go-sqlite3
	go get github.com/stretchr/testify/assert

.PHONY: lint
lint:
	$(GOPATH)/bin/golint ${GOFILES}

.PHONY: fmt
fmt:
	@if [ -n "$$(gofmt -l ${GOFILES})" ]; then echo 'Please run gofmt -l -w on your code.' && exit 1; fi

.PHONY: gofmt
gofmt:
	gofmt -l -w ${GOFILES}

.PHONY: vet
vet:
	go tool vet -composites=false ${GOFILES}

.PHONY: test
test: vet fmt lint
	go test -v -covermode=count -coverprofile=cover.out github.com/reaxun/multidex/api

.PHONY: coverage
coverage:
	go tool cover -html=cover.out -o=cover.html
