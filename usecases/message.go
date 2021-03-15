package usecases

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func CreateMessage(privateKey *ecdsa.PrivateKey, message []byte) {
	hash := sha256.Sum256(message)
	
	sign, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		panic(err)
	}

	fmt.Println(sign)
}

func VerifyMessage(privateKey *ecdsa.PrivateKey, message []byte) {
	hash := sha256.Sum256(message)

	sig, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		panic(err)
	}

	valid := ecdsa.VerifyASN1(&privateKey.PublicKey, hash[:], sig)
	fmt.Println("signature verified:", valid)
}
