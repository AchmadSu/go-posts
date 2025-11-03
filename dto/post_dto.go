package dto

type PublicPost struct {
	ID       uint   `json:"id" binding:"omitempty"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Status   string `json:"status"`
}
