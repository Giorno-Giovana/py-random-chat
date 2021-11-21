# for faster building
DOCKER_FEATURES:=DOCKER_BUILDKIT=1

all:
	docker run -t recognition

build:
	${DOCKER_FEATURES} docker build -t recognition .
