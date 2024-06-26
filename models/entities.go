package models

// Struct who will represent our Todo, like a DB table
type Todo struct {
	//struct tags, when encoding/decoding JSON, ID becomes id and so on
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
