package dnsserver

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/miekg/dns"
	"net"
)

type client struct {
	// the sha256 hash of the trusted dns resolver's tls certificate (also known as SPKI)
	trustedDnsResolverPin string
	// the trusted dns resolver's address and port and port
	trustedDnsResolverAddress string
	// trusted dns fully qualified domain name on certificate common name field
	trustedDnsResolverFqdn string
	// safe tls connection
}

func (c *client) establishTrust() error {
	cert, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client.key")
	if err != nil {
		return err
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert}, ServerName: c.trustedDnsResolverFqdn}

	conn, err := net.Dial("tcp", c.trustedDnsResolverAddress)
	if err != nil {
		return err
	}

	safeConn := tls.Client(conn, config)
	defer func() {
		if err := safeConn.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	err = safeConn.Handshake()
	if err != nil {
		return err
	}

	state := safeConn.ConnectionState()
	return c.verifyConnectionState(state)
}

func (c *client) verifyConnectionState(state tls.ConnectionState) error {
	for _, v := range state.PeerCertificates {
		pubDer, err := x509.MarshalPKIXPublicKey(v.PublicKey.(*ecdsa.PublicKey))
		if err != nil {
			return err
		}
		sum := sha256.Sum256(pubDer)
		pin := make([]byte, base64.StdEncoding.EncodedLen(len(sum)))
		base64.StdEncoding.Encode(pin, sum[:])
		if v.Subject.CommonName == c.trustedDnsResolverFqdn {
			if c.trustedDnsResolverPin == "" {
				// This is the first tls connection, so we store the pin in memory
				c.trustedDnsResolverPin = string(pin)
				fmt.Println("stored the pin!")
				return nil

			} else if c.trustedDnsResolverPin == string(pin) && c.trustedDnsResolverPin != "" {
				// we've already established the pin, so we check that they match
				fmt.Println("pin matched stored!")
				return nil

			} else {
				// the pin has not matched, so we do not trust this connection
				fmt.Println("pin matched stored!")
				return errors.New("the dns resolver pin does not match the expected pin")
			}
		} else {
			fmt.Println("common name not found")
			return errors.New("the provided trusted dns resolver fqdn was not a common name")
		}
	}
	return errors.New("no peer certs to check")
}

func (c *client) request(w dns.ResponseWriter, req *dns.Msg) {
	fmt.Println("got a dns question")
	cert, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client.key")
	if err != nil {
		fmt.Println(err.Error())
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert},
		ServerName: c.trustedDnsResolverFqdn,
		VerifyConnection: c.verifyConnectionState,
	}
	dnsClient := &dns.Client{Net: "tcp-tls", TLSConfig: config}

	resp, _, err := dnsClient.Exchange(req, c.trustedDnsResolverAddress)
	if err != nil {
		return
	}
	err = w.WriteMsg(resp)
	if err != nil {
		return
	}
}
