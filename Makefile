.DEFAULT_TARGET: all

.PHONY: all
all: clean lint test build

.PHONY: build
build:
	go build -o gograveyard ./cmd/gograveyard

.PHONY: clean
clean:
	rm -f gograveyard coverage.out

.PHONY: help
help:
	@echo 'Available Targets:'
	@echo '  all    remove artifacts, lint and test code, then build the binary'
	@echo '  build  build the gograveyard binary'
	@echo '  clean  delete the build and test artifacts'
	@echo '  help   print this output'
	@echo '  lint   run golangci-lint'
	@echo '  test   run all unit tests'

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -cover -coverprofile=coverage.out ./...
