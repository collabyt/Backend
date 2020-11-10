package main

import (
	"net/http"

	"time"

	"github.com/collabyt/Backend/database"
	"github.com/collabyt/Backend/handler"
	"github.com/collabyt/Backend/limiter"
	"github.com/gorilla/mux"
)

func main() {
	// Database connection pool
	database.Connect()

	// API routes
	r := mux.NewRouter()

	// Playlist operations
	r.HandleFunc("/api/v1/playlists", handler.CreatePlaylist).Methods("POST")
	r.HandleFunc("/api/v1/playlists/{PublicID}", handler.GetPlaylistByPublicID).Methods("GET")
	r.HandleFunc("/api/v1/playlists", handler.GetPublicPlaylists).Methods("GET")
	// Private Playlist operations
	r.HandleFunc("/api/v1/auth/{PublicID}", handler.RequestAccessToPlaylist).Methods("POST")
	r.HandleFunc("/api/v1/exit/{PublicID}", handler.DeauthorizeToPlaylist).Methods("GET")
	// Keyword operations
	r.HandleFunc("/api/v1/keywords", handler.CreateKeyword).Methods("POST")
	r.HandleFunc("/api/v1/keywords", handler.GetKeywords).Methods("GET")
	// Video operations
	r.HandleFunc("/api/v1/playlists/{PublicID}/videos", handler.CreateVideoInPlaylist).Methods("POST")
	r.HandleFunc("/api/v1/playlists/{PublicID}/videos/{VideoID}", handler.DeleteVideo).Methods("DELETE")
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	visitors := limiter.NewCatalog(256, 5, 10) // TODO: Change to environment variabes
	timeToDumpVisitor := 60 * time.Minute      // TODO: Change to environment variabe
	go limiter.CleanupVisitors(visitors, timeToDumpVisitor)

	server := http.Server{
		Addr:         ":8080",
		Handler:      limiter.Limit(visitors, r),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
