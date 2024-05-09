.PHONY: all ubuntu alpine clean

.DEFAULT_GOAL := all

all: ubuntu alpine

ubuntu:
	@echo "Building for Ubuntu"
	@docker build -t go-cshared-ubuntu-example:latest -f Dockerfile.ubuntu .
	@docker run -it --rm go-cshared-ubuntu-example:latest /bin/bash -c "./entrypoint.sh"

alpine:
	@echo "Building for Alpine"
	@docker build -t go-cshared-alpine-example:latest -f Dockerfile.alpine .
	@docker run -it --rm go-cshared-alpine-example:latest /bin/sh -c "./entrypoint.sh"

clean:
	-@docker rmi go-cshared-ubuntu-example:latest
	-@docker rmi go-cshared-alpine-example:latest
