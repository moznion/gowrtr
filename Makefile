.PHONY: errgen

PKGS := $(shell go list ./... | grep -v go-errgen)

check: test lint vet fmt-check

test: errgen
	go test -v -cover -race -coverprofile=coverage.txt -covermode=atomic $(PKGS)

test-coverage: errgen
	go test -v -cover -coverprofile cover.out $(PKGS)
	go tool cover -html=cover.out -o cover.html

lint:
	staticcheck ./...

vet:
	go vet $(PKGS)

fmt-check:
	gofmt -l -s **/*.go | grep [^*][.]go$$; \
	EXIT_CODE=$$?; \
	if [ $$EXIT_CODE -eq 0 ]; then exit 1; fi; \
	goimports -l **/*.go | grep [^*][.]go$$; \
	EXIT_CODE=$$?; \
	if [ $$EXIT_CODE -eq 0 ]; then exit 1; fi \

fmt:
	gofmt -w -s **/*.go
	goimports -w **/*.go

installdeps:
	GO111MODULE=on go mod vendor
	GO111MODULE=on go mod tidy

bootstrap: installdeps
	GO111MODULE=on go get -u golang.org/x/lint/golint \
		golang.org/x/tools/cmd/goimports \
		github.com/moznion/go-errgen/cmd/errgen

errgen:
	go generate ./...

