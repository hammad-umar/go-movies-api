package routes

import (
	"github.com/gorilla/mux"
	"github.com/hammad-umar/go-movies-api/pkg/controllers"
)

func RegisterMovieRoutes(handler *mux.Router) {
	handler.HandleFunc("/movies", controllers.CreateMovie).Methods("POST")
	handler.HandleFunc("/movies", controllers.GetAllMovies).Methods("GET")
	handler.HandleFunc("/movies/{bookId}", controllers.GetMovie).Methods("GET")
	handler.HandleFunc("/movies/{bookId}", controllers.UpdateMovie).Methods("PUT")
	handler.HandleFunc("/movies/{bookId}", controllers.DeleteMovie).Methods("DELETE")
}
