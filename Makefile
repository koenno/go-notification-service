# compile: ## compile
# 	mkdir -p $(BUILD_DIR)
# 	CGO_ENABLED=0 GOOS=linux go build -o $(BUILD_DIR)/${APP_NAME} cmd/main.go

help:  ## print help
	@grep -E '^[\.a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	  awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run: ## run
	CGO_ENABLED=0 GOOS=linux go run cmd/main.go

test: ## test
	CGO_ENABLED=0 GOOS=linux go test ./...

.EXPORT_ALL_VARIABLES:
BUILD_DIR             ?= build

.PHONY: compile help test
