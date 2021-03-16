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

	"github.com/unidoc/unipdf/v3/annotator"
	"github.com/unidoc/unipdf/v3/core"
	"github.com/unidoc/unipdf/v3/model"
	"github.com/unidoc/unipdf/v3/model/sighandler"
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
	randomNameFilePdf :=  helpers.RandomString(10) + "." + formatPdf
	
	tmpfile, err := ioutil.TempFile("", randomNameFilePdf)	
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpfile.Name())
	tmpfile.Write(contentPdf)
	defer tmpfile.Close()
	
	reader, err := model.NewPdfReader(tmpfile)
	if err != nil {
		return "", err
	}
	
	// Create appender.
	appender, err := model.NewPdfAppender(reader)
	if err != nil {
		return "", err
	}
	
	// Create sig handler.
	handler, err := sighandler.NewAdobePKCS7Detached(decodedKey, parsedCert)
	if err != nil {
		return "", err
	}
	
	// Create sig.
	sig := model.NewPdfSignature(handler)
	sig.SetName("Test Self Signed PDF")
	sig.SetReason("TestSelfSignedPDF")
	sig.SetDate(now, "")
	
	if err := sig.Initialize(); err != nil {
		return "", err
	}
	
	// Create sig field and appearance.
	opts := annotator.NewSignatureFieldOpts()
	opts.FontSize = 10
	opts.Rect = []float64{10, 25, 75, 60}
	
	field, err := annotator.NewSignatureField(
		sig,
		[]*annotator.SignatureLine{
			annotator.NewSignatureLine("Name", "John Doe"),
		},
		opts,
	)

	field.T = core.MakeString("Self signed PDF")
	if err = appender.Sign(1, field); err != nil {
		return "", err
	}
	
	rootPathPdf := viper.GetString("pdf.root-path")
	outputPath := helpers.SetPath(rootPathPdf, randomNameFilePdf)
	
	err = appender.WriteToFile(outputPath)
	if err != nil {
		return "", err
	}

	publicLink := helpers.SetPublicLink(outputPath)

	return publicLink, err
}

// VerifyFile ..
func VerifyFile(publicKey string, sig, file []byte) (bool, error) {
	hash := sha256.Sum256(file)
	decodedKey := helpers.DecodePublicKey(publicKey)

	return true, rsa.VerifyPKCS1v15(decodedKey, crypto.SHA256, hash[:], sig)
}
