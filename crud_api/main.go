package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/exp/slices"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rodrigoschaer/go_projects/crud_api/dto"
)

var movies []dto.Movie

func getMovies(writer http.ResponseWriter, req *http.Request) {
	if len(movies) == 0 {
		writer.WriteHeader(http.StatusNoContent)
		writer.Write([]byte("No movies available"))
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(movies)
}

func getMovieById(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	movieId := params["id"]

	idx := slices.IndexFunc(movies, func(m dto.Movie) bool { return m.Id == movieId })

	if idx == -1 {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("No movie was found with this Id"))
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(movies[idx])

}

func createMovie(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var movie dto.Movie
	_ = json.NewDecoder(req.Body).Decode(&movie)

	movie.Id = uuid.New().String()

	movies = append(movies, movie)
	json.NewEncoder(writer).Encode(movie)
}

func updateMovie(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	movieId := params["id"]

	var updatedMovie dto.Movie
	err := json.NewDecoder(req.Body).Decode(&updatedMovie)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Invalid request payload"))
	}

	idx := slices.IndexFunc(movies, func(m dto.Movie) bool { return m.Id == movieId })

	if idx == -1 {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("No movie was found with this Id"))
		return
	}

	movies[idx] = updatedMovie
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Movie updated successfully"))
}

func deleteMovie(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	movieId := params["id"]

	idx := slices.IndexFunc(movies, func(m dto.Movie) bool { return m.Id == movieId })
	if idx == -1 {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("No movie was found with this Id"))
		return
	}

	movies = slices.DeleteFunc(movies, func(m dto.Movie) bool { return m.Id == movieId })
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Movie deleted with success"))
}

func main() {
	movies = append(movies, dto.Movie{Id: "1", Isbn: "12345", Title: "Movie 1", Director: &dto.Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, dto.Movie{Id: "2", Isbn: "54321", Title: "Movie 2", Director: &dto.Director{Firstname: "Mary", Lastname: "Doe"}})

	router := mux.NewRouter()

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", router))

}
