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

	// Single playlist operations
	r.HandleFunc("/api/v1/playlist/:playlist", handler.Playlist) // GET & POST
	// r.HandleFunc("/api/v1/video", handler.Video)                 // POST

	// Multiple playlists operations
	// r.HandleFunc("/api/v1/playlists/", handler.LatestPlaylists) // GET (keyword, afterid)

	// Single Keyword operations
	r.HandleFunc("/api/v1/keyword", handler.Keyword) // POST

	// Multiple Keywords operations
	r.HandleFunc("/api/v1/keywords/", handler.GetKeywords) // GET

	// Serve static Pages
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// Server initialization
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
