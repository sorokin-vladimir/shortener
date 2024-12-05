.PHONY: devw devt downw downt rebuildw rebuildt logs test shell deps

# Run dev environment for Web
devw:
		docker compose --env-file .env.dev -f configs/docker/docker-compose.yml -f configs/docker/docker-compose.dev.w.yml up

# Run dev environment for Telegram
devt:
		docker compose --env-file .env.dev -f configs/docker/docker-compose.yml -f configs/docker/docker-compose.dev.t.yml up

# Stop and remove all dev containers for Web
downw:
		docker compose --env-file .env.dev -f configs/docker/docker-compose.yml -f configs/docker/docker-compose.dev.w.yml down

# Stop and remove all dev containers for Telegram
downt:
		docker compose --env-file .env.dev -f configs/docker/docker-compose.yml -f configs/docker/docker-compose.dev.t.yml down

# Full rebuild with delete volumes for Web
rebuildw:
		docker compose --env-file .env.dev -f configs/docker/docker-compose.yml -f configs/docker/docker-compose.dev.w.yml down -v
		docker compose --env-file .env.dev -f configs/docker/docker-compose.yml -f configs/docker/docker-compose.dev.w.yml up --build --force-recreate

# Full rebuild with delete volumes for Telegram
rebuildt:
		docker compose --env-file .env.dev -f configs/docker/docker-compose.yml -f configs/docker/docker-compose.dev.t.yml down -v
		docker compose --env-file .env.dev -f configs/docker/docker-compose.yml -f configs/docker/docker-compose.dev.t.yml up --build --force-recreate

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
