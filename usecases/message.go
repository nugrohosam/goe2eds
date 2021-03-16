package usecases

import (
	"crypto/rsa"
	"crypto"
	"crypto/rand"
	"crypto/sha256"
	helpers "github.com/nugrohosam/goe2eds/helpers"
)

// CreateMessage ..
func CreateMessage(privateKey string, message []byte) ([]byte, error) {
	hash := sha256.Sum256(message)
	decodedKey := helpers.DecodePrivateKey(privateKey)
	return rsa.SignPKCS1v15(rand.Reader, decodedKey, crypto.SHA256, hash[:])
}

// VerifyMessage ..
func VerifyMessage(publicKey string, sig, message []byte) (bool, error) {
	hash := sha256.Sum256(message)
	decodedKey := helpers.DecodePublicKey(publicKey)
	if decodedKey == nil {
		return false, nil
	}

	return true, rsa.VerifyPKCS1v15(decodedKey, crypto.SHA256, hash[:], sig)
}
