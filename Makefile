.PHONY: build run clean lint deploy

IMAGE_NAME=telegram-bot-video
CONTAINER_NAME=telegram-bot-video-installer

build:
	docker build -t $(IMAGE_NAME) .

run: build
	docker run --rm --name $(CONTAINER_NAME) --env-file .env $(IMAGE_NAME)

clean:
	docker rmi -f $(IMAGE_NAME) || true

lint:
	go vet ./...
	golangci-lint run ./...

# deploy:

