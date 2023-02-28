.PHONY: all

-include .env

SHELL=/bin/bash -e

migrate-up:  ## DB migrate
	@docker exec -it go sh -c "cd migrations/ && goose postgres postgres://golang:golang@postgres:5432/postgres up"

migrate-down:  ## DB migrate
	@docker exec -it go sh -c "cd migrations/ && goose postgres postgres://golang:golang@postgres:5432/postgres down"

build:
	cd frontend && npm install && cd .. && docker-compose build

run:
	docker-compose up -d
