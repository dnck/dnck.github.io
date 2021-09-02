FROM docker.io/golang:alpine as build

WORKDIR /dns-tool

COPY go.mod .
COPY go.sum .
COPY dns-tool.go .
COPY server ./server
COPY certs ./certs

RUN apk add --update openssl && \
    rm -rf /var/cache/apk/* && \
    /dns-tool/certs/make-client-certs.sh && \
    go build .

FROM docker.io/golang:alpine

WORKDIR /dns-tool

COPY --from=build /dns-tool/certs ./certs
COPY --from=build /dns-tool/dns-tool .

ENTRYPOINT ["/dns-tool/dns-tool"]