package router

import (
	"gonetflixapimongo/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("Get")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.GetAMovie).Methods("GET")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PATCH")
	router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controller.DeleteAllMovie).Methods("DELETE")

	return router
}
