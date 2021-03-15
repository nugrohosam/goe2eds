package v1

// ExmpleItems ...
type KeyItems struct {
	DownloadPrivateKeyUrl string `structs:"download_private_key_url" json:"download_private_key_url"`
	DownloadPublicKeyUrl string `structs:"download_public_key_url" json:"download_public_key_url"`
}

// ExmpleDetail ...
type KeyDetail struct {
	DownloadPrivateKeyUrl string `structs:"download_private_key_url" json:"download_private_key_url"`
	DownloadPublicKeyUrl string `structs:"download_public_key_url" json:"download_public_key_url"`
}

// KeyListItems ..
type KeyListItems []KeyItems
