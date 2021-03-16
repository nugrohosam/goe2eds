package v1

import "mime/multipart"

// CreateFile ...
type CreateFile struct {
	PrivateKey string `form:"private_key" json:"private_key" binding:"required"`
	File    *multipart.FileHeader `form:"file" binding:"required"`
	Cert   *multipart.FileHeader `form:"cert" binding:"required"`
}

// VerifyFile ...
type VerifyFile struct {
	PublicKey string `form:"public_key" json:"public_key" binding:"required"`
	SignatureFile *multipart.FileHeader `form:"signature_file" json:"signature_file" binding:"required"`
	File   *multipart.FileHeader `form:"file" binding:"required"`
}
