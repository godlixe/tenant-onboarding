build:
	docker build

up:
	docker compose up -d

down:
	docker compose down

compile:
	echo "Building app"
	go build -o ./cmd/server/main ./cmd/server/main.go