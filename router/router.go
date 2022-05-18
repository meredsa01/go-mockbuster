package router

import (
	"github.com/meredsa01/go-mockbuster/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/films", middleware.GetAllFilms).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/filmsbytitle/{1}", middleware.GetFilmsByTitle).Methods("GET", "OPTIONS")

	return router
}
