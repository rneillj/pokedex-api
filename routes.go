package main

import (
	"github.com/censhin/pokedex-api/moves"
	"github.com/censhin/pokedex-api/pokemon"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/pokemon", pokemon.CollectionResource).Methods("GET")
	router.HandleFunc("/pokemon/{id}", pokemon.MemberResource).Methods("GET")
	router.HandleFunc("/moves", moves.MovesResource).Methods("GET")
	router.HandleFunc("/moves/{id}", moves.MoveResource).Methods("GET")

	return router
}
