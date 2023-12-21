DOCKER ?= docker
DOCKERCOMPOSE ?= docker-compose

GO ?= go
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")
GOMODULES := $(shell go list ./...)

# run & build & clean

run:
	$(GO) run cmd/httpd/main.go

clean:
	@rm -r build

build:
	$(GO) build -o build/program/app cmd/httpd/main.go

# fmt

fmt:
	$(GOFMT) -w $(GOFILES)

# testing

test:
	$(GO) clean -testcache
	$(GO) mod tidy
	$(GO) test -cover $(GOMODULES)

update:
	$(GO) get -u ./...
	$(GO) mod tidy

info:
	@$(GO) vet $(GOMODULES)
	@$(GO) list $(GOMODULES)
	@$(GO) version

# deployment

deploy:
	$(DOCKER) build --tag todo-api .
	$(DOCKER) run -rm todo-api

tdeploy:
	$(DOCKER) build --tag todo-api .
	$(DOCKER) run todo-api

compose:
	$(DOCKERCOMPOSE) up -d

# phony

.PHONY: clean build run fmt test update info deploy tdeploy compose