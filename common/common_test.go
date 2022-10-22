package common

import (
	"os"
	"path/filepath"
	"testing"
)

//TODO: Add test to ensure path returned exists.
//TODO: Add test to test files are valid certs and keys
func TestGetCertificateAndKeyPathForRegion(t *testing.T) {
	cPath, kPath := getCertificateAndKeyPathsForRegion("us-east-1")
	certFolder := os.Getenv("BART_CERT_FOLDER")
	wantCPath := filepath.Join(certFolder, "us-east-1.pem")
	wantKPath := filepath.Join(certFolder, "us-east-1.key")
	if cPath != wantCPath && kPath != wantKPath {
		t.Errorf("got %s and %s when expecting %s and %s", cPath, kPath, wantCPath, wantKPath)
	}

}
