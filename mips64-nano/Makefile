PHONY: build
build:
	mkdir -p output
	docker build -t nano-mips64-build .
	docker run --rm nano-mips64-build | tar --extract --verbose
