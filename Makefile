# RFC3339 (to match GoReleaser)
DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
GIT_COMMIT := $(shell git rev-parse HEAD)
GO_LDFLAGS := -ldflags="-s -w -X 'main.version=$(GIT_COMMIT)' -X 'main.commit=$(GIT_COMMIT)' -X 'main.date=$(DATE)'"

# overridable
LINTARGS := "-set_exit_status"
COVERPKG ?= -coverpkg ./internal/... -coverpkg ./lambdas/... -coverpkg ./pkg/...


# things to build
rwildcard=$(foreach d,$(wildcard $(1:=/*)),$(call rwildcard,$d,$2) $(filter $(subst *,%,$2),$d))
LAMBDAS := $(patsubst lambdas/%/main.go,%,$(call rwildcard,lambdas,*main.go))
CMDS := $(patsubst cmd/%/main.go,%,$(call rwildcard,cmd,*main.go))

.PHONY: all
all: $(LAMBDAS) $(CMDS)

.PHONY: $(LAMBDAS)
$(LAMBDAS): clean lint
	@echo "building $@..."
	@GOOS=linux GOARCH=amd64 go build -trimpath $(GO_LDFLAGS) $(BUILDARGS) -o build/$@ ./lambdas/$@/.
	@touch -mt 202001010000 build/$@
	@echo "ziping..."
	@zip -j -X build/$@.zip build/$@

.PHONY: $(CMDS)
$(CMDS): clean lint
	@echo "building $@..."
	@GOOS=linux GOARCH=amd64 go build -trimpath $(GO_LDFLAGS) $(BUILDARGS) -o build/$@-linux-amd64 ./cmd/$@/.
	@GOOS=darwin GOARCH=amd64 go build -trimpath $(GO_LDFLAGS) $(BUILDARGS) -o build/$@-darwin-amd64 ./cmd/$@/.
.PHONY: clean

clean:
	@rm -rf ./build/*

.PHONY: lint
lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run ./cmd/... ./internal/... ./lambdas/... ./pkg/...

.PHONY: gen
gen:
	go generate ./lambdas/... ./internal/... ./pkg/...

.PHONY: test
test:
	@go clean -testcache
	go test -v -timeout=180s -tags 'test' ./cmd/... ./internal/... ./lambdas/... ./pkg/...

cover:
	@go clean -testcache
	go test -v -timeout=180s $(COVERPKG) -coverprofile=coverage.out -tags 'test' ./cmd/... ./internal/... ./lambdas/... ./pkg/...
	@go tool cover -html=coverage.out

