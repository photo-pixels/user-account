LOCAL_DB_HOST:=localhost
LOCAL_DB_NAME:=user_account
LOCAL_DB_DSN:="host=$(LOCAL_DB_HOST) dbname=$(LOCAL_DB_NAME) sslmode=disable"

.PHONY: generate
generate:
	mkdir -p vendor.protogen
	cp -R api/user_account/ vendor.protogen/
	buf generate

.PHONY: format
format:
	smartimports -local "github.com/photo-pixels/user-account/"

.PHONY: lint-full
lint-full:
	goimports -w ./internal/..
	golangci-lint run --config=.golangci.yaml ./...

.PHONY: run
run:
	go run cmd/main.go

migrate-create\:%:
	goose -dir=./migrations create $* sql

.PHONY: migrate-up
migrate-up:
	goose -allow-missing -dir migrations postgres $(LOCAL_DB_DSN) up

.PHONY: db
db:
	psql -c "drop database if exists $(LOCAL_DB_NAME)"
	psql -c "create database $(LOCAL_DB_NAME)"
	goose -allow-missing -dir migrations postgres $(LOCAL_DB_DSN) up

.PHONY: schema
schema:
	pg_dump -d $(LOCAL_DB_NAME) --schema-only --no-owner --no-privileges --no-tablespaces --no-security-labels --no-comments |  sed -e '/^--/d' > schema.sql
	sqlc generate

clear-cache:
	buf mod clear-cache