package models

// schema of the film table
type Film struct {
	Film_id      int64  `json:"film_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Release_year int64  `json:"release_year"`
}
