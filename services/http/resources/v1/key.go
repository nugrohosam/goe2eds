package v1

// ExmpleItems ...
type KeyItems struct {
	PrivateKey string `structs:"private_key" json:"private_key"`
	PublicKey string `structs:"public_key" json:"public_key"`
}

// KeyDetail ...
type KeyDetail struct {
	PrivateKey string `structs:"private_key" json:"private_key"`
	PublicKey string `structs:"public_key" json:"public_key"`
}

// KeyListItems ..
type KeyListItems []KeyItems
