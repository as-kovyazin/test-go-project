package server

import (
	"context"
	"iSpringTest/database"
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

	JsonResponse200WithBody(w, ResponseTodo{ID: todo.ID, Text: todo.Text, CreatedAt: todo.CreatedAt})
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

func (s *Server) getUncompletedTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := database.FindUncompletedTodos(s.database, context.Background())
	if err != nil {
		JsonResponse404WithBody(w, RequestErr{Error: err.Error()})
		return
	}

	JsonResponse200WithBody(w, getResponseTodoList(todos))
}

func (s *Server) deleteTodo(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) getCompletedTodo(w http.ResponseWriter, r *http.Request) {
	todos, err := database.FindCompletedTodos(s.database, context.Background())
	if err != nil {
		JsonResponse404WithBody(w, RequestErr{Error: err.Error()})
		return
	}

	JsonResponse200WithBody(w, getResponseTodoList(todos))
}

func getResponseTodoList(todos []database.Todo) ResponseTodoList {
	responseTodos := make([]ResponseTodo, 0)
	for _, todo := range todos {
		responseTodos = append(responseTodos, ResponseTodo{ID: todo.ID, Text: todo.Text, CreatedAt: todo.CreatedAt})
	}
	return ResponseTodoList{Todos: responseTodos}
}
