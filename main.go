package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	repo := NewRedisRepo(os.Getenv("REDIS_ADDR"))
	router := NewRouter(repo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
}
