SERVICE_NAME = url_minifier

# Сборка исполняемого файла сервиса
build:
	@echo "Building..."
	cd src && go build && mv src $(SERVICE_NAME)

# Запуск сервиса с локальным конфигом
run: build
	@echo "Running..."
	cd src && ./$(SERVICE_NAME)

## Сборка Docker-образа
image:
	@echo "Docker image building..."
	$Q docker build -t $(SERVICE_NAME) .

## Запуск Docker-образа
run\:image:
	@echo "Running docker image..."
	docker run -p 8080:8080 $(SERVICE_NAME)