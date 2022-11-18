package router

import (
	"mongoapi/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.Getallmovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.Createmovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.Markwatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.Deletemovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovies", controller.Deleteallmovies).Methods("DELETE")

	return router
}
