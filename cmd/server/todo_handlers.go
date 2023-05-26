package server

import (
	"iSpringTest/models"
	"net/http"
)

func (s *Server) addTodo(w http.ResponseWriter, r *http.Request) {
	var payload models.RequestTodo
	err := getJsonBody(r, &payload)

	if err != nil {
		JsonResponse404WithBody(w, RequestErr{Error: err.Error()})
		return
	}

	todo, err := models.AddTodo(payload, s.database)
	if err != nil {
		JsonResponse404WithBody(w, RequestErr{Error: err.Error()})
		return
	}

	JsonResponse200WithBody(w, Todo{ID: todo.ID, Text: todo.Text, CreatedAt: todo.CreatedAt})
}

func (s *Server) completeTodo(w http.ResponseWriter, r *http.Request) {
	id, err := getRequestIntVal("id", r)
	if err != nil {
		JsonResponse404WithBody(w, RequestErr{Error: err.Error()})
		return
	}

	if err := models.CompleteTodo(id, s.database); err != nil {
		JsonResponse404WithBody(w, RequestErr{Error: err.Error()})
		return
	}

	JsonResponse200(w)
}

func (s *Server) getUncompletedTodo(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) deleteTodo(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) getCompletedTodo(w http.ResponseWriter, r *http.Request) {

}
