package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Movie struct to model the data of a movie
type Movie struct {
	ID       string    `json:"id"`       // Unique identifier for each movie
	Isbn     string    `json:"isbn"`     // ISBN identifier for the movie
	Title    string    `json:"title"`    // Title of the movie
	Director *Director `json:"director"` // Director information (struct)
}

// Director struct to hold the name of the director
type Director struct {
	Firstname string `json:"firstname"` // Director's first name
	Lastname  string `json:"lastname"`  // Director's last name
}

var movies []Movie // Global slice to store all movies

// getMovies returns all movies in JSON format
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set the content type to JSON
	json.NewEncoder(w).Encode(movies)                  // Encode movies slice to JSON and write to response
}

// deleteMovie removes a movie by ID
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set content type to JSON
	params := mux.Vars(r)                              // Get URL params
	for index, item := range movies {                  // Iterate over movies
		if item.ID == params["id"] { // If ID matches, delete movie
			movies = append(movies[:index], movies[index+1:]...)
			json.NewEncoder(w).Encode(movies) // Return remaining movies
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound) // If movie is not found, return 404
}

// getMovie returns a specific movie by ID
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set content type to JSON
	params := mux.Vars(r)                              // Get URL params
	for _, item := range movies {                      // Iterate over movies
		if item.ID == params["id"] { // If ID matches, return the movie
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound) // If movie is not found, return 404
}

// createMovie adds a new movie to the list
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set content type to JSON
	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie) // Decode incoming JSON request body
	if err != nil {
		http.Error(w, "Invalid movie data", http.StatusBadRequest) // If decoding fails, return 400
		return
	}
	movie.ID = uuid.New().String()    // Generate unique ID for the new movie
	movies = append(movies, movie)    // Add the new movie to the list
	w.WriteHeader(http.StatusCreated) // Return 201 status for created resource
	json.NewEncoder(w).Encode(movie)  // Return the created movie
}

// updateMovie updates an existing movie by ID
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set content type to JSON
	params := mux.Vars(r)                              // Get URL params
	for index, item := range movies {                  // Iterate over movies
		if item.ID == params["id"] { // If ID matches, update movie
			var movie Movie
			err := json.NewDecoder(r.Body).Decode(&movie) // Decode incoming JSON request body
			if err != nil {
				http.Error(w, "Invalid movie data", http.StatusBadRequest) // If decoding fails, return 400
				return
			}
			movie.ID = params["id"]          // Ensure the ID remains the same as the original
			movies[index] = movie            // Update the movie in place
			json.NewEncoder(w).Encode(movie) // Return the updated movie
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound) // If movie is not found, return 404
}

// main function to define routes and start the server
func main() {
	r := mux.NewRouter() // Create a new Gorilla Mux router

	// Seed some initial movie data
	movies = append(movies, Movie{
		ID:    uuid.New().String(),
		Isbn:  "123456",
		Title: "Inception",
		Director: &Director{
			Firstname: "Christopher",
			Lastname:  "Nolan",
		},
	})

	movies = append(movies, Movie{
		ID:    uuid.New().String(),
		Isbn:  "654321",
		Title: "The Godfather",
		Director: &Director{
			Firstname: "Francis",
			Lastname:  "Coppola",
		},
	})

	movies = append(movies, Movie{
		ID:    uuid.New().String(),
		Isbn:  "112233",
		Title: "Pulp Fiction",
		Director: &Director{
			Firstname: "Quentin",
			Lastname:  "Tarantino",
		},
	})

	movies = append(movies, Movie{
		ID:    uuid.New().String(),
		Isbn:  "334455",
		Title: "The Dark Knight",
		Director: &Director{
			Firstname: "Christopher",
			Lastname:  "Nolan",
		},
	})

	movies = append(movies, Movie{
		ID:    uuid.New().String(),
		Isbn:  "556677",
		Title: "Fight Club",
		Director: &Director{
			Firstname: "David",
			Lastname:  "Fincher",
		},
	})

	movies = append(movies, Movie{
		ID:    uuid.New().String(),
		Isbn:  "778899",
		Title: "Forrest Gump",
		Director: &Director{
			Firstname: "Robert",
			Lastname:  "Zemeckis",
		},
	})

	// Define route handlers
	r.HandleFunc("/movies", getMovies).Methods("GET")           // Get all movies
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")       // Get a movie by ID
	r.HandleFunc("/movies", createMovie).Methods("POST")        // Create a new movie
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")    // Update a movie by ID
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE") // Delete a movie by ID

	// Start the HTTP server
	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r)) // Start the server on port 8080
}
