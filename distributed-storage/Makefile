.DEFAULT_GOAL := help

TEST_PATH := $(word 2,$(MAKECMDGOALS))
EXTRA_GOALS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

.PHONY: help build run test test-race test-cover $(EXTRA_GOALS)

help:
	@echo "Usage:"
	@echo "  make build"
	@echo "  make run"
	@echo "  make test <path>"
	@echo "  make test-race <path>"
	@echo "  make test-cover <path>"
	@echo "Examples:"
	@echo "  make build"
	@echo "  make run"
	@echo "  make test ./practice/05-backend-use-cases"
	@echo "  make test-race ./practice/05-backend-use-cases"
	@echo "  make test-cover ./practice/05-backend-use-cases"
	@echo "  make test ./..."

build:
	@mkdir -p ./bin
	@go build -o ./bin/fs .

run: build
	@./bin/fs

test:
	@if [ -z "$(TEST_PATH)" ]; then \
		echo "missing path: use 'make test <path>'"; \
		exit 1; \
	fi
	@if ! command -v gotestsum >/dev/null 2>&1; then \
		echo "gotestsum not found. Install with:"; \
		echo "  go install gotest.tools/gotestsum@latest"; \
		exit 1; \
	fi
	@pkg="$(TEST_PATH)"; \
	case "$$pkg" in \
		./*|../*|/*) ;; \
		*) pkg="./$$pkg" ;; \
	esac; \
	gotestsum --format testname -- "$$pkg"

test-race:
	@if [ -z "$(TEST_PATH)" ]; then \
		echo "missing path: use 'make test-race <path>'"; \
		exit 1; \
	fi
	@if ! command -v gotestsum >/dev/null 2>&1; then \
		echo "gotestsum not found. Install with:"; \
		echo "  go install gotest.tools/gotestsum@latest"; \
		exit 1; \
	fi
	@pkg="$(TEST_PATH)"; \
	case "$$pkg" in \
		./*|../*|/*) ;; \
		*) pkg="./$$pkg" ;; \
	esac; \
	gotestsum --format testname -- -race "$$pkg"

test-cover:
	@if [ -z "$(TEST_PATH)" ]; then \
		echo "missing path: use 'make test-cover <path>'"; \
		exit 1; \
	fi
	@if ! command -v gotestsum >/dev/null 2>&1; then \
		echo "gotestsum not found. Install with:"; \
		echo "  go install gotest.tools/gotestsum@latest"; \
		exit 1; \
	fi
	@pkg="$(TEST_PATH)"; \
	case "$$pkg" in \
		./*|../*|/*) ;; \
		*) pkg="./$$pkg" ;; \
	esac; \
	gotestsum --format testname -- -cover "$$pkg"

# Swallow extra make goals so `make test <path>` works.
$(EXTRA_GOALS):
	@:
