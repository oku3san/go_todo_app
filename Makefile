.PHONE: help build build-local up down logs ps test
.DEFAULT_GOLE := help

DOCKER_TAG := latest

build:
	docker build -t oku3san/gotodo:${DOCKER_TAG} \
		--target deploy ./

build-local:
	docker compose build --no-cache

up:
	docker compose up -d

down:
	docker compose down

logs:
	docker compose logs -f

ps:
	docker compose ps

test:
	go test -race -shuffle=on ./...
