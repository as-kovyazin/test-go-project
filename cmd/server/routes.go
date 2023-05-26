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
		Route{"todo_complete", "POST", "/todo/complete/{id:[0-9]+}", s.completeTodo},
		Route{"todo_delete", "DELETE", "/todo/delete/{id:[0-9]+}", s.deleteTodo},
		Route{"todo_get_uncompleted", "GET", "/todo/uncompleted", s.getUncompletedTodos},
		Route{"todo_get_completed", "GET", "/todo/completed", s.getCompletedTodo},
	}
}
