package usecases

import (
	"log"
	"os"
	"io/ioutil"
	"time"
	"crypto/sha256"
	"crypto"
	"crypto/rsa"

	helpers "github.com/nugrohosam/goe2eds/helpers"

	"github.com/unidoc/unipdf/v3/annotator"
	"github.com/unidoc/unipdf/v3/core"
	"github.com/unidoc/unipdf/v3/model"
	"github.com/unidoc/unipdf/v3/model/sighandler"
)

var now = time.Now()

const usage = "Usage: %s INPUT_PDF_PATH OUTPUT_PDF_PATH\n"

func CreateFile(privateKey string, contentPdf, cert []byte) {

	decodedKey := helpers.DecodePrivateKey(privateKey)
	parsedCert, err := helpers.ParseCert(cert)
	if err != nil {
		log.Fatal("Fail: %v\n", err)
	}
	
	randomNameFilePdf :=  helpers.RandomString(10) + ".pdf"
	
	tmpfile, err := ioutil.TempFile("", randomNameFilePdf)	
	if err != nil {
		log.Fatal("Fail: %v\n", err)
	}
	defer os.Remove(tmpfile.Name())
	tmpfile.Write(contentPdf)
	defer tmpfile.Close()

	reader, err := model.NewPdfReader(tmpfile)
	if err != nil {
		log.Fatal("Fail: %v\n", err)
	}

	// Create appender.
	appender, err := model.NewPdfAppender(reader)
	if err != nil {
		log.Fatal("Fail: %v\n", err)
	}

	// Create signature handler.
	handler, err := sighandler.NewAdobePKCS7Detached(decodedKey, parsedCert)
	if err != nil {
		log.Fatal("Fail: %v\n", err)
	}

	// Create signature.
	signature := model.NewPdfSignature(handler)
	signature.SetName("Test Self Signed PDF")
	signature.SetReason("TestSelfSignedPDF")
	signature.SetDate(now, "")

	if err := signature.Initialize(); err != nil {
		log.Fatal("Fail: %v\n", err)
	}

	// Create signature field and appearance.
	opts := annotator.NewSignatureFieldOpts()
	opts.FontSize = 10
	opts.Rect = []float64{10, 25, 75, 60}

	field, err := annotator.NewSignatureField(
		signature,
		[]*annotator.SignatureLine{
			annotator.NewSignatureLine("Name", "John Doe"),
			annotator.NewSignatureLine("Date", "2019.16.04"),
			annotator.NewSignatureLine("Reason", "External signature test"),
		},
		opts,
	)
	field.T = core.MakeString("Self signed PDF")

	if err = appender.Sign(1, field); err != nil {
		log.Fatal("Fail: %v\n", err)
	}

	outputPath := helpers.SetPath("assets/pdf", randomNameFilePdf)

	// Write output PDF file.
	err = appender.WriteToFile(outputPath)
	if err != nil {
		log.Fatal("Fail: %v\n", err)
	}

	log.Printf("PDF file successfully signed. Output path: %s\n", outputPath)
}

// VerifyFile ..
func VerifyFile(publicKey string, sig, file []byte) (bool, error) {
	hash := sha256.Sum256(file)
	decodedKey := helpers.DecodePublicKey(publicKey)

	return true, rsa.VerifyPKCS1v15(decodedKey, crypto.SHA256, hash[:], sig)
}
