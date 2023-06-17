VERSION="v0.0.2"

PROJECT_NAME="helm-changelog"
BINDIR ?= $(CURDIR)/bin
TMPDIR ?= $(CURDIR)/tmp
ARCH   ?= amd64

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

.PHONY: help build docker all clean

test-unit: ## Run unit-tests
	go test ./...

test-integration: build ## Run integration-tests
	bash hack/integration-tests.sh

test: test-unit test-integration ## Run unit and integration tests

build: ## Build binary
	mkdir -p $(BINDIR)
	CGO_ENABLED=0 go build -o ./bin/${PROJECT_NAME} .

verify: test build ## tests and builds

image: ## build docker image
	docker build -t mogensen/${PROJECT_NAME}:${VERSION} .

clean: ## clean up created files
	rm -rf \
		$(BINDIR) \
		$(TMPDIR)

all: test build docker ## Runs test, build and docker

test-coverage: ## Generate test coverage report
	mkdir -p $(TMPDIR)
	go test ./... --coverprofile $(TMPDIR)/outfile
	go tool cover -html=$(TMPDIR)/outfile

lint: ## Generate static analysis report
	golint ./...

update-docs: build ## Upgrade automatic documentations
	bash hack/update-readme.sh
