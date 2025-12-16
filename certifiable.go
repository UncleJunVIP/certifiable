package certifiable

import (
	"crypto/tls"
	"crypto/x509"
	_ "embed"
	"log"
	"net/http"
)

//go:embed certificates.crt
var certs []byte

func init() {
	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(certs) {
		log.Fatal("Failed to append certificates")
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
		RootCAs: pool,
	}
}
