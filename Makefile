#!make

include .env
export

MAKEFLAGS += --always-make

PACKAGES=./pkg/...
APP_BASE_IMAGE=golang:1.15-alpine
APP_ENV_IMAGE=docker.pkg.github.com/egnd/go-tghandler/golang:1.15-alpine-env

.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

%:
	@:

########################################################################################################################

lint: ## Lint source code
	golangci-lint run --color=always --config=.golangci.yml $(PACKAGES)

docker-lint:
	docker run --rm -t --volume $$(pwd):/src:rw $(APP_ENV_IMAGE) make lint
	@echo "Success"

test: mocks ## Test source code
	go test -mod=readonly -cover -covermode=count -coverprofile=coverage.out $(PACKAGES)
	go tool cover -html=coverage.out -o coverage.html

docker-test:
	docker run --rm -t --volume $$(pwd):/src:rw $(APP_ENV_IMAGE) make test
	@echo "Success, read report at file://$$(pwd)/coverage.html"

mocks: ## Generate mocks
	@rm -rf gen/mocks && mkdir -p gen/mocks
	mockery --all --case=underscore --recursive --outpkg=mocks --output=gen/mocks --dir=pkg

docker-mocks:
	docker run --rm -t --volume $$(pwd):/src:rw $(APP_ENV_IMAGE) make mocks
	@echo "Success"

vendor: ## Resolve dependencies
	go mod tidy

docker-vendor:
	docker run --rm -t --volume $$(pwd):/src:rw $(APP_ENV_IMAGE) make vendor
	@echo "Success"

owner: ## Reset folder owner
	sudo chown -R $$(id -u):$$(id -u) ./
	@echo "Success"

image-env: ## Build golang env image
	docker build --build-arg BASE_IMG=$(APP_BASE_IMAGE) --tag=$(APP_ENV_IMAGE) --target=env .
