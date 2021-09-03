FROM docker.io/golang:alpine as build

WORKDIR /dns-tool

COPY go.mod .
COPY go.sum .
COPY dns-tool.go .
COPY server ./server

RUN go build .

ENTRYPOINT ["/dns-tool/dns-tool"]