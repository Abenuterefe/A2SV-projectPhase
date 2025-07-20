package models

type Task struct {
	ID          string `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"` // e.g., "pending", "in_progress", "done"
}
