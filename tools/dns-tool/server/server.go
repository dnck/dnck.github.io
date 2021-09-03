package dnsserver

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type dnsProxyServer struct {
	address         string
	port            string
	timeoutSeconds  int
	dnsStubResolver *dnsClient
}

func (proxy *dnsProxyServer) listenAndServe() error {
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

func (proxy *dnsProxyServer) readBytes(conn net.Conn) ([]byte, error) {
	_ = conn.SetReadDeadline(time.Now().Add(time.Duration(proxy.timeoutSeconds) * time.Second))
	reader := bufio.NewReader(conn)
	// TCP DNS messages are prefixed with a two byte length field:
	// https://datatracker.ietf.org/doc/html/rfc1035#section-4.2.2
	firstTwoBytes, err := reader.Peek(2)
	if err != nil {
		return nil, err
	}
	lengthData := binary.BigEndian.Uint16(firstTwoBytes)
	allBytes, err := reader.Peek(2+int(lengthData))
	debugf(fmt.Sprintf("read %d bytes from frontend", len(allBytes)))
	if err != nil {
		return nil, err
	}
	return allBytes, nil
}

func (proxy *dnsProxyServer) handle(conn net.Conn, dnsClient dnsClient) {
	infof("received dns query from frontend")
	buf, err := proxy.readBytes(conn)
	if err != nil {
		debugf(err.Error())
		return
	}
	//fmt.Println(hex.Dump(buf))
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
