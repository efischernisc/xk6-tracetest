EXTENSION_NAME=github.com/efischernisc/xk6-tracetest
.PHONY: help
help: Makefile ## show list of commands
	@echo "Choose a command run:"
	@echo ""
	@awk 'BEGIN {FS = ":.*?## "} /[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-40s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort
.PHONY: build
build: ## build this extension with K6
	xk6 build --with $(EXTENSION_NAME)=.