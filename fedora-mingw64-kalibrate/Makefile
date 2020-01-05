.PHONY: build
build:
	mkdir -p output
	docker build -t kalibrate-build .
	docker run --rm kalibrate-build | tar --extract --verbose
