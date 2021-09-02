package dnsserver

import (
	"fmt"
	"net"
	"time"
)

type DnsQuestion struct {
	ID          string
	Flags       string
	QueryDomain string
	RequestType string
}

type dnsProxyServer struct {
	address             string
	port                int
	readTimeoutDuration time.Duration
	dnsStubResolver     *dnsClient
}

func (proxy *dnsProxyServer) listenAndServe() error {
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", proxy.address, proxy.port))
	if err != nil {
		return err
	}
	for {
		clientConn, err := ln.Accept()
		if err != nil {
			return err
		}
		go proxy.handle(clientConn, *proxy.dnsStubResolver)
	}
}

func (proxy *dnsProxyServer) handle(clientConn net.Conn, dnsClient dnsClient) {
	fmt.Println("[dnsProxyServer.handle] [DEBUG] dispatching dns-request over tls")
	buf := make([]byte, 0, 4096) // big buffer for perhaps large domain names in the dns query
	tmp := make([]byte, 256)
	err := clientConn.SetReadDeadline(time.Now().Add(proxy.readTimeoutDuration))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for {
		n, err := clientConn.Read(tmp)
		if err != nil {
			fmt.Println("[dnsProxyServer.handle] [ERROR]", err.Error())
			break
		}
		buf = append(buf, tmp[:n]...)
	}
	answer, err := dnsClient.sendQuery(buf)
	if err != nil {
		return
	}
	n, err := clientConn.Write(answer)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(fmt.Sprintf("[dnsProxyServer.handle] [DEBUG] wrote %d bytes to dnsClient connection", n))
}

func (settings *Settings) MakeDnsProxyServer() error {
	fmt.Println("[MakeDnsProxyServer] [DEBUG] responding to DNS requests @ tcp://localhost:5353")
	dnsClient, err := settings.makeDnsClient()
	if err != nil {
		return err
	}
	proxyServer := dnsProxyServer{settings.Address, settings.Port,
		2 * time.Second, dnsClient,
	}
	return proxyServer.listenAndServe()
}
