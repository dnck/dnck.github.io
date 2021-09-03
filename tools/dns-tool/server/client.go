package dnsserver

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
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

func makeTlsCert() (*tls.Certificate, error) {
	caCert := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization:  []string{"N26"},
			Country:       []string{"DE"},
			Province:      []string{"Brandenburg"},
			Locality:      []string{"Berlin"},
			StreetAddress: []string{"123 Cloud Strasse"},
			PostalCode:    []string{"12047"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(2, 6, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	privateKeyPem := new(bytes.Buffer)
	err = pem.Encode(privateKeyPem, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	if err != nil {
		return nil, err
	}
	caCertBytes, err := x509.CreateCertificate(rand.Reader, caCert, caCert, &privateKey.PublicKey, privateKey)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	caCertPem := new(bytes.Buffer)
	err = pem.Encode(caCertPem, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caCertBytes,
	})
	if err != nil {
		return nil, err
	}
	dnsClientCerts, err := tls.X509KeyPair(caCertPem.Bytes(), privateKeyPem.Bytes())
	if err != nil {
		return nil, err
	}
	return &dnsClientCerts, nil
}

func (settings *Settings) makeDnsClient() (*dnsClient, error) {
	tlsCert, err := makeTlsCert()
	//tlsCert, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client.key")
	if err != nil {
		return nil, err
	}
	c := dnsClient{settings.DnsResolverPin,
		settings.DnsResolverAddress,
		settings.DnsResolverFqdn,
		&tls.Config{Certificates: []tls.Certificate{*tlsCert}, ServerName: settings.DnsResolverFqdn},
	}
	c.tlsConfig.VerifyConnection = c.verifyConnectionState
	err = c.establishTrust()
	if err != nil {
		return nil, err
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
	fmt.Println(fmt.Sprintf("[dnsClient.sendQuery] [DEBUG] wrote %d bytes to safeConn", n))
	buf := make([]byte, 4096)
	tmp := make([]byte, 256)
	for {
		n, err := safeConn.Read(tmp)
		if err != nil {
			fmt.Println("[dnsClient.sendQuery] [ERROR]", err.Error())
			break
		}
		buf = append(buf, tmp[:n]...)
	}
	// fmt.Println(fmt.Printf("[dnsClient.sendQuery] [DEBUG] DNS ANSWER\n%s", hex.Dump(r.Bytes())))
	return buf, nil
}
