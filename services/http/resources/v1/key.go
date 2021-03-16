package v1

// KeyItem ...
type KeyItem struct {
	PrivateKey   string	`structs:"private_key" json:"private_key" copier:"field:PrivateKey"`
	PublicKey 	string 	`structs:"public_key" json:"public_key" copier:"field:PublicKey"`
	CertUrl 	string 	`structs:"cert_url" json:"cert_url" copier:"field:CertUrl"`
}

// KeyDetail ...
type KeyDetail struct {
	PrivateKey   string	`structs:"private_key" json:"private_key" copier:"field:PrivateKey"`
	PublicKey 	string 	`structs:"public_key" json:"public_key" copier:"field:PublicKey"`
	CertUrl 	string 	`structs:"cert_url" json:"cert_url" copier:"field:CertUrl"`
}

// KeyListItems ..
type KeyListItems []KeyItem
