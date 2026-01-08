
CURRENT_UID := $(shell id -u)
CURRENT_GID := $(shell id -g)

lint:
	docker run -t --rm -v $(PWD):/app -w /app golangci/golangci-lint:v2.8.0 golangci-lint run --fix