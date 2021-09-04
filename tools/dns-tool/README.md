# dns-tool

The Go binary, `dns-tool` is a server that routes un-encrypted DNS requests over
TCP to a TLS DNS server. As such, it handles the encryption of DNS requests for
its frontend clients and returns cleartext answers from a trusted DNS resolver.

## Build and run Go binary

```bash
go build . && \
  ./dns-tool
```

### Usage
```
Usage of ./dns-tool:
  -address string
    	address the dns-server listens on 
  -port string
    	port the dns-server binds to (default "5353")
  -resolver-addr string
    	the trusted dns resolver's ip address and port (default "1.1.1.1:853")
  -resolver-fqdn string
    	the trusted dns resolver's common name (default "cloudflare-dns.com")
  -resolver-pin string
    	the base64 encoded sha256 hash of the trusted dns resolver's tls cert (SPKI)
  -timeout int
    	global read/write deadline for proxy server and dns-client (default 3)
```

## Build image and run as container

```bash
make build && \
  docker run --rm --net=host dnck.github.io/dns-tool
```

## Test

To check on encryption, use the command line utilities `tcpdump` and `dig`.

First, show the encrypted TCP client/tls-server packets:

```bash
CLIENT_ADDRESS="dnck.fritz.box"
TLS_SERVER_ADDRESS="1.1.1.1"
sudo tcpdump -X host "$HOST_ADDRESS" and "$TLS_SERVER_ADDRESS"
```

Next, open a new terminal and send a DNS query to the our `dns-tool` frontend
server:

```bash
DNS_TOOL_ADDRESS="127.0.0.1"
DNS_TOOL_PORT="5353"
dig +timeout=10 +tcp -p "$DNS_TOOL_PORT" @"$DNS_TOOL_ADDRESS" www.danjcook.com
```

## Questions

### Imagine this proxy being deployed in an infrastructure. What would be the security concerns you would raise?

denial-of-server attacks

### How would you integrate that solution in a distributed, microservices-oriented and containerized architecture?

We could run it as a side-car container similar to istio.

### What other improvements do you think would be interesting to add to the project?

Caching, monitoring/statistics