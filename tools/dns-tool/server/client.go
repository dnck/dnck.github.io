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

// dnsClient is the primary means by which the program communicates with a
// trusted dns resolver over tcp-tls
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

// dnsClientCerts is a container for holding keys and certificates for the
// dnsClient
type dnsClientCerts struct {
	PrivateKey             *rsa.PrivateKey
	PrivateKeyPemBytes     []byte
	SelfSignedCertPemBytes []byte
	SelfSignedX509Certs    *tls.Certificate
}

// makePrivateKey creates a new rsa public/private key pair. It is a helper
// function used when generating self-signed
// certificates for the dnsClient
func (c *dnsClientCerts) makePrivateKey() error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	c.PrivateKey = privateKey
	if err := c.getPrivateKeyPemBytes(); err != nil {
		return err
	}
	return nil
}

// getPrivateKeyPemBytes derives the pem encoding of the passed rsa private key.
// It is a helper function used when generating self-signed certificates.
func (c *dnsClientCerts) getPrivateKeyPemBytes() error {
	privateKeyPemBytes := new(bytes.Buffer)
	err := pem.Encode(privateKeyPemBytes, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(c.PrivateKey),
	})
	if err != nil {
		return err
	}
	c.PrivateKeyPemBytes = privateKeyPemBytes.Bytes()
	return nil
}

// makeSelfSignedCertificate generates the pem encoded ca cert for the dnsClient
func (c *dnsClientCerts) makeSelfSignedCertificate() error {
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
		&c.PrivateKey.PublicKey,
		c.PrivateKey,
	)
	if err != nil {
		return err
	}
	caCertBytesPem := new(bytes.Buffer)
	err = pem.Encode(caCertBytesPem, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caCertBytes,
	})
	if err != nil {
		return err
	}
	c.SelfSignedCertPemBytes = caCertBytesPem.Bytes()
	return nil
}

// chainCerts creates chains together the pem encoded self-signed ca
// certificates and private key for use in a tls configuration
func (c *dnsClientCerts) chainCerts() error {
	if err := c.makePrivateKey(); err != nil {

		return err
	}
	if err := c.makeSelfSignedCertificate(); err != nil {
		return err
	}
	selfSignedChainedTlsCerts, err := tls.X509KeyPair(c.SelfSignedCertPemBytes,
		c.PrivateKeyPemBytes)
	if err != nil {
		return err
	}
	c.SelfSignedX509Certs = &selfSignedChainedTlsCerts
	return nil
}

// makeTlsCert creates a chains the self-signed certificates which are used in
// the dnsClient to establish communication with a tls server
func (c *dnsClientCerts) makeTlsCert() error {
	if err := c.chainCerts(); err != nil {
		return err
	}
	return nil
}

// makeTlsConfig configures the tls connection for communication with a dns
// tls server
func (d *dnsClient) makeTlsConfig() error {
	selfSigner := dnsClientCerts{}
	if err := selfSigner.makeTlsCert(); err != nil {
		return err
	}
	d.tlsConfig = &tls.Config{Certificates: []tls.Certificate{*selfSigner.SelfSignedX509Certs},
		ServerName: d.commonName}
	d.tlsConfig.VerifyConnection = d.verifyConnection
	return nil
}

// establishTrust performs the first handshake with the tls server at program
// start; if a pin was not provided, it will store one for the trusted
// `-resolver-addr` `-resolver-fqdn` pair provided on the command line
func (d *dnsClient) establishTrust() error {
	if err := d.makeTlsConfig(); err != nil {
		return err
	}
	tlsConn, err := d.getConnection()
	if err != nil {
		return err
	}
	defer func() {
		if err := tlsConn.Close(); err != nil {
			log.Println(err.Error())
		}
	}()
	err = tlsConn.Handshake()
	if err != nil {
		return err
	}
	debugf("completed first handshake with tls server")
	return nil
}

// verifyConnection checks on each tls handshake that the tls server's
// certificates have not changed; it uses the sha256
// hash of the trusted tls server's certificate as a "pin". If the pin has
//  not been given at program start, then a pin
// is stored in memory for the first tls handshake (this is risky and
// should be avoided). If, on subsequent handshakes,
// the pin does not match expectations, then an error is returned.
func (d *dnsClient) verifyConnection(state tls.ConnectionState) error {
	for _, v := range state.PeerCertificates {
		pubKeyBytes, err := x509.MarshalPKIXPublicKey(v.PublicKey.(*ecdsa.PublicKey))
		if err != nil {
			return err
		}
		checkSum := sha256.Sum256(pubKeyBytes)
		pin := make([]byte, base64.StdEncoding.EncodedLen(len(checkSum)))
		base64.StdEncoding.Encode(pin, checkSum[:])
		// If the name on the certificate matches the name of the server we trust,
		// we will also check that the
		// hash of the certificate matches our expectations; if not, an error returns.
		if v.Subject.CommonName == d.commonName {
			if d.certPin == "" {
				// Unless the pin was passed at run time, then this block never gets
				// executed; however, if this is the first tls connection, and the
				// dnsClient was not provided with a pin, this block stores the pin in
				// memory for later checking.
				d.certPin = string(pin)
				debugf("stored pin of tls server's certificate")
				return nil

			} else if d.certPin == string(pin) && d.certPin != "" {
				// This block will only execute if either 1. the dnsClient was provided
				// with the pin at program start,
				// or 2. it stored the pin on its first tls connection/handshake;
				// Thereafter, we check that the pin  matches the stored pin on each
				// handshake, returning nil to indicate that the pin matches and it was
				// not the empty string.
				// log.Println("[dnsClient.verifyConnection] [DEBUG] tls server's pin
				return nil
			} else {
				// The pin has not matched, so we do not trust this connection and
				// return an error.
				return errors.New("pin mismatch")
			}
		} else {
			// We did not find the name of the trusted tls server on the certificates,
			// so we return an error.
			return errors.New("tls server's name not found in certificate")
		}
	}
	// Unlikely, but we return an error if the certificates were not in the tls
	// connection state
	return errors.New("peer certificate missing")
}

// getConnection returns the dnsClient's configured connection to the tls server.
// After acquiring the connection and using it, the callee should be careful
// to close it. If the dial fails, an error is returned
func (d *dnsClient) getConnection() (*tls.Conn, error) {
	conn, err := net.Dial("tcp", d.addressPort)
	if err != nil {
		return nil, err
	}
	tlsConn := tls.Client(conn, d.tlsConfig)
	return tlsConn, nil
}

func (d *dnsClient) setTlsConnDeadline(tlsConn *tls.Conn) error {
	err := tlsConn.SetDeadline(time.Now().Add(time.Duration(d.timeoutSeconds) * time.Second))
	if err != nil {
		return err
	}
	return nil
}

// readBytes reads tcp dns messages from the tls connection. Note that tcp dns
// messages are prefixed with two bytes indicating the length of the message
// (https://datatracker.ietf.org/doc/html/rfc1035#section-4.2.2)
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

// sendQuery dispatches the frontend server's tcp dns request over tls to the
// trusted dns server and returns
// the response to the frontend for serving the original client
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
