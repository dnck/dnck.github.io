package dnsserver

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"time"
)

type dnsClient struct {
	// the sha256 hash of the trusted dns resolver's tls certificate (also known as SPKI)
	certPin string
	// the trusted dns resolver's address and port and port
	addressPort string
	// trusted dns fully qualified domain name on certificate common name field
	commonName string
	// tlsConfig
	tlsConfig *tls.Config
}

func (settings *Settings) makeDnsClient() (*dnsClient, error) {
	cert, err := tls.LoadX509KeyPair(settings.ClientCertDir+"/client.pem", settings.ClientCertDir+"/client.key")
	if err != nil {
		return nil, errors.New("[settings.makeDnsClient]: cert error")
	}
	c := dnsClient{settings.DnsResolverPin,
		settings.DnsResolverAddress,
		settings.DnsResolverFqdn,
		&tls.Config{Certificates: []tls.Certificate{cert}, ServerName: settings.DnsResolverFqdn},
	}
	c.tlsConfig.VerifyConnection = c.verifyConnectionState
	err = c.establishTrust()
	if err != nil {
		fmt.Println(err.Error())
	}
	return &c, nil
}

func (c *dnsClient) establishTrust() error {
	conn, err := net.Dial("tcp", c.addressPort)
	if err != nil {
		return err
	}
	safeConn := tls.Client(conn, c.tlsConfig)
	defer func() {
		if err := safeConn.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()
	err = safeConn.Handshake()
	if err != nil {
		return err
	}
	return nil
}

func (c *dnsClient) verifyConnectionState(state tls.ConnectionState) error {
	for _, v := range state.PeerCertificates {
		pubDer, err := x509.MarshalPKIXPublicKey(v.PublicKey.(*ecdsa.PublicKey))
		if err != nil {
			return err
		}
		sum := sha256.Sum256(pubDer)
		pin := make([]byte, base64.StdEncoding.EncodedLen(len(sum)))
		base64.StdEncoding.Encode(pin, sum[:])
		if v.Subject.CommonName == c.commonName {
			if c.certPin == "" {
				// This is the first tls connection, so we store the pin in memory
				c.certPin = string(pin)
				fmt.Println("[dnsClient.verifyConnectionState] [DEBUG] stored pin")
				return nil

			} else if c.certPin == string(pin) && c.certPin != "" {
				// we've already established the pin, so we check that they match
				fmt.Println("[dnsClient.verifyConnectionState] [DEBUG] pin matched")
				return nil

			} else {
				// the pin has not matched, so we do not trust this connection
				return errors.New("[dnsClient.verifyConnectionState] pin mismatch")
			}
		} else {
			return errors.New("[dnsClient.verifyConnectionState] common name not found")
		}
	}
	return errors.New("[dnsClient.verifyConnectionState] missing peer certs")
}

func (c *dnsClient) sendQuery(msg []byte) ([]byte, error) {
	conn, err := net.Dial("tcp", c.addressPort)
	if err != nil {
		return nil, err
	}
	timeoutDuration := 3 * time.Second
	safeConn := tls.Client(conn, c.tlsConfig)
	defer func() {
		if err := safeConn.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()
	err = safeConn.SetDeadline(time.Now().Add(timeoutDuration))
	if err != nil {
		return nil, err
	}
	n, err := safeConn.Write(msg)
	if err != nil {
		return nil, err
	}
	fmt.Println(fmt.Sprintf("[sendQuery] [DEBUG] wrote %d bytes to safeConn", n))
	buf := make([]byte, 0, 4096) // big buffer for possibly very large domain names
	tmp := make([]byte, 256)
	// TODO (dnck): what is the end of a dns message? How can we know it terminated?
	for {
		n, err := safeConn.Read(tmp)
		if err != nil {
			fmt.Println("[sendQuery] [ERROR]", err.Error())
			break
		}
		buf = append(buf, tmp[:n]...)
	}
	return buf, nil
}
