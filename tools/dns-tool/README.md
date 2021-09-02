# dns-tool

Proxies DNS requests over TCP to a TCP TLS DNS service

### Build and run go binary

```bash
go build .
```

```bash
mkdir certs && \
  openssl req \
      -new -nodes -x509 \
      -out certs/client.pem \
      -keyout certs/client.key \
      -days 3650 \
      -subj "/C=DE/ST=NRW/L=Earth/O=N26/OU=IT/CN=www.n26.com/emailAddress=cookdj0128@gmail.com" && \
  ./dns-tool
```

### Build and run container image

```bash
make build
```

```bash
podman run --rm --net=host dnck.github.io/dns-tool
```

## Test

To check on the encryption, we can use `tcpdump`:

1. Show the cleartext TCP client/server packets between loopback addresses:

```bash
sudo tcpdump -i lo -X port 5353
```

2. Show the encrypted TCP client/server packets between remote addresses:

```bash
sudo tcpdump -X dst port 853
```

3. Send a DNS query using `dig` to our TCP DNS proxy:

```bash
dig +timeout=5 +tries=1 +tcp -p 5353 @0.0.0.0 www.danjcook.com
```

## Questions

### Imagine this proxy being deployed in an infrastructure. What would be the security concerns you would raise?

### How would you integrate that solution in a distributed, microservices-oriented and containerized architecture?

### What other improvements do you think would be interesting to add to the project?