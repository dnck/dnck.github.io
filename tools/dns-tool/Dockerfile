FROM docker.io/golang:alpine as build

WORKDIR /dns-tool

COPY go.mod .
COPY go.sum .
COPY dns-tool.go .
COPY run.sh .
COPY server ./server

RUN go build .

FROM docker.io/alpine:latest

# For testing in container
RUN apk update && apk add bind-tools

WORKDIR /dns-tool

COPY --from=build /dns-tool/dns-tool .
COPY --from=build /dns-tool/run.sh .

ENTRYPOINT ["/dns-tool/dns-tool"]