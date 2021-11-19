# path to actual config - the one that is copied to the docker container
CONFIG:=resources/config/config.yaml

# path to docker compose file
DCOMPOSE:=docker-compose.yaml

# path to external config which will copied to CONFIG
CONFIG_PATH=resources/config/config_default.yaml

# for faster building
DOCKER_COMPOSE_FEATURES:=COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1

all: down build up

debug: down build-debug up-debug

mod:
	go mod tidy && go get ./... && go mod vendor

# ---------------------------------------------------------- #

down:
	docker-compose -f ${DCOMPOSE} down

build:
	cp ${CONFIG_PATH} ${CONFIG}
	${DOCKER_COMPOSE_FEATURES} docker-compose build --no-cache --pull --build-arg CONFIG=${CONFIG}
	
build-debug:
	cp ${CONFIG_PATH} ${CONFIG}
	${DOCKER_COMPOSE_FEATURES} docker-compose build --build-arg CONFIG=${CONFIG}

up:
	docker-compose -f ${DCOMPOSE} up -d

up-debug:
	docker-compose -f ${DCOMPOSE} up
