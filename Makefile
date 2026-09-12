SHELL := /bin/bash

IMGPROXY_IMAGE ?= imgproxy-local:dev
IMGPROXY_DOCKERFILE ?= docker/Dockerfile
COMPOSE_DIR ?= /root/docker/sites/imgproxy
COMPOSE := docker compose -f $(COMPOSE_DIR)/docker-compose.yml --env-file $(COMPOSE_DIR)/.env

.PHONY: help build redeploy up down logs ps health

help:
	@echo "Targets (run from code/imgproxy):"
	@echo "  build     Build local imgproxy image from current source"
	@echo "  redeploy  Build + recreate imgproxy container"
	@echo "  up        Start imgproxy container"
	@echo "  down      Stop imgproxy container"
	@echo "  logs      Tail imgproxy logs"
	@echo "  ps        Show imgproxy container state"
	@echo "  health    Check imgproxy health"

build:
	docker build -f $(IMGPROXY_DOCKERFILE) -t $(IMGPROXY_IMAGE) .

up:
	IMGPROXY_IMAGE=$(IMGPROXY_IMAGE) $(COMPOSE) up -d --no-deps imgproxy

redeploy: build
	IMGPROXY_IMAGE=$(IMGPROXY_IMAGE) $(COMPOSE) up -d --force-recreate --no-deps imgproxy

down:
	$(COMPOSE) stop imgproxy

logs:
	$(COMPOSE) logs -f --tail=120 imgproxy

ps:
	$(COMPOSE) ps imgproxy

health:
	curl -fsS http://127.0.0.1:8080/health || curl -fsS https://imgproxy.l.l0l.in/health
