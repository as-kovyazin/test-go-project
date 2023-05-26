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
		Route{"todo_add", "PUT", "/todo/add", s.addTodo},
		Route{"todo_complete", "POST", "/todo/complete", s.completeTodo},
		Route{"todo_get_uncompleted", "GET", "/todo/uncompleted", s.getUncompletedTodo},
		Route{"todo_delete", "GET", "/todo/delete", s.deleteTodo},
		Route{"todo_get_completed", "GET", "/todo/completed", s.getCompletedTodo},
	}
}
