package router

import (
	"github.com/meredsa01/go-mockbuster/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/films", middleware.GetAllFilms).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/films/title/{title}", middleware.GetFilmsByTitle).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/films/rating/{rating}", middleware.GetFilmsByRating).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/films/categoryid/{id}", middleware.GetFilmsByCategoryID).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/films/category/{category}", middleware.GetFilmsByCategory).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/filmdetails/{id}", middleware.GetFilmDetails).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/films/comment/{comment}", middleware.InsertComment).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/films/comments/{film_id}", middleware.GetCommentsByFilmID).Methods("GET", "OPTIONS")
	return router
}
