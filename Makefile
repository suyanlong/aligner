# golang1.14.1 or latest
# 1. make help
# 2. make build
# ...

.PHONY: build test fmt vet clean help

build: ## build binaray
	@go build -v

fmt: ## go fmt
	@go fmt ./...
	@#gofumpt -s -w ./..

test: ## Run unittests
	@go test ./...

clean: ## Remove previous build
	@rm -rf $(shell find . -name 'datadir' -not -path "./vendor/*")
	@rm -rf bin
	@rm -rf aligner
	@rm -rf tmp

help: ## Display this help screen
	@printf "Help doc:\nUsage: make [command]\n"
	@printf "[command]\n"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
