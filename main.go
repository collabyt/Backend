package main

import (
	"log"
	"net/http"

	"github.com/collabyt/Backend/cache"
	"github.com/collabyt/Backend/database"
	"github.com/collabyt/Backend/handler"
	"github.com/collabyt/Backend/logger"
	"github.com/collabyt/Backend/startup"
	"github.com/gorilla/mux"
)

func main() {
	// initializing Application Logs
	logger.Setup()
	// Database connection pool
	database.Connect()
	log.Println("Server connected to the PostgreSQL Database")
	// TODO: Implement POSTGRES connection successfully log
	cache.Connect()
	log.Println("Server connected to the REDIS Cache")
	// TODO: Implement REDIS connection successfully log

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

	server, err := startup.SetupServer(r)
	if err != nil {
		panic(err)
	}
	log.Println("http.Server successfully created")
	// TODO: Implement initialization error log

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	log.Println("Server running and accepting connections...")
	// TODO: Implement successful start log
}
