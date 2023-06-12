PROJECT_BINARY=smoothlife
PROJECT_BINARY_OUTPUT=bin

.PHONY: all

all: help

## Build:
tidy: ## Tidy project
	@go mod tidy

clean: tidy ## Clean project
	@rm -rfv bin

build: clean ## Build project
	@GO111MODULE=on CGO_ENABLED=0 go build -ldflags="-w -s" -o ${PROJECT_BINARY_OUTPUT}/${PROJECT_BINARY} ./... 

test: build ## Tests project
	@go clean -testcache
	@go test -v ./... -race -count=1

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  make <target>'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    %-20s%s\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  %s\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)
