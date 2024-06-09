---
layout: post
title: dns-proxy-server
date: 2021-09-01
---

# [dns-tool](https://github.com/dnck/dnck.github.io/tree/master/tools/dns-tool)

The Go binary `dns-tool` is a proxy server that routes un-encrypted DNS quries over
TCP to a TLS DNS server. As such, it handles the encryption of a DNS quries for
frontend clients and returns them cleartext answers from a trusted DNS resolver.


### Usage

```
Usage of ./dns-tool:
  -address string
    	address the proxy server listens on
  -port string
    	port the proxy server binds to (default "53")
  -resolver-addr string
    	the trusted dns resolver's ip address and port (default "1.1.1.1:853")
  -resolver-fqdn string
    	the trusted dns resolver's common name (default "cloudflare-dns.com")
  -resolver-pin string
    	the base64 encoded sha256 hash of the trusted dns resolver's tls cert (SPKI)
  -timeout int
    	global read/write deadline for proxy server and tls-server client (default 3)
```

## Build and run Docker image and test container

```bash
make build && \
  sudo docker run -d --dns 127.0.0.1 --name dns-tool dnck.github.io/dns-tool
```

*Note that running with --dns will overwrite the container's `/etc/resolv.conf`
file*

### Test
```bash
sudo docker exec -it dns-tool dig +tcp www.danjcook.com && \
  sudo docker logs dns-tool
```

## Build and run Go binary and test

```bash
go build . && \
  ./dns-tool --port 5353
```

### Test
```bash
dig +tcp -p 5353 @127.0.0.1 www.danjcook.com
```

## Inspect Encryption

To check on encryption, you can use the command line utilities `tcpdump` and `dig`.

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

**Bootstrap Problem**
The DNS Client in the current design of the TCP Proxy Server can be provided with
the sha256 hash of the TLS-server's SPKI (its "pin") at program start, along with
the TLS-Server's IP address and its common name on the certificate. These values
are used to confirm the TLS-Server's identity on each connection. As such, they
provide an extra layer of security. However, the program itself can be provided
with an untrusted pin at start, which would clearly have implications for its
security. For this reason, it is best to follow a Principle of Least Privilege
approach on the deployment of the proxy.

**Denial-of-service**
Under the assumption that the frontend clients are trusted, the current design
dispatches all requests to a remote TLS DNS Server without validating the
requests, rate-limiting, or caching.

On the one hand,this is a nice feature since it reduces the latency of the
response. However, if the frontend clients are not trusted, or are prone to
misbehavior, then this "nice" feature can result in the remote TLS DNS Server
limiting the rate or dns queries or entirely banning the dns client address.

**Outside of control**
Clearly, there are things which are not within the control of the proxy,
but nonetheless, there are some valid concerns such as DNS Cache poisoning in
which an attacker impersonates an authoritative nameserver and taints the cache
of the DNS resolver.

### How would you integrate that solution in a distributed, microservices-oriented and containerized architecture?

**Sidecars & CoreDNS**
If the infrastructure is already using a service-mesh such as istio, then
deploying the current proxy alongside the existing containers could come at
minimal cost. The idea would be to use the caching already built into the
[istio dns proxy](https://istio.io/latest/docs/ops/configuration/traffic-management/dns-proxy/).

However, this alone would be insufficient since each of the pod/containers
would need to modify their `/etc/resolv.conf` file to point to the proxy, which
would be burdensome for developers. For this reason, we could also
modify the Kubernetes CoreDNS configuration to forward all queries for domains
outside of the cluster to a predefined nameserver (e.g. the sidecar proxy).

https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/

However, before this is done, the current design should add a feature to
serve UDP DNS requests to frontend clients, which is a current limitation.

### What other improvements do you think would be interesting to add to the project?

**Add Unit-tests and an end-to-end test**

**Pin expiration**
The current design stores the pin of the TLS Server at start. If the pin is
non-maliciously changed (due to certificate rotation), then requests will no
longer be served. It would be interesting to think about ways to solve the
aforementioned "bootstrap problem".

https://www.ietf.org/proceedings/82/slides/websec-1.pdf

**Caching**
As already mentioned, the proxy does not interpret the contents of the queries
nor store answers from the nameserver. Caching the most queried records could
reduce latency for frontend clients.

**Monitoring/statistics**
The current design is devoid of metrics collection without which it is difficult
to determine whether it is meeting any of its SLOs.
