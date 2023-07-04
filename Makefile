# Makefile

.PHONY: run-server run-dev build-vue build-go all clean

all: build run-server

run-server: build
	go run server/cmd/main.go

run-dev:
	cd client && npm run dev & cd server && go run cmd/main.go

build-vue:
	cd client && npm run build

build-go:
	cd server && go build -o cmd/vue-fiber.exe cmd/main.go

clean:
	rm -rf client/dist
	rm -f server/cmd/main
