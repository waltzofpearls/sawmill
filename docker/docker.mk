COMPOSE_ENV = env COMPOSE_PROJECT_NAME=sawmill \
				  COMPOSE_FILE=docker/docker-compose.yml

docker: dc

dc: | dc-build dc-up dc-scale
	@$(COMPOSE_ENV) docker-compose up -d --force-recreate lb
	@$(COMPOSE_ENV) docker-compose exec api sh -c 'go run fixture/main.go'

dc-down:
	@$(COMPOSE_ENV) docker-compose down

dc-ps:
	@$(COMPOSE_ENV) docker-compose ps

dc-build:
	@echo 'Building services with docker...'
	@$(COMPOSE_ENV) docker-compose build

dc-up:
	@echo 'Starting up docker services...'
	@$(COMPOSE_ENV) docker-compose up -d --force-recreate

dc-scale:
	@echo 'Scaling docker services api=2 db_mem=2...'
	@$(COMPOSE_ENV) docker-compose scale api=2 db_mem=2
