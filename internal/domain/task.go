package domain

type TaskItem struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Difficulty  int    `json:"difficulty"` // 1-5
	IsClosed    bool   `json:"isClosed"`
}
