package dnsserver

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"time"
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
		go proxy.handle(clientConn, *proxy.dnsStubResolver)
	}
}

// readBytes reads tcp dns messages from the tcp connection. Note that tcp
// dns messages are prefixed with two bytes indicating the message length
// (https://datatracker.ietf.org/doc/html/rfc1035#section-4.2.2)
func (proxy *dnsProxyServer) readBytes(conn net.Conn) ([]byte, error) {
	_ = conn.SetReadDeadline(time.Now().Add(time.Duration(proxy.timeoutSeconds) * time.Second))
	reader := bufio.NewReader(conn)
	firstTwoBytes, err := reader.Peek(2)
	if err != nil {
		return nil, err
	}
	lengthData := binary.BigEndian.Uint16(firstTwoBytes)
	allBytes, err := reader.Peek(2 + int(lengthData))
	debugf(fmt.Sprintf("read %d bytes from frontend", len(allBytes)))
	if err != nil {
		return nil, err
	}
	return allBytes, nil
}

// handle dispatches the net.Conn data to the dnsProxyServer configured dns
// over tls client and responds with the tls client's message to the
// dnsProxyServer's frontend client connection
func (proxy *dnsProxyServer) handle(conn net.Conn, dnsClient dnsClient) {
	infof("received dns query from frontend")
	buf, err := proxy.readBytes(conn)
	if err != nil {
		debugf(err.Error())
		return
	}
	//fmt.Println(hex.Dump(buf))//print the dns query for debugging
	answer, err := dnsClient.sendQuery(buf)
	if err != nil {
		debugf(err.Error())
		return
	}
	_, err = conn.Write(answer)
	if err != nil {
		debugf(err.Error())
		return
	}
	infof("responded to frontend")
}
