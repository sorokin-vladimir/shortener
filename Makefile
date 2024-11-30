.PHONY: dev down rebuild logs test shell deps

# Run dev environment
dev:
		docker compose --env-file .env.dev -f configs/docker/docker-compose.yml -f configs/docker/docker-compose.dev.yml up

# Stop and remove all containers
down:
		docker compose --env-file .env.dev -f configs/docker/docker-compose.yml -f configs/docker/docker-compose.dev.yml down

# Full rebuild with delete volumes
rebuild:
		docker compose --env-file .env.dev -f configs/docker/docker-compose.yml -f configs/docker/docker-compose.dev.yml down -v
		docker compose --env-file .env.dev -f configs/docker/docker-compose.yml -f configs/docker/docker-compose.dev.yml up --build --force-recreate

# Watch logs
logs:
		docker compose logs -f

# Run tests inside container
test:
		docker compose exec shortener go test ./...

# Connect to app shell
shell:
		docker compose exec shortener sh

# Update dependencies
deps:
		docker compose exec shortener go mod download
