up:
	docker compose up -d

down:
	docker compose down -v

prod:
	go run cmd/producer/main.go

cons:
	go run cmd/consumer/main.go

PHONY: up down