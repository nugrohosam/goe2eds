package usecases

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"

	helpers "github.com/nugrohosam/goe2eds/helpers"
)

// CreateMessage ..
func CreateMessage(privateKey string, message []byte) ([]byte, error) {
	hash := sha256.Sum256(message)
	decodedKey := helpers.DecodePrivateKey(privateKey)
	return ecdsa.SignASN1(rand.Reader, decodedKey, hash[:])
}

// VerifyMessage ..
func VerifyMessage(publicKey string, sig, message []byte) (bool, error) {
	hash := sha256.Sum256(message)
	decodedKey := helpers.DecodePublicKey(publicKey)
	return ecdsa.VerifyASN1(decodedKey, hash[:], sig), nil
}
