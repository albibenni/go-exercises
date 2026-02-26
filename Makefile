.DEFAULT_GOAL := help

TEST_PATH := $(word 2,$(MAKECMDGOALS))
EXTRA_GOALS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

.PHONY: help test $(EXTRA_GOALS)

help:
	@echo "Usage:"
	@echo "  make test <path>"
	@echo "Examples:"
	@echo "  make test ./practice/05-backend-use-cases"
	@echo "  make test ./..."

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

# Swallow extra make goals so `make test <path>` works.
$(EXTRA_GOALS):
	@:
