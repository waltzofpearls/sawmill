COMPOSE_ENV = env COMPOSE_PROJECT_NAME=sawmill \
				  COMPOSE_FILE=docker/docker-compose.yml

setup: dc
teardown: dc-down

NUM = 100
SCALE = api=2 db_mem=2

.PHONY: fixture
fixture:
	@$(COMPOSE_ENV) docker-compose exec api sh -c 'go run fixture/gen.go $(NUM)'

dc: | dc-build dc-up dc-scale
	@$(COMPOSE_ENV) docker-compose up -d --force-recreate lb

dc-down:
	@$(COMPOSE_ENV) docker-compose down -v

dc-ps:
	@$(COMPOSE_ENV) docker-compose ps

dc-build:
	@echo 'Building services with docker...'
	@$(COMPOSE_ENV) docker-compose build

dc-up:
	@echo 'Starting up docker services...'
	@$(COMPOSE_ENV) docker-compose up -d --force-recreate

dc-scale:
	@echo 'Scaling docker services $(SCALE)...'
	@$(COMPOSE_ENV) docker-compose scale $(SCALE)
