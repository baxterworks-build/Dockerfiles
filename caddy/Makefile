.PHONY: build
build:
	mkdir -p output
	docker build -t mycaddy .
	docker run --rm mycaddy | tar --extract --verbose
