package v1

// Example ...
type Example struct {
	ExampleID int `form:"user_id" json:"user_id" binding:"required,numeric"`
}
