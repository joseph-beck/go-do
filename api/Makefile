# clean & build

.PHONY: clean build

clean:
	@rm -r build

build: 
	go build -o build/program/app cmd/httpd/main.go 

# run & dev

.PHONY: run dev

run:
	make build
	./build/program/app

dev: 
	go run cmd/httpd/main.go

# testing

.PHONY: test

test:
	go clean -testcache 
	go mod tidy
	go test -cover -v ./...

# deployment

.PHONY: deploy tdeploy

deploy:
	docker build --tag todo-api .
	docker run -rm todo-api

deploy:
	docker build --tag todo-api .
	docker run todo-api