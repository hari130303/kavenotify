# Variables
IMAGE_NAME=kavenotify-image
DOCKERFILE=dockerfile


SERVICE_NAME=kavenotify-service
BINARY_NAME=kavenotify-binary
# Build Docker image
build:
	docker build -t $(IMAGE_NAME) -f $(DOCKERFILE) .

# Run Docker Compose to start the containers
up: 
	set GOOS=linux&& set CGO_ENABLED=0 && go build -o ${BINARY_NAME} .
	docker build -t $(IMAGE_NAME) -f $(DOCKERFILE) .
	docker-compose down
	docker-compose up -d

# Stop and remove containers, networks, and images created by `docker-compose up`
down:
	docker-compose down

# Clean up Docker images
clean:
	docker rmi $(IMAGE_NAME)

build_binary:
	@echo "Building ${SERVICE_NAME} binary..."
	set GOOS=linux&& set CGO_ENABLED=0 && go build -o ${BINARY_NAME} .
# chdir ..\broker-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done"