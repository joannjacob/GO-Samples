package entity

// Duck represents a duck entity stored in repository
type Duck struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}