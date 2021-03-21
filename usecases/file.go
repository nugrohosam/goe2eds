package usecases

import (
	"os"
	"io/ioutil"
	"time"
	"crypto/sha256"
	"crypto"
	"crypto/rsa"
	"github.com/spf13/viper"
	helpers "github.com/nugrohosam/goe2eds/helpers"
)

var now = time.Now()

const usage = "Usage: %s INPUT_PDF_PATH OUTPUT_PDF_PATH\n"

func CreateFile(privateKey string, contentPdf, cert []byte) (string, error) {
	decodedKey := helpers.DecodePrivateKey(privateKey)
	parsedCert, err := helpers.ParseCert(cert)
	if err != nil {
		return "", err
	}
	
	formatPdf := viper.GetString("pdf.format")
	randomNameFilePdf := helpers.RandomString(10) + "." + formatPdf
	
	tmpfile, err := ioutil.TempFile("", randomNameFilePdf)	
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpfile.Name())
	tmpfile.Write(contentPdf)
	defer tmpfile.Close()
	
	return "", err
}

// VerifyFile ..
func VerifyFile(publicKey string, sig, file []byte) (bool, error) {
	hash := sha256.Sum256(file)
	decodedKey := helpers.DecodePublicKey(publicKey)
	return true, rsa.VerifyPKCS1v15(decodedKey, crypto.SHA256, hash[:], sig)
}
