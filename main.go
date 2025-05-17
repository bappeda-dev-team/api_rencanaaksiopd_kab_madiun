package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"renaksiopdService/helper"
	"renaksiopdService/middleware"

	"github.com/joho/godotenv"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	host := os.Getenv("host")
	port := os.Getenv("port")
	addr := fmt.Sprintf("%s:%s", host, port)

	cors := helper.NewCORSMiddleware()

	if addr == ":" {
		addr = "localhost:8080"
	}

	return &http.Server{
		Addr:    addr,
		Handler: cors.Handler(authMiddleware),
	}
}

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	// Initialize dan jalankan server
	server := InitializeServer()
	log.Printf("Server berjalan di %s", server.Addr)
	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
