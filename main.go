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
	duplicate = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "sudoku",
		Subsystem: "api",
		Name:      "puzzles_duplicate",
		Help:      "The number of puzzles sent that already exist.",
	})
)

func init() {
	prometheus.MustRegister(added)
	prometheus.MustRegister(duplicate)
}

func main() {
	repo := newRedisRepo(os.Getenv("REDIS_ADDR"))
	router := newRouter(repo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
}
