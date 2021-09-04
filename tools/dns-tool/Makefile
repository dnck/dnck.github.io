.PHONY: build podman-build usage help
help: ## The default task is help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.DEFAULT_GOAL := help

build: ## Build and tag the image 'dnck.github.io/dns-tool' with docker
	 sudo docker build -t dnck.github.io/dns-tool .

usage: ## sudo docker run -d --dns 127.0.0.1 --name dns-tool dnck.github.io/dns-tool