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
	r.HandleFunc("/api/v1/playlist/new", handler.NewPlaylist)
	r.HandleFunc("/api/v1/playlist/:playlist", handler.Playlist)
	r.HandleFunc("/api/v1/playlist/add", handler.AddToPlaylist)
	// Multiple playlists operations
	r.HandleFunc("/api/v1/playlists/latest/:lastid", handler.LatestPlaylists)
	r.HandleFunc("/api/v1/playlists/latest/:keyword", handler.LatestPlaylists)
	// Single Keyword operations
	r.HandleFunc("/api/v1/keyword/new", handler.NewKeyword)
	// Multiple Keywords operations
	r.HandleFunc("/api/v1/keywords/:word", handler.GetKeywords)
	// Serve static Pages
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	// Server initialization
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
