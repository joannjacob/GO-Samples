package models

// Post type details
type Post struct {
	// gorm.Model
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	// created_at time.Time `json"created_at"`
	// updated_at time.Time `json:"pdated_at"`
}
