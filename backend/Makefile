.PHONY: run
run:
	go build main.go router.go router_gen.go
	rm -f main.pid
	./main & echo $$! > main.pid

.PHONY: stop
stop:
	kill `cat main.pid` && rm main.pid

GO_FMTTER_VERSION = 0.3.8
.PHONY: install-go-fmtter
install-go-fmtter: $(LOCALBIN)
	if ! test -x $(LOCALBIN)/gosimports || ! $(LOCALBIN)/gosimports -version | grep $(GO_FMTTER_VERSION) >/dev/null; then \
		GOBIN=$(LOCALBIN) go install github.com/rinchsan/gosimports/cmd/gosimports@v$(GO_FMTTER_VERSION); \
	fi


.PHONY: fmt-go
fmt-go: install-go-fmtter
# go mod tidy doesn't recognize the go.work file, see https://github.com/golang/go/issues/50750.
# It will report 'missing directory' error even if the 'missing directory' is already added in the other module.
# Running `go work sync` first doesn't solve the problem, if the 'missing directory' is not released.
# So we add `-e` to attempt to proceed despite errors encountered while loading packages.
	$(foreach PKG, $(GO_MODULES_EXCLUDE_SITE), \
		pushd ./${PKG} && \
			go mod tidy -e || exit 1; \
		popd; \
	)
	$(foreach PKG, $(GO_MODULES), \
		pushd ./${PKG} && \
			$(LOCALBIN)/gosimports -w -local ${PROJECT_NAME} . || exit 1; \
		popd; \
	)


