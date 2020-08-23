package main

import (
	"net/http"

	"github.com/collabyt/Backend/database"
	"github.com/collabyt/Backend/handler"
	"github.com/gorilla/mux"
)

func main() {
	// Database connection pool
	database.Connect()

	// API routes
	r := mux.NewRouter()

	// Single Video operations
	// r.HandleFunc("/api/v1/video", handler.Video).Methods("POST")
	// Multiple playlists operations
	// r.HandleFunc("/api/v1/playlists/", handler.LatestPlaylists).Methods("GET") // (keyword, afterid)

	// Single playlist operations
	r.HandleFunc("/api/v1/playlist", handler.CreatePlaylist).Methods("POST")
	r.HandleFunc("/api/v1/playlist", handler.GetPlaylist).Methods("GET")
	// Single Keyword operations
	r.HandleFunc("/api/v1/keyword", handler.CreateKeyword).Methods("POST")
	// Multiple Keywords operations
	r.HandleFunc("/api/v1/keywords/", handler.GetKeywords).Methods("GET")
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
