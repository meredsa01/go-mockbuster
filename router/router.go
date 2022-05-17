package router

import (
	"rxb-project/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/films", middleware.GetAllFilms).Methods("GET", "OPTIONS")

	return router
}
