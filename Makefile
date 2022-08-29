test: lint test-ci
test-ci: unit-test

lint: lint-bin ## Run go golangci-lint against code.
	$(GOLANGCI_LINT) run --timeout 5m0s

unit-test: mocks
	go test -v ./... --vet=off -coverprofile cover.out

# Generate mocks
mocks: mockgen
	$(MOCKGEN) -source utils/http.go -destination pkg/mocks/utils/mock.go HttpClient

GOLANGCI_LINT = ./bin/golangci-lint
lint-bin: ## Download golangci-lint locally if necessary.
	$(call go-get-tool,$(GOLANGCI_LINT),github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2)

MOCKGEN = ./bin/mockgen
mockgen: ## Download mockgen locally if necessary.
	$(call go-get-tool,$(MOCKGEN),github.com/golang/mock/mockgen@v1.6.0)

# go-get-tool will 'go get' any package $2 and install it to $1.
PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
define go-get-tool
@[ -f $(1) ] || { \
set -e ;\
echo "Downloading $(2)" ;\
GOBIN=$(PROJECT_DIR)/bin go install $(2) ;\
}
endef

