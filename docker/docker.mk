COMPOSE_ENV = env COMPOSE_PROJECT_NAME=sawmill \
				  COMPOSE_FILE=docker/docker-compose.yml

docker: | compose-build compose-up compose-scale

compose-build:
	@echo 'Building services with docker...'
	@$(COMPOSE_ENV) docker-compose build

compose-up:
	@echo 'Starting up docker services...'
	@$(COMPOSE_ENV) docker-compose up -d --force-recreate

compose-scale:
	@echo 'Scaling docker services api=2...'
	@$(COMPOSE_ENV) docker-compose scale api=2
