package router

import (
	controller "mongo-api/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	//get all movies
	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")

	//create a movie
	router.HandleFunc("/api/movie", controller.InsertMovie).Methods("POST")

	//update a movie
	router.HandleFunc("/api/movie/{id}", controller.UpdateMovie).Methods("PUT")

	//delete a movie and all movies
	router.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovies", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
