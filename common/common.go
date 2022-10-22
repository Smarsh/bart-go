package common

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GetCertificateForRegion(region string) tls.Certificate {
	certificate_path, key_path := getCertificateAndKeyPathsForRegion(region)
	cert, err := loadCertificateFromPath(certificate_path, key_path)
	if err != nil {
		log.Fatal("unexpected error when trying to get certificate for region=", region, err)

	}
	return cert
}

func getCertificateAndKeyPathsForRegion(region string) (string, string) {
	baseCertPath := os.Getenv("BART_CERT_FOLDER")
	var certificate_path string
	var key_path string
	switch region := region; region {
	case "us-east-1":
		certificate_path = filepath.Join(baseCertPath, "us-east-1.pem")
		key_path = filepath.Join(baseCertPath, "us-east-1.key")
	case "eu-west-2":
		certificate_path = filepath.Join(baseCertPath, "eu-west-2.pem")
		key_path = filepath.Join(baseCertPath, "eu-west-2.pem")
	default:
		log.Fatal("unexpected error when trying to get certificate for region=", region)
	}

	return certificate_path, key_path
}

func loadCertificateFromPath(certificate_path string, key_path string) (tls.Certificate, error) {
	cert, err := tls.LoadX509KeyPair(certificate_path, key_path)
	return cert, err
}

func GetApiUrlForApp(app string, region string, tier string) string {
	urlMap := map[string]string{
		"shda__us-east-1__production":     "https://shda.cc.smarsh.cloud",
		"shda-spo__us-east-1__production": "https://shda-spo.cc.smarsh.cloud",
		"shda__eu-west-2__production":     "https://shda.cc.mt.eu-west-2.aws.smarsh.cloud",
		"shda-spo__eu-west-2__production": "https://shda-spo.cc.mt.eu-west-2.aws.smarsh.cloud",

		"prva__us-east-1__production": "https://prva.cc.smarsh.cloud",
		"prva__eu-west-2__production": "https://prva.cc.mt.eu-west-2.aws.smarsh.cloud",
	}

	key := fmt.Sprintf("%s__%s__%s", app, region, tier)
	url := urlMap[key]
	if url == "" {
		log.Fatal("Could not get api url for", app, region, tier)
	}
	return url
}
