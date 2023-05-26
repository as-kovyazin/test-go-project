package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (s *Server) NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range s.GetRoutes() {
		var handler http.Handler

		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	router.Use(middleware)

	return router
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
