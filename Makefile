.PHONY: build run clean lint

build:
	go build -o main.exe cmd/main.go

run: build
	./main.exe

clean:
	del /f main.exe

lint:
	go vet ./...
	golangci-lint run ./...
