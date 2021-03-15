package usecases

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	helpers "github.com/nugrohosam/goe2eds/helpers"
)

// CreateKey ..
func CreateKey() (string, string, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", "", err
	}

	privKey, pubKey := helpers.EncodeKey(privateKey, &privateKey.PublicKey)

	return privKey, pubKey, err
}
