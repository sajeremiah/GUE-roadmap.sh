include .env

.PHONY: 

PROJECT_ROOT := $(PWD)

%:
	@:

build:
	@go build -o $(PROJECT_ROOT)/bin/github-activity $(PROJECT_ROOT)/cmd/gue/main.go

run: build
	@$(PROJECT_ROOT)/bin/github-activity $(filter-out $@,$(MAKECMDGOALS))

race:
	@go run -race $(PROJECT_ROOT)/cmd/gue/main.go