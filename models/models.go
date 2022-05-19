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

// schema of the film details
type Film_details struct {
	FilmBase   Film
	Categories []Film_category
	Language   string `json:"language"`
	Actors     []Film_actor
}

type Film_category struct {
	Category_id int64  `json:"category_id"`
	Name        string `json:"name"`
}

type Category struct {
	Category_id int64  `json:"category_id"`
	Name        string `json:"name"`
	Last_update string `json:"last_update"`
}

type Film_actor struct {
	Actor_id   int64  `json:"actor_id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
}

type Film_comment struct {
	Comment_id  int64  `jason:"comment_id"`
	Film_id     int64  `json:"film_id"`
	Customer_id int64  `json:"customer_id"`
	Comment     string `json:"comment"`
	Last_update string `json:"last_update"`
}

