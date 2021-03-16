package v1

// FileItem ...
type FileItem struct {
	IsValid 	bool 	`structs:"is_valid" json:"is_valid" copier:"field:IsValid"`
}

// FileDetail ...
type FileDetail struct {
	IsValid 	bool 	`structs:"is_valid" json:"is_valid" copier:"field:IsValid"`
}

// SignatureFileItem ...
type SignatureFileItem struct {
	FileUrl 	string 	`structs:"file_url" json:"file_url" copier:"field:IsValid"`
}

// SignatureFileDetail ...
type SignatureFileDetail struct {
	FileUrl 	string 	`structs:"file_url" json:"file_url" copier:"field:IsValid"`
}

// FileListItems ..
type FileListItems []FileItem
