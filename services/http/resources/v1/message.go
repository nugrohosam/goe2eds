package v1

// MessageItem ...
type MessageItem struct {
	IsValid 	bool 	`structs:"is_valid" json:"is_valid" copier:"field:IsValid"`
}

// MessageDetail ...
type MessageDetail struct {
	IsValid 	bool 	`structs:"is_valid" json:"is_valid" copier:"field:IsValid"`
}

// SignatureMessageItem ...
type SignatureMessageItem struct {
	Signature 	[]byte 	`structs:"signature" json:"signature" copier:"field:IsValid"`
}

// SignatureMessageDetail ...
type SignatureMessageDetail struct {
	Signature 	[]byte 	`structs:"signature" json:"signature" copier:"field:IsValid"`
}

// MessageListItems ..
type MessageListItems []MessageItem
