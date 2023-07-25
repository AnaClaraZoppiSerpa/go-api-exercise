package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type Album struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

var albums = []Album{
	{1, "Kind of Blue", "Miles Davis", 1959},
	{2, "A Love Supreme", "John Coltrane", 1965},
	{3, "Blue Train", "John Coltrane", 1957},
}

func getAlbumsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(albums)
}

func addAlbumHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newAlbum Album
	err := json.NewDecoder(r.Body).Decode(&newAlbum)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	// Add the new album to the list
	albums = append(albums, newAlbum)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newAlbum)
}
/*
func main() {
	fmt.Println("Starting albums server")
	http.HandleFunc("/albums", getAlbumsHandler)
	http.HandleFunc("/new", addAlbumHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
*/
