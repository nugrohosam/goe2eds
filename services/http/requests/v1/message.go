package v1

// CreateMessage ...
type CreateMessage struct {
	PrivateKey string `form:"private_key" json:"private_key" binding:"required"`
	Message    string `form:"message" json:"message" binding:"required"`
}

// VerifyMessage ...
type VerifyMessage struct {
	PublicKey string `form:"public_key" json:"public_key" binding:"required"`
	Signature []byte `form:"signature" json:"signature" binding:"required"`
	Message   string `form:"message" json:"message" binding:"required"`
}
