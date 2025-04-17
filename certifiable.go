package certifiable

import (
	_ "embed"
	"log"
	"os"
	"path/filepath"
)

//go:embed certificates.crt
var certs string

func init() {

	cwd, _ := os.Getwd()

	if err := os.MkdirAll(filepath.Join(cwd, "certs"), 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	certPath := filepath.Join(cwd, "certs", "certificates.crt")
	if err := os.WriteFile(certPath, []byte(certs), 0644); err != nil {
		log.Fatalf("Failed to write certificates to file: %v", err)
	}

	if err := os.Setenv("SSL_CERT_DIR", filepath.Join(cwd, "certs")); err != nil {
		log.Fatalf("Failed to set SSL_CERT_DIR: %v", err)
	}
}
