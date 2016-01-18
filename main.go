package main

import (
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	added = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "sudoku",
		Subsystem: "api",
		Name:      "puzzles_added",
		Help:      "The number of puzzles added.",
	})
)

func init() {
	prometheus.MustRegister(added)
}

func main() {
	repo := NewRedisRepo(os.Getenv("REDIS_ADDR"))
	router := NewRouter(repo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
}
