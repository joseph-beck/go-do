DOCKER ?= docker
GO ?= go
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")
GOMODULES := $(shell go list ./...)

# clean & build

clean:
	@rm -r build

build: 
	$(GO) build -o build/program/app cmd/httpd/main.go 

# run & dev

run:
	make build
	./build/program/app

dev: 
	$(GO) run cmd/httpd/main.go

# fmt

fmt:
	$(GOFMT) -w $(GOFILES)

# testing

test:
	$(GO) clean -testcache 
	$(GO) mod tidy
	$(GO) test -cover $(GOMODULES)

# deployment

deploy:
	$(DOCKER) build --tag todo-api .
	$(DOCKER) run -rm todo-api

tdeploy:
	$(DOCKER) build --tag todo-api .
	$(DOCKER) run todo-api

# phony

.PHONY: clean build run dev fmt test deploy tdeploy