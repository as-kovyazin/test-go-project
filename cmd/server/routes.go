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
		Route{"task_add", "PUT", "/task", s.addTask},
		Route{"task_complete", "POST", "/task/{id:[0-9]+}", s.updateTask},
		Route{"task_delete", "DELETE", "/task/{id:[0-9]+}", s.deleteTask},
		Route{"task_get", "GET", "/task/{id:[0-9]+}", s.getTask},
		Route{"get_uncompleted_task", "GET", "/task/uncompleted", s.getUncompletedTasks},
		Route{"get_completed_task", "GET", "/task/completed", s.getCompletedTasks},
	}
}
