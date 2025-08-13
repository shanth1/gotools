.PHONY: build

go:
	go run main.go

build:
	go build -o ./build/ main.go
