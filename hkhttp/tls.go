package hkhttp

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"time"
)

type (
	_GenTLSConfigOptions struct {
		ip       string
		hostname string
		subject  pkix.Name
	}
	GenTLSConfigOption func(o *_GenTLSConfigOptions)
)

func WithIP(ip string) GenTLSConfigOption {
	return func(o *_GenTLSConfigOptions) {
		o.ip = ip
	}
}

func WithHostname(hostname string) GenTLSConfigOption {
	return func(o *_GenTLSConfigOptions) {
		o.hostname = hostname
	}
}

func WithSubject(subject pkix.Name) GenTLSConfigOption {
	return func(o *_GenTLSConfigOptions) {
		o.subject = subject
	}
}

func GenTLSConfig(ip, serverName string, subject pkix.Name) (conf *tls.Config, err error) {

	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, max)
	if err != nil {
		return
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP(ip)},
	}

	pk, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return
	}
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
	if err != nil {
		return
	}
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"h3", "h2", "h1"},
		ServerName:   serverName,
	}, nil
}

func GenTLSConfigNoErr(ops ...GenTLSConfigOption) *tls.Config {
	op := &_GenTLSConfigOptions{
		hostname: "localhost",
		ip:       "127.0.0.1",
		subject: pkix.Name{
			Country:            []string{"CN"},
			Province:           []string{"BeiJing"},
			Organization:       []string{"Devops"},
			OrganizationalUnit: []string{"certDevops"},
			CommonName:         "127.0.0.1",
		},
	}
	for _, o := range ops {
		o(op)
	}
	c, _ := GenTLSConfig(op.ip, op.hostname, op.subject)
	return c
}
