.PHONY: generate
generate:
	buf generate

.PHONY: run
run-photos-server:
	go run cmd/main.go

.PHONY: format
format:
	smartimports -local "github.com/photo-pixels/user-account/"

.PHONY: lint-full
lint-full:
	goimports -w ./internal/..
	golangci-lint run --config=.golangci.yaml ./...