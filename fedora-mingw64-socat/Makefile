SHELL=bash
.PHONY: container push all

container:
	docker build -t voltagex/fedora-mingw64-socat .

push:
	docker push voltagex/fedora-mingw64-socat

all: debug container push
