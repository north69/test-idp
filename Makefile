rebuild:
	docker compose down --remove-orphans
	docker compose build go
	docker-compose up -d