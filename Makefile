# Variables
DOCKER_COMPOSE = docker-compose

# Cibles
.PHONY: all build up down logs restart clean

all: build up

build:
	$(DOCKER_COMPOSE) build

up:
	$(DOCKER_COMPOSE) up -d

down:
	$(DOCKER_COMPOSE) down

logs:
	$(DOCKER_COMPOSE) logs api
	$(DOCKER_COMPOSE) logs db

status:
	$(DOCKER_COMPOSE) ps

db:
	docker exec -it apicookmaster_db_1 psql -U postgres

restart:
	$(DOCKER_COMPOSE) restart

clean:
	$(DOCKER_COMPOSE) down -v --remove-orphans
