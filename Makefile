APP?=echo
PROJECT?=github.com/nvogel/${APP}
RELEASE=$(shell git symbolic-ref -q --short HEAD 2> /dev/null || git describe --tags --exact-match 2> /dev/null || echo $(TRAVIS_BRANCH))
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

BIN_DIR := $(GOPATH)/bin
GOMETALINTER := $(BIN_DIR)/gometalinter

PLATFORMS := windows linux darwin
os = $(word 1, $@)

.PHONY: release
release: $(PLATFORMS) ### Release

.PHONY: $(PLATFORMS)
$(PLATFORMS): ### Build per platform
	mkdir -p release
	GOOS=$(os) GOARCH=amd64 go build \
		-ldflags "-s -w \
		-X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} \
		-X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o release/${APP}-$(RELEASE)-$(os)-amd64

.PHONY: clean
clean: ### Clean
	rm -f ${APP}
	rm -rf "release"

.PHONY: test
test: ### Tests
	go test -race -v ./...

.PHONY: lint
lint: $(GOMETALINTER) ### Lint
	gometalinter --vendor --disable-all --enable=errcheck --enable=vet --enable=deadcode ./...

$(GOMETALINTER): ### Install Gometalinter if needed
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

.PHONY: help
help: ## Help
	@echo '--------------------------------------------------------------------------'
	@grep -E '### .*$$' $(MAKEFILE_LIST) | grep -v '@grep' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo "---"
	@echo "Platforms are $(PLATFORMS)"
	@echo "GOMETALINTER is $(GOMETALINTER)"
	@echo '--------------------------------------------------------------------------'
