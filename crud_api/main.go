package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

import (
	"github.com/rodrigoschaer/go_projects/crud_api/dto"
)

var movies []Movie

func main() {
	router := mux.newRouter()

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", router))
}
