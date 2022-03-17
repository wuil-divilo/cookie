## Configure test directories. Use valid go test package format.
lambdas = $(wildcard ./lambda/*)
UNIT_TEST_PACKAGES=./internal/... $(addsuffix /internal/...,${lambdas})
COMPONENT_TEST_PACKAGES=${lambdas}

## Target definitions
OUTPUT_DIR=.make-out/
need-setup:
	@[ -d ${OUTPUT_DIR} ] || (echo "Run 'make setup' the first time" && exit 1)

help:
	@echo "This is a helper makefile for utils-go"
	@echo "Targets:"
	@echo "    build:            sam build"
	@echo "    validate:         sam validate --profile erik"
	@echo "    lint:             run code quality report"
	@echo "    test-unit:        run unit tests"
	@echo "    test-component:   run component tests"
	@echo "    generate:         regenerate all generated files"
	@echo "    tidy:             tidy go mod"
	@echo "    setup:            setup needed development tools"

build:
	@sam build

validate:
	@sam validate --profile erik

lint: need-setup
	@echo Analyzing code
	@goreportcard-cli -v | tee "${OUTPUT_DIR}report.out"

test-unit: need-setup
	@echo Unit tests
	@-go test -coverprofile "${OUTPUT_DIR}coverage.out" ${UNIT_TEST_PACKAGES} | tee "${OUTPUT_DIR}unit-run.out"
	@go tool cover -func="${OUTPUT_DIR}coverage.out" | tee "${OUTPUT_DIR}parsed_coverage.out"

test-component: need-setup
	@echo Component tests
	@-go test ${COMPONENT_TEST_PACKAGES} | tee "${OUTPUT_DIR}component-run.out"

generate:
	@echo "generate..."
	go generate ./...
	oapi-codegen -generate types -package model -o internal/model/api.go api/openapi.yaml

tidy:
	@echo "tidy..."
	go mod tidy

setup:
	@echo "setup..."
	@./cmd/setup.sh

qa-pipeline: lint test-unit test-component
	@./cmd/checks.sh ${OUTPUT_DIR} lint test-unit test-component

pre-commit: build validate qa-pipeline
