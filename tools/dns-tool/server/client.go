package dnsserver

import (
	"bufio"
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/binary"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net"
	"time"
)

type dnsClient struct {
	// sha256 hash of the dns resolver's tls certificate
	certPin string
	// the trusted dns resolver's address and port and port
	addressPort string
	// trusted dns fully qualified domain name on certificate common name field
	commonName string
	// read write deadline
	timeoutSeconds int
	// tlsConfig is configured when the dnsClient establishTrust method is called
	tlsConfig *tls.Config
}

func makePrivateKey() (*rsa.PrivateKey, []byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	privateKeyPemBytes := new(bytes.Buffer)
	err = pem.Encode(privateKeyPemBytes, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	if err != nil {
		return nil, nil, err
	}
	return privateKey, privateKeyPemBytes.Bytes(), nil
}

func makeSelfSignedCertificate(privateKey *rsa.PrivateKey) ([]byte, error) {
	caCert := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization:  []string{"N26"},
			Country:       []string{"DE"},
			Province:      []string{"Brandenburg"},
			Locality:      []string{"Berlin"},
			StreetAddress: []string{"123 Cloud Street"},
			PostalCode:    []string{"12047"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(2, 6, 0),
		IsCA:      true,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageClientAuth,
			x509.ExtKeyUsageServerAuth,
		},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}
	caCertBytes, err := x509.CreateCertificate(
		rand.Reader,
		caCert,
		caCert,
		&privateKey.PublicKey,
		privateKey,
	)
	if err != nil {
		return nil, err
	}
	caCertBytesPem := new(bytes.Buffer)
	err = pem.Encode(caCertBytesPem, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caCertBytes,
	})
	if err != nil {
		return nil, err
	}
	return caCertBytesPem.Bytes(), nil
}

func makeTlsCert() (*tls.Certificate, error) {
	privateKey, privateKeyPemBytes, err := makePrivateKey()
	certPemBytes, err := makeSelfSignedCertificate(privateKey)
	if err != nil {
		return nil, err
	}
	selfSignedTlsCert, err := tls.X509KeyPair(certPemBytes, privateKeyPemBytes)
	if err != nil {
		return nil, err
	}
	return &selfSignedTlsCert, nil
}

func makeTlsConfig(dnsResolverFqdn string) (*tls.Config, error) {
	tlsCert, err := makeTlsCert()
	if err != nil {
		return nil, err
	}
	return &tls.Config{Certificates: []tls.Certificate{*tlsCert},
		ServerName: dnsResolverFqdn}, nil
}

// establishTrust dials the address for the trusted dns resolver and performs a
// tls handshake to acquire the sha256 hash of the resolver's tls cert
func (d *dnsClient) establishTrust() error {
	tlsConfig, err := makeTlsConfig(d.commonName)
	if err != nil {
		return err
	}
	tlsConfig.VerifyConnection = d.verifyConnection
	conn, err := net.Dial("tcp", d.addressPort)
	if err != nil {
		return err
	}
	safeConn := tls.Client(conn, tlsConfig)
	defer func() {
		if err := safeConn.Close(); err != nil {
			log.Println(err.Error())
		}
	}()
	err = safeConn.Handshake()
	if err != nil {
		return err
	}
	d.tlsConfig = tlsConfig
	debugf("completed first handshake with tls server")
	return nil
}

func (d *dnsClient) verifyConnection(state tls.ConnectionState) error {
	for _, v := range state.PeerCertificates {
		pubDer, err := x509.MarshalPKIXPublicKey(v.PublicKey.(*ecdsa.PublicKey))
		if err != nil {
			return err
		}
		sum := sha256.Sum256(pubDer)
		pin := make([]byte, base64.StdEncoding.EncodedLen(len(sum)))
		base64.StdEncoding.Encode(pin, sum[:])
		if v.Subject.CommonName == d.commonName {
			if d.certPin == "" {
				// This is the first tls connection, so we store the pin in memory
				// (unless passed at run time)
				d.certPin = string(pin)
				debugf("stored hash of tls server's certificate")
				return nil

			} else if d.certPin == string(pin) && d.certPin != "" {
				// we've already established the pin, so we check that they match
				// log.Println("[dnsClient.verifyConnection] [DEBUG] tls server's pin
				// matched known pin")
				return nil

			} else {
				// the pin has not matched, so we do not trust this connection
				return errors.New("pin mismatch")
			}
		} else {
			return errors.New("tls server's name not found in cert")
		}
	}
	return errors.New("peer certs missing")
}

func (d *dnsClient) getConnection() (*tls.Conn, error) {
	conn, err := net.Dial("tcp", d.addressPort)
	if err != nil {
		return nil, err
	}
	tlsConn := tls.Client(conn, d.tlsConfig)
	err = tlsConn.SetDeadline(time.Now().Add(time.Duration(d.timeoutSeconds) * time.Second))
	if err != nil {
		return nil, err
	}
	return tlsConn, nil
}

func (d *dnsClient) readBytes(conn *tls.Conn) ([]byte, error) {
	_ = conn.SetReadDeadline(time.Now().Add(time.Duration(d.timeoutSeconds) * time.Second))
	reader := bufio.NewReader(conn)
	firstTwoBytes, err := reader.Peek(2)
	if err != nil {
		return nil, err
	}
	lengthData := binary.BigEndian.Uint16(firstTwoBytes)
	allBytes, err := reader.Peek(2 + int(lengthData))
	debugf(fmt.Sprintf("read %d bytes from tls server", len(allBytes)))
	if err != nil {
		return nil, err
	}
	return allBytes, nil
}

func (d *dnsClient) sendQuery(msg []byte) ([]byte, error) {
	tlsConn, err := d.getConnection()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := tlsConn.Close(); err != nil {
			log.Println(err.Error())
		}
	}()
	_, err = tlsConn.Write(msg)
	if err != nil {
		return nil, err
	}
	debugf("dispatched query to tls server")
	buf, err := d.readBytes(tlsConn)
	if err != nil {
		return nil, err
	}
	// fmt.Println(hex.Dump(buf))
	return buf, nil
}
