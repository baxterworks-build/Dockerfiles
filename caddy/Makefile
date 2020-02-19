.PHONY: build install clean
build:
	mkdir -p output
	docker build -t mycaddy .
	docker run --rm mycaddy | tar --extract --verbose

install:
	sudo service caddy stop
	sudo cp output/caddy /usr/local/bin/caddy
	sudo setcap cap_net_bind_service=+ep /usr/local/bin/caddy
	sudo service caddy start

clean:
	rm output/caddy
