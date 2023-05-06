# Variables
API_NAME := ApiCookMaster
API_PORT := 8080
API_DOCKER_IMAGE := my-api-image
DB_DOCKER_IMAGE := my-db-image

.PHONY: help
help: ## Affiche ce message d'aide
	@echo "Utilisation: make [cible]"
	@echo ""
	@echo "Cibles:"
	@awk -F ': ## ' \
		'{printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' \
		$(MAKEFILE_LIST) | column -t -s $$'\t' |  sed 's/^/  /'

.PHONY: build
build:
	go build -o bin/ApiCookMaster


.PHONY: run
run: ## Démarre l'application
	go run main.go

.PHONY: docker-build
docker-build: ## Compile et construit les images Docker
	docker build -t $(API_DOCKER_IMAGE) .

.PHONY: docker-compose-up
docker-compose-up: ## Démarre les services de l'application via Docker Compose
	docker-compose up

.PHONY: docker-compose-down
docker-compose-down: ## Arrête les services de l'application via Docker Compose
	docker-compose down

.PHONY: clean
clean: ## Supprime les artefacts de compilation
	rm -rf bin/*

.PHONY: test
test: ## Lance les tests
	go test ./...

