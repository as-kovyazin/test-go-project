package server

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func (s *Server) GetRoutes() []Route {
	return Routes{
		Route{"todo_add", "PUT", "/todo/add", nil},
		Route{"todo_complete", "POST", "/todo/complete", nil},
		Route{"todo_get_uncompleted", "GET", "/todo/uncompleted", nil},
		Route{"todo_delete", "GET", "/todo/delete", nil},
		Route{"todo_get_completed", "GET", "/todo/completed", nil},
	}
}
