package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

func newRouter(repo repo) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range ourRoutes {
		var handler http.Handler

		handler = handlerWithRepo(repo, route.HandlerFunc)
		handler = logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	router.Handle("/metrics", prometheus.Handler())

	return router
}

func handlerWithRepo(repo repo, handlerFunc func(repo, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handlerFunc(repo, w, r)
	}
}
