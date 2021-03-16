package usecases

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"time"
	"math/big"
	"github.com/spf13/viper"

	helpers "github.com/nugrohosam/goe2eds/helpers"
)

// CreateKey ..
func CreateKey() (string, string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", "", "", err
	}

	userData := helpers.GetAuth()

	// Initialize X509 certificate template.
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{userData.Email},
		},
		NotBefore: now.Add(-time.Hour),
		NotAfter:  now.Add(time.Hour * 24 * 365),

		KeyUsage:              x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	// Generate X509 certificate.
	certData, err := x509.CreateCertificate(rand.Reader, &template, &template, privateKey.Public(), privateKey)
	if err != nil {
		return "", "", "", err
	}

	rootPathCert := viper.GetString("cert.root-path")
	formatCert := viper.GetString("cert.format")
	nameCert := helpers.RandomString(20) + "." + formatCert
	filePath := helpers.SetPath(rootPathCert, nameCert) 

	helpers.StoreFile(certData, filePath)
	publicLink := helpers.GetPublicLink(filePath)
	
	privKey, pubKey := helpers.EncodeKey(privateKey, &privateKey.PublicKey)

	return privKey, pubKey, publicLink, err
}
