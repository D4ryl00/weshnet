##
## Config
##

GO ?= go
GOPATH ?= $(HOME)/go
GO_TAGS ?= -tags "fts5 sqlite sqlite_unlock_notify"
GO_TEST_OPTS ?= -test.timeout=300s -race -cover -coverprofile=coverage.txt -covermode=atomic $(GO_TAGS)
GO_TEST_PATH ?= ./...
GO_TEST_ENV ?=
CI ?= false

VCS_REF ?= `git rev-parse --short HEAD`
VERSION ?= `go run github.com/mdomke/git-semver/v5`
LDFLAGS ?= -ldflags="-X berty.tech/weshprotocol/pkg/bertyversion.VcsRef=$(VCS_REF) -X berty.tech/weshprotocol/pkg/bertyversion.Version=$(VERSION)"

# @FIXME(gfanton): on macOS Monterey (12.0.x) we currently need to set the
# environment variable `MallocNanoZone` to 0 to avoid a SIGABRT or SIGSEGV
# see https://github.com/golang/go/issues/49138
MACOS_VERSION=$(shell defaults read /System/Library/CoreServices/SystemVersion.plist 'ProductVersion' 2>/dev/null | sed 's/\.[0-9]$$//')
ifeq ($(MACOS_VERSION),12.0)
GO_TEST_ENV := MallocNanoZone=0 $(GO_TEST_ENV)
endif

ifeq ($(MACOS_VERSION),12.1)
GO_TEST_ENV := MallocNanoZone=0 $(GO_TEST_ENV)
endif

##
## General rules
##

all: help
.PHONY: all


help:
	@echo "Available make commands:"
	@cat Makefile | grep '^[a-z]' | grep -v '=' | cut -d: -f1 | sort | sed 's/^/  /'
.PHONY: help


test: unittest lint tidy
.PHONY: test


unittest: go.unittest
.PHONY: unittest


generate: pb.generate
.PHONY: generate


regenerate: gen.clean generate
.PHONY: regenerate


clean:
	rm -rf out/
.PHONY: clean


tidy: go.tidy
.PHONY: tidy


lint: go.lint
.PHONY: lint


lint.fix: go.fmt
.PHONY: lint.fix

##
## Other rules
##


check-program = $(foreach exec,$(1),$(if $(shell PATH="$(PATH)" which $(exec)),,$(error "No $(exec) in PATH")))

go.tidy: pb.generate
	$(call check-program, $(GO))
	GO111MODULE=on $(GO) mod tidy
.PHONY: go.tidy

go.lint: pb.generate
	$(call check-program, golangci-lint)
	golangci-lint run --timeout=5m $(if $(filter $(CI), false), --verbose) ./...
.PHONY: go.lint

go.unittest: pb.generate
	$(call check-program, $(GO))
	$(GO_TEST_ENV) GO111MODULE=on $(GO) test $(GO_TEST_OPTS) $(GO_TEST_PATH)
.PHONY: go.unittest


go.flappy-tests: pb.generate
	TEST_STABILITY=flappy go run moul.io/testman test -v -test.v -timeout=600s -retry=10 -run ^TestFlappy    ./
	TEST_STABILITY=flappy go run moul.io/testman test -v -test.v -timeout=600s -retry=10 -run ^TestScenario_ ./
	TEST_STABILITY=flappy go run moul.io/testman test -v -test.v -timeout=600s -retry=10 -run ^TestFlappy    ./internal/tinder
	# FIXME: run on other packages too
.PHONY: go.flappy-tests


go.broken-tests: pb.generate
	TEST_STABILITY=broken go run moul.io/testman test -continue-on-error -timeout=1200s -test.timeout=60s -retry=5 -run ^TestScenario_ ./
.PHONY: go.broken-tests


print-%:
	@echo $* = $($*)


minimum_go_minor_version = 14
validate-go-version:
	@if [ ! "x`$(GO) version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f1`" = "x1" ]; then \
		echo "error: Golang version should be \"1.x\". Please use 1.$(minimum_go_minor_version) or more recent."; \
		exit 1; \
	fi
	@if [ `$(GO) version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f2` -lt $(minimum_go_minor_version) ]; then \
		echo "error: Golang version is not supported. Please use 1.$(minimum_go_minor_version) or more recent."; \
		exit 1; \
	fi
.PHONY: validate-go-version


##
## Code gen
##


protos_src := $(wildcard api/*.proto) $(wildcard api/go-internal/*.proto)
gen_src := $(protos_src) Makefile
gen_sum := gen.sum
protoc_opts := -I api:`go list -m -mod=mod -f {{.Dir}} github.com/grpc-ecosystem/grpc-gateway`/third_party/googleapis:`go list -m -mod=mod -f {{.Dir}} github.com/gogo/protobuf`:/protobuf
pb.generate: gen.sum validate-go-version
$(gen_sum): $(gen_src)
	$(call check-program, shasum docker $(GO))
	@shasum $(gen_src) | sort -k 2 > $(gen_sum).tmp
	@diff -q $(gen_sum).tmp $(gen_sum) || ( \
	  uid=`id -u`; \
	  set -xe; \
	  $(GO) mod download; \
	  docker run \
		--user="$$uid" \
		--volume="`go env GOPATH`/pkg/mod:/go/pkg/mod" \
		--volume="$(PWD):/go/src/berty.tech/weshprotocol" \
		--workdir="/go/src/berty.tech/weshprotocol" \
		--entrypoint="sh" \
		--rm \
		bertytech/protoc:29 \
		-xec 'make generate_local'; \
	  $(MAKE) tidy \
	)
.PHONY: pb.generate


generate_local:
	go version
	$(call check-program, shasum protoc)
	$(GO) run github.com/buicongtan1997/protoc-gen-swagger-config -i api/protocoltypes.proto -o api/protocoltypes.yaml
	$(GO) run github.com/buicongtan1997/protoc-gen-swagger-config -i api/pushtypes.proto -o api/pushtypes.yaml
	$(GO) run github.com/buicongtan1997/protoc-gen-swagger-config -i api/bertyreplication.proto -o api/bertyreplication.yaml
	protoc $(protoc_opts) --gogo_out=plugins=grpc:$(GOPATH)/src api/errcode.proto
	protoc $(protoc_opts) --gogo_out=plugins=grpc:$(GOPATH)/src api/pushtypes.proto
	protoc $(protoc_opts) --gogo_out=plugins=grpc:$(GOPATH)/src api/go-internal/records.proto
	protoc $(protoc_opts) --gogo_out=plugins=grpc:$(GOPATH)/src api/go-internal/handshake.proto
	protoc $(protoc_opts) --gogo_out=plugins=grpc:$(GOPATH)/src --grpc-gateway_out=logtostderr=true,grpc_api_configuration=api/protocoltypes.yaml:$(GOPATH)/src api/protocoltypes.proto
	protoc $(protoc_opts) --gogo_out=plugins=grpc:$(GOPATH)/src --grpc-gateway_out=logtostderr=true,grpc_api_configuration=api/bertyreplication.yaml:$(GOPATH)/src api/bertyreplication.proto
	protoc $(protoc_opts) --gogo_out=plugins=grpc:$(GOPATH)/src --grpc-gateway_out=logtostderr=true,grpc_api_configuration=api/pushtypes.yaml:$(GOPATH)/src api/pushtypes.proto
	$(MAKE) go.fmt
	shasum $(gen_src) | sort -k 2 > $(gen_sum).tmp
	mv $(gen_sum).tmp $(gen_sum)
.PHONY: generate_local


go.fmt:
	go run github.com/daixiang0/gci write . \
		--skip-generated -s 'standard,default,prefix(berty.tech)'
	go run mvdan.cc/gofumpt -w .

.PHONY: go.fmt

gen.clean:
	rm -f gen.sum $(wildcard */*/*.pb.go) $(wildcard */*/*pb_test.go) $(wildcard */*/*pb.gw.go)
.PHONY: gen.clean

asdf.install_plugins:
	$(call check-program, asdf)
	@echo "Installing asdf plugins..."
	@set -e; \
	for PLUGIN in $$(cut -d' ' -f1 .tool-versions | grep "^[^\#]"); do \
		asdf plugin add $$PLUGIN || [ $$? == 2 ] || exit 1; \
	done

asdf.install_tools: asdf.install_plugins
	$(call check-program, asdf)
	@echo "Installing asdf tools..."
	@asdf install
