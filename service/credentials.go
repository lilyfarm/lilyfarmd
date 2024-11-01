package service

import (
	"fmt"
	"os"
)

// This file is for loading and working with TLS certificates

const tlsCertificateEnvironmentVariable = "LILYFARM_TLS_CERTIFICATE"

const tlsKeyEnvironmentVariable = "LILYFARM_TLS_KEY"

type tlsCredentials struct {
	certificate string
	key         string
}

func (c *tlsCredentials) load() error {
	c.certificate = os.Getenv(tlsCertificateEnvironmentVariable)
	if c.certificate == "" {
		return fmt.Errorf("TLS Certificate not found in %s", tlsCertificateEnvironmentVariable)
	}

	c.key = os.Getenv(tlsKeyEnvironmentVariable)
	if c.key == "" {
		return fmt.Errorf("TLS Key not found in %s", tlsKeyEnvironmentVariable)
	}

	return nil
}

var defaultTlsCredentials tlsCredentials
