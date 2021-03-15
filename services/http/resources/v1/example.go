package v1

// ExmpleItems ...
type ExmpleItems struct {
	Name string `structs:"name" json:"name"`
}

// ExmpleDetail ...
type ExmpleDetail struct {
	Name  string `structs:"name" json:"name"`
	Email string `structs:"email" json:"email"`
}

// ExmpleListItems ..
type ExmpleListItems []ExmpleItems
