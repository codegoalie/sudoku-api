package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(repo, http.ResponseWriter, *http.Request)
}

type Routes []Route

var routes = Routes{
	Route{"PuzzleIndex", "GET", "/puzzle", PuzzleIndex},
	Route{"PuzzleShow", "GET", "/puzzle/{id}", PuzzleShow},
	Route{"PuzzleCreate", "POST", "/puzzle", PuzzleCreate},
}
