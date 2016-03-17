package main

import "net/http"

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(repo, http.ResponseWriter, *http.Request)
}

type routes []route

var ourRoutes = routes{
	route{"PuzzleIndex", "GET", "/puzzle", puzzleIndex},
	route{"PuzzleShow", "GET", "/puzzle/{id}", puzzleShow},
	route{"PuzzleCreate", "POST", "/puzzle", puzzleCreate},
	route{"StatsIndex", "GET", "/stats", statsIndex},
}
