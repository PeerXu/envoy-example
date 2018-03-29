GO := $(shell which go)
GLIDE := $(shell which glide)
DOCKER_COMPOSE := $(shell which docker-compose)

pre_build:
	$(GLIDE) install

build_: pre_build
	$(MAKE) -C protos all
	script/build.sh

build_cmd: pre_build
	$(GO) build -o bin/greet cmd/main.go

up: build_
	$(DOCKER_COMPOSE) up --build -d

down:
	$(DOCKER_COMPOSE) down -v
