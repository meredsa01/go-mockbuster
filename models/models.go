package models

// schema of the film table
type film struct {
	film_id      int64  `json:"film_id"`
	title        string `json:"title"`
	description  string `json:"description"`
	release_year int64  `json:"release_year"`
}
