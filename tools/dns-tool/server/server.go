package dnsserver

import (
	"fmt"
	"net"
)

// dnsProxyServer is the server that the frontend clients which do not speak
// tls use to request dns records over tcp it proxies the frontend clients
// queries to a tls server through its configured dnsStubResolver
type dnsProxyServer struct {
	// address - ip the tcp protocol will be server on for frontend clients
	address string
	// port the dnsProxyServer will bind-to to serve the tcp dns requests
	port string
	// timeoutSeconds - limit to the reads/writes to/from the frontend clients
	timeoutSeconds int
	// dnsStubResolver - client that speaks to the backend tls-tcp dns resolver
	dnsStubResolver *dnsClient
}

// listenAndServe accepts incoming requests and passes them blindly to its handle
func (proxy *dnsProxyServer) listenAndServe() error {
	infof(fmt.Sprintf("dns proxy running @ tcp://%s",
		net.JoinHostPort(proxy.address, proxy.port)))
	ln, err := net.Listen("tcp", net.JoinHostPort(proxy.address, proxy.port))
	if err != nil {
		return err
	}
	for {
		clientConn, err := ln.Accept()
		if err != nil {
			return err
		}
		//go proxy.handle(clientConn, *proxy.dnsStubResolver)
		go proxy.pipeHandle(clientConn)
	}
}

func (proxy *dnsProxyServer) pipeHandle(conn net.Conn) {
	err := proxy.dnsStubResolver.Send(conn)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}