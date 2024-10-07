ifneq (,$(wildcard ./.env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

CURRENT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
CMD_DIR := cmd
CMDS := $(wildcard $(CMD_DIR)/*)
BINARIES := $(patsubst $(CMD_DIR)/%,bin/%,$(CMDS))

.PHONY: all build clean dump help

# HELP - will output the help for each task in the Makefile
# In sorted order.
# The width of the first column can be determined by the `width` value passed to awk
#
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html for the initial version.
#
help: ## This help.
	@grep -E -h "^[a-zA-Z_-]+:.*?## " $(MAKEFILE_LIST) \
	  | sort \
	  | awk -v width=20 'BEGIN {FS = ":.*?## "} {printf "\033[36m%-*s\033[0m %s\n", width, $$1, $$2}'

all: @build ## Build all binaries.

$(BINARIES): bin/%: $(CMD_DIR)/%
	@$(GO) build -o $@ ./$<

build: clean ## Clean and build all binaries.
	@$(MAKE) $(BINARIES)

clean: ## Clean up built binaries.
	@$(GO) clean
	@rm -f bin/*

dump: ## Dump environment variables and current branch information.
	@echo "Environment Variables"
	@echo "----------------------"
	@echo ""
	@echo "----------------"
	@echo "Test Variables:"
	@echo "----------------"
	@echo ""
	@echo "Current Branch:"
	@echo "---------------"
	@echo "CURRENT_BRANCH: $(CURRENT_BRANCH)"
	@echo ""
	@echo "Binary Targets:"
	@echo "---------------"
	@echo "$(BINARIES)"

test: ## Run all tests.
	@$(GO) test ./...

reset: clean build ## Clean and build binaries.
