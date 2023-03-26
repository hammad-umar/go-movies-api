package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hammad-umar/go-movies-api/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterMovieRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
