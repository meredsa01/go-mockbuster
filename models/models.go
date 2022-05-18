package models

// schema of the film table
type Film struct {
	Film_id          int64   `json:"film_id"`
	Title            string  `json:"title"`
	Description      string  `json:"description"`
	Release_year     int64   `json:"release_year"`
	Language_id      int64   `json:"language_id"`
	Rental_duration  int64   `json:"rental_duration"`
	Rental_rate      float64 `json:"rental_rate"`
	Length           int64   `json:"length"`
	Replacement_cost float64 `json:"replacement_cost"`
	Rating           string  `json:"rating"`
	Last_update      string  `json:"last_update"`
	Special_features string  `json:"speacial_features"`
	Fulltext         string  `json:"fulltext"`
}
