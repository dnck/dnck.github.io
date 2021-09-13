# If you got podman installed, we're going to build with it!
PODMAN := $(shell command -v podman 2>/dev/null)
DOCKER := $(shell command -v docker 2>/dev/null)
ifeq ($(PODMAN), /usr/bin/podman)
CONTAINER_ENGINE := $(PODMAN)
else
CONTAINER_ENGINE := $(DOCKER)
endif

.PHONY: help
help: ## The default task is help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.DEFAULT_GOAL := help

build_image: help_deploy## Build the image
	 "$(CONTAINER_ENGINE)" build -f ./Dockerfile -t aliash_tool
	 "$(CONTAINER_ENGINE)" image prune --filter label=stage=aliash_tool_base_image

prune_baseimage: ## Remove the intermediate image
	 "$(CONTAINER_ENGINE)" image prune --filter label=stage=aliash_tool_base_image

run_test: ## Run aliash_tool for a help message after image build
	 "$(CONTAINER_ENGINE)" run aliash_tool

help_deploy: ## Build image and run tests in container
	@echo "--------------------------------------------------------------------"
	@echo "Usage:"
	@echo "podman run \\"
	@echo "    aliash_tool <options> <commands> <options>"
	@echo "--------------------------------------------------------------------"
