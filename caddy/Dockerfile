FROM golang:latest
ENV GO111MODULE on
RUN mkdir /usr/src/mycaddy
WORKDIR /usr/src/mycaddy
RUN go mod init caddy
RUN go get github.com/caddyserver/caddy
COPY myplugins.go .
RUN go build myplugins.go
RUN mkdir /output
RUN mv myplugins /output/caddy
RUN chmod +x /output/caddy
ENTRYPOINT tar --create /output
