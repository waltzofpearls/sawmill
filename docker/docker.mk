COMPOSE_ENV = env COMPOSE_PROJECT_NAME=sawmill \
				  COMPOSE_FILE=docker/docker-compose.yml

docker: | compose-build compose-up

compose-build:
	@$(COMPOSE_ENV) docker-compose build

compose-up:
	@$(COMPOSE_ENV) docker-compose up -d --force-recreate
