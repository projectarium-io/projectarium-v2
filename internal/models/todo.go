package models

type Todo struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	Checked     bool   `json:"checked"`
	ProjectID   *int   `json:"project_id"`
}
