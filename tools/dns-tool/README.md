# dns-tool

Proxies DNS requests over TCP to a TCP TLS DNS server 

### Build and run go binary

```bash
go build .
```

```bash
./dns-tool
```

### Build and run container image

```bash
make build
```

```bash
docker run --rm --net=host dnck.github.io/dns-tool
```

## Test

To check on the encryption, we can use `tcpdump`:

1. Show the cleartext TCP client/server packets between loopback addresses:

```bash
sudo tcpdump -i lo -X port 5353
```

2. Show the encrypted TCP client/server packets between remote addresses:

```bash
HOST_ADDRESS="dnck.fritz.box"
sudo tcpdump -X host "$HOST_ADDRESS" and 1.1.1.1
```

3. Send a DNS query using `dig` to our TCP DNS proxy:

```bash
dig +timeout=10 +tries=1 +tcp -p 5353 @0.0.0.0 www.danjcook.com
```

## Questions

### Imagine this proxy being deployed in an infrastructure. What would be the security concerns you would raise?

denial-of-server attacks

### How would you integrate that solution in a distributed, microservices-oriented and containerized architecture?

We could run it as a side-car container similar to istio.

### What other improvements do you think would be interesting to add to the project?

Caching, monitoring/statistics