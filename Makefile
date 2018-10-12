SERVICE_NAME = url_minifier

build:  ## Build the executable file of service.
	@echo "Building..."
	cd src && go build && mv src $(SERVICE_NAME)

run: build  ## Run a service.
	@echo "Running..."
	cd src && ./$(SERVICE_NAME)

image: ## Build a docker image.
	@echo "Docker image building..."
	$Q docker build -t $(SERVICE_NAME) .

run\:image: ## Run a docker image.
	@echo "Running docker image..."
	docker run -p 8080:8080 $(SERVICE_NAME)

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-\:]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ": .*?## "}; {gsub(/[\\]*/,""); printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'