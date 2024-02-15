package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Keep first letter in capitaal if you want to export
type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	dir1 := Director{"John", "Smith"}
	dir2 := Director{"Christopher", "Nolan"}
	dir3 := Director{"Tim", "Burton"}

	movies = append(movies, Movie{"1", "12345", "Movie1", &dir1})
	movies = append(movies, Movie{"2", "23456", "Movie2", &dir2})
	movies = append(movies, Movie{"3", "34567", "Movie3", &dir3})

	r.HandleFunc("/movies", GetMovies).Methods("GET")     // Movies for fetching all
	r.HandleFunc("/movies/{id}", GetMovie).Methods("GET") // Movie for fetching one
	r.HandleFunc("/movies", CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port : 8000")

	log.Fatal(http.ListenAndServe(":8000", r))

}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&newMovie)

	newMovie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, newMovie)

	json.NewEncoder(w).Encode(newMovie)

}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var newMovie Movie
			_ = json.NewDecoder(r.Body).Decode(&newMovie)

			newMovie.ID = params["id"]
			movies = append(movies, newMovie)

			json.NewEncoder(w).Encode(newMovie)
			return
		}
	}
	json.NewEncoder(w).Encode(movies)

}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}
