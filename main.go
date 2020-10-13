package main

import (
	"net/http"

	"time"

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
	// r.HandleFunc("/api/v1/video", handler.CreateVideo).Methods("POST")
	// Multiple playlists operations
	// r.HandleFunc("/api/v1/playlists/", handler.LatestPlaylists).Methods("GET") // (keyword, afterid)

	// Playlist operations
	r.HandleFunc("/api/v1/playlists", handler.CreatePlaylist).Methods("POST")             // DONE
	r.HandleFunc("/api/v1/playlists/{pid}", handler.GetPlaylistByPublicID).Methods("GET") // DONE
	r.HandleFunc("/api/v1/playlists", handler.GetPlaylists).Methods("GET")                // DONE
	// Keyword operations
	r.HandleFunc("/api/v1/keywords", handler.CreateKeyword).Methods("POST") // DONE
	r.HandleFunc("/api/v1/keywords/", handler.GetKeywords).Methods("GET")   // DONE
	// Video operations
	r.HandleFunc("/api/v1/playlists/{pid}/videos", handler.CreateVideoInPlaylist).Methods("POST") // DONE
	r.HandleFunc("/API/v1/playlists/{pid}/videos/{vid}", handler.DeleteVideo).Methods("DELETE")   // TODO: CHECK MODEL!
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	server := http.Server{
		Addr:         ":8080",
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
