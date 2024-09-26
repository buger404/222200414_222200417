.PHONY: run
run:
	go build main.go router.go router_gen.go
	rm -f main.pid
	./main & echo $$! > main.pid

.PHONY: stop
stop:
	kill `cat main.pid` && rm main.pid

GO_FMTTER_VERSION = 0.3.8
SHELL = /bin/bash
OS = $(shell uname)

ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
LOCALBIN := $(ROOT_DIR)/bin
$(LOCALBIN):
	@mkdir -p $(LOCALBIN)

GO_PROD_MODULES = biz pack utils
GO_MODULES = $(GO_PROD_MODULES) e2e site tools ./examples/dev_your_plugin
GO_MODULES_EXCLUDE_SITE = $(filter-out site,$(GO_MODULES))

PROJECT_NAME = Olympic

.PHONY: install-go-fmtter
install-go-fmtter: $(LOCALBIN)
	if ! test -x $(LOCALBIN)/gosimports || ! $(LOCALBIN)/gosimports -version | grep $(GO_FMTTER_VERSION) >/dev/null; then \
		GOBIN=$(LOCALBIN) go install github.com/rinchsan/gosimports/cmd/gosimports@v$(GO_FMTTER_VERSION); \
	fi

.PHONY: fmt-go
fmt-go: install-go-fmtter
	$(foreach PKG, $(GO_PROD_MODULES), \
		pushd ./${PKG} && \
			go mod tidy -e || exit 1; \
		popd; \
	)
	$(foreach PKG, $(GO_PROD_MODULES), \
		pushd ./${PKG} && \
			$(LOCALBIN)/gosimports -w -local $(PROJECT_NAME) . || exit 1; \
		popd; \
	)
