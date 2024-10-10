PROJECT ?= "go-getting-started"
PROJECT_NAME := $(PROJECT)

.PHONY: all test build migration

all: test build

test:
	@go test -cover ./...

build:
	@env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./$(PROJECT_NAME)

migration:
	@touch "./migration/scripts/$$(date +%Y%m%d%H%M%S)_$(name).up.sql"