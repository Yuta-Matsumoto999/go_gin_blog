build:
	docker compose build
up:
	docker compose up -d
init air:
	docker compose run --rm app air init
mod tidy:
	docker-compose run --rm app go mod tidy
app:
	docker compose exec app bash
db:
	docker compose exec db bash
redis:
	docker compose exec redis bash
lacalstack:
	docker compose exec localstack bash
down:
	docker compose down
destroy:
	docker compose down --rmi all --volumes --remove-orphans
ps:
	docker compose ps