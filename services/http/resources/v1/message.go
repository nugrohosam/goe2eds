package v1

// MessageItem ...
type MessageItem struct {
	IsValid 	bool 	`structs:"is_valid" json:"is_valid" copier:"field:IsValid"`
}

// MessageDetail ...
type MessageDetail struct {
	IsValid 	bool 	`structs:"is_valid" json:"is_valid" copier:"field:IsValid"`
}

// MessageListItems ..
type MessageListItems []MessageItem
