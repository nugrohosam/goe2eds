package v1

// KeyItem ...
type KeyItem struct {
	PrivateKey   string	`structs:"private_key" json:"private_key" copier:"field:PrivateKey"`
	PublicKey 	string 	`structs:"public_key" json:"public_key" copier:"field:PublicKey"`
	UrlCert 	string 	`structs:"url_cert" json:"url_cert" copier:"field:UrlCert"`
}

// KeyDetail ...
type KeyDetail struct {
	PrivateKey   string	`structs:"private_key" json:"private_key" copier:"field:PrivateKey"`
	PublicKey 	string 	`structs:"public_key" json:"public_key" copier:"field:PublicKey"`
	UrlCert 	string 	`structs:"url_cert" json:"url_cert" copier:"field:UrlCert"`
}

// KeyListItems ..
type KeyListItems []KeyItem
