#!make

MAKEFLAGS += --always-make

GOLANG_BASE_IMAGE=golang:1.15-alpine
GOLANG_IMAGE=docker.pkg.github.com/egnd/go-tghandler/golang:1.15-alpine

.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

%:
	@:

########################################################################################################################

lint: ## Lint source code
	golangci-lint run --color=always --config=.golangci.yml ./pkg/...

docker-lint:
	docker run --rm -t --volume $$(pwd):/src:rw --entrypoint make $(GOLANG_IMAGE) lint
	@echo "All is OK"

test: mocks ## Test source code
	go test -mod=readonly -cover -covermode=count -coverprofile=coverage.out ./pkg/...
	go tool cover -html=coverage.out -o coverage.html
	go tool cover -func=coverage.out

docker-test:
	docker run --rm -t --volume $$(pwd):/src:rw --entrypoint make $(GOLANG_IMAGE) test
	@echo "Detailed report at file://$$(pwd)/coverage.html"

mocks: ## Generate mocks
	@rm -rf gen/mocks && mkdir -p gen/mocks
	mockery --all --case=underscore --recursive --outpkg=mocks --output=gen/mocks --dir=pkg

docker-mocks:
	docker run --rm -t --volume $$(pwd):/src:rw --entrypoint make $(GOLANG_IMAGE) mocks
	@echo "Success"

vendor: ## Resolve dependencies
	go mod tidy

docker-vendor:
	docker run --rm -t --volume $$(pwd):/src:rw --entrypoint make $(GOLANG_IMAGE) vendor
	@echo "Success"

owner: ## Reset folder owner
	sudo chown -R $$(id -u):$$(id -u) ./
	@echo "All is OK"

image-golang: ## Build golang env image
	docker build --build-arg BASE_IMG=$(GOLANG_BASE_IMAGE) --tag=$(GOLANG_IMAGE) --file=golang.Dockerfile build

check-conflicts: ## Find git conflicts
	@if grep -rn '^<<<\<<<< ' .; then exit 1; fi
	@if grep -rn '^===\====$$' .; then exit 1; fi
	@if grep -rn '^>>>\>>>> ' .; then exit 1; fi
	@echo "All is OK"

check-todos: ## Find TODO's
	@if grep -rn '@TO\DO:' .; then exit 1; fi
	@echo "All is OK"

check-master: ## Check for latest master in current branch
	@git remote update
	@if ! git log --pretty=format:'%H' | grep $$(git log --pretty=format:'%H' -n 1 origin/master) > /dev/null; then exit 1; fi
	@echo "All is OK"
