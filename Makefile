.PHONY: install-devtools
install-devtools:
	go install golang.org/x/tools/cmd/goimports@v0.24.0
	go install honnef.co/go/tools/cmd/staticcheck@v0.4.7
	go install github.com/onsi/ginkgo/v2/ginkgo@v2.8.1

.PHONY: format
format:
	gofmt -s -w .

.PHONY: imports
imports:
	goimports -l -w .

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: vet
vet:
	go vet ./...

.PHONY: lint
lint: format imports tidy vet

.PHONY: build
build:
	go build ./...

.PHONY: build-examples
build-examples:
	cd examples
	go build ./...

.PHONY: precommit
precommit: lint test-momento

.PHONY: test-momento
test-momento:
	ginkgo momento-redis/ -- -UseRedis=false

.PHONY: test-redis
test-redis:
	ginkgo momento-redis/ -- -UseRedis=true

.PHONY: test
test:
	ginkgo momento-redis/ -- -UseRedis=false
	ginkgo momento-redis/ -- -UseRedis=true
