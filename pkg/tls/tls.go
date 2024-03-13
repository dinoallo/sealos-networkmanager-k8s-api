package tls

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
)

// snippet credit to https://github.com/islishude/grpc-mtls-example/blob/main/cmd/client/main.go#L40
func LoadTLSConfig(certFile, keyFile, caFile string) (*tls.Config, error) {
	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	ca, err := os.ReadFile(caFile)
	if err != nil {
		return nil, err
	}

	capool := x509.NewCertPool()
	if !capool.AppendCertsFromPEM(ca) {
		return nil, fmt.Errorf("faild to append the CA certificate to CA pool")
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{certificate},
		RootCAs:      capool,
	}
	return tlsConfig, nil
}
