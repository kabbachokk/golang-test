default: help

.PHONY: help
help: ## помощь
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build $(cmd)
build: ## сборка
	@go build -o $(cmd) ./cmd/$(cmd)/main.go 

.PHONY: run-server
run-server: ## запуск сервера
	@go run ./cmd/apiserver/main.go 

.PHONY: run-client
run-client: ## запуск клиента
	@go run ./cmd/apiserver-client/main.go -s $(string) -u $(url)

PHONY: test
test: ## тесты
	go test ./...