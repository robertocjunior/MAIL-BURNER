package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"tempmail/internal/database"
	"tempmail/internal/handlers"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	database.InitDB()

	http.Handle("/", http.FileServer(http.Dir("./static")))

	// API Routes
	http.HandleFunc("/api/config", handlers.HandleConfig)
	http.HandleFunc("/api/destinations", handlers.HandleDestinations)
	http.HandleFunc("/api/check", handlers.HandleCheck)
	http.HandleFunc("/api/create", handlers.HandleCreate)
	http.HandleFunc("/api/pin", handlers.HandlePin)
	http.HandleFunc("/api/active", handlers.HandleListActive)
	http.HandleFunc("/api/history", handlers.HandleHistory)
	http.HandleFunc("/api/delete", handlers.HandleDelete)
	http.HandleFunc("/api/tags", handlers.HandleTags)

	addr := ":" + port
	fmt.Printf("🚀 Sistema Cloudflare Mail v6.1 Refatorado rodando em http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}