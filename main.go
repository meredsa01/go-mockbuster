package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/meredsa01/go-mockbuster/router"
)

func main() {
	r := router.Router()

	fmt.Println("Starting server on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", r))
}
