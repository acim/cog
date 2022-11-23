.PHONY: lint test test-bench test-cov

lint:
	@golangci-lint run \
		--enable-all \
		--disable deadcode \
		--disable exhaustivestruct \
		--disable golint \
		--disable ifshort \
		--disable interfacer \
		--disable ireturn \
		--disable maligned \
		--disable nosnakecase \
		--disable scopelint \
		--disable structcheck \
		--disable varcheck \
		--disable varnamelen

test:
	@go test -race ./...

test-bench:
	@go test -bench=. ./...

test-cov:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -func coverage.out
