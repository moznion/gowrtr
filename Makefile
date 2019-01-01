PKGS := $(shell go list ./...)

check: test lint vet fmt-check

test:
	go test -v $(PKGS)

lint:
	golint $(PKGS)

vet:
	go vet $(PKGS)

fmt-check:
	gofmt -l -s *.go | grep [^*][.]go$$; \
	EXIT_CODE=$$?; \
	if [ $$EXIT_CODE -eq 0 ]; then exit 1; fi; \
	goimports -l *.go | grep [^*][.]go$$; \
	EXIT_CODE=$$?; \
	if [ $$EXIT_CODE -eq 0 ]; then exit 1; fi \

fmt:
	gofmt -w -s *.go
	goimports -w *.go

