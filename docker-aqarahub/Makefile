.PHONY: build
build:
	mkdir -p output
	docker build -t aqarahub-armhf-build .
	docker run --rm aqarahub-armhf-build | tar --extract --verbose
