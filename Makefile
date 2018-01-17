HOOKS_PATH=hooks
DIFF_FILES=$(shell git diff-index --cached --name-only develop | xargs printf -- '--include=%s\n')
MODIFIED_FILES=$(shell git diff-index --cached --name-only HEAD | xargs printf -- '--include=%s\n')

.PHONY: dep setup test coverage mocks

dep:
	@go get -t -v ./... \
	github.com/onsi/ginkgo/ginkgo \
	github.com/onsi/gomega/...  \
	github.com/axw/gocov/gocov \
	github.com/vektra/mockery/.../ \
	github.com/alecthomas/gometalinter

fmt:
	@go fmt ./...

setup: dep
	@cp -f $(HOOKS_PATH)/** .git/hooks

check: setup
	@gometalinter ./... --fast $(MODIFIED_FILES)

deepcheck: setup
	@gometalinter ./... $(DIFF_FILES)

fullcheck: setup
	gometalinter ./...

test: setup
	@ginkgo -gcflags=-l ./...	

coverage: setup
	@gocov test ./... | gocov report

mocks: setup
	@scripts/build_mocks
