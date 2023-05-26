package server

import (
	"context"
	"iSpringTest/database"
	"iSpringTest/models"
	"net/http"
)

func (s *Server) addTask(w http.ResponseWriter, r *http.Request) {
	var payload models.RequestTask
	err := getJsonBody(r, &payload)

	if err != nil {
		JsonResponse400WithBody(w, RequestErr{Error: err.Error()})
		return
	}

	task, err := models.AddTask(payload, s.database)
	if err != nil {
		JsonResponse400WithBody(w, RequestErr{Error: err.Error()})
		return
	}

	JsonResponse200WithBody(w, ResponseTask{ID: task.ID, Text: task.Text, CreatedAt: task.CreatedAt})
}

func (s *Server) updateTask(w http.ResponseWriter, r *http.Request) {
	id, err := getRequestIntVal("id", r)
	if err != nil {
		JsonResponse400WithBody(w, RequestErr{Error: err.Error()})
		return
	}

	err = models.CompleteTask(id, s.database)
	if err == models.NotFoundByIdError {
		JsonResponse404(w)
		return
	}
	if err != nil {
		JsonResponse400WithBody(w, RequestErr{Error: err.Error()})
		return
	}

	JsonResponse200(w)
}

func (s *Server) getUncompletedTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := database.FindUncompletedTasks(s.database, context.Background())
	if err != nil {
		JsonResponse400WithBody(w, RequestErr{Error: err.Error()})
		return
	}

	JsonResponse200WithBody(w, getResponseTaskList(tasks))
}

func (s *Server) deleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := getRequestIntVal("id", r)
	if err != nil {
		JsonResponse400WithBody(w, RequestErr{Error: err.Error()})
		return
	}

	err = models.DeleteTask(id, s.database)
	if err == models.NotFoundByIdError {
		JsonResponse404(w)
		return
	}
	if err != nil {
		JsonResponse400WithBody(w, RequestErr{Error: err.Error()})
		return
	}

	JsonResponse200(w)
}

func (s *Server) getCompletedTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := database.FindCompletedTasks(s.database, context.Background())
	if err != nil {
		JsonResponse400WithBody(w, RequestErr{Error: err.Error()})
		return
	}

	JsonResponse200WithBody(w, getResponseTaskList(tasks))
}

func (s *Server) getTask(w http.ResponseWriter, r *http.Request) {
	id, err := getRequestIntVal("id", r)
	if err != nil {
		JsonResponse400WithBody(w, RequestErr{Error: err.Error()})
		return
	}

	task, err := models.GetTask(id, s.database)
	if err == models.NotFoundByIdError {
		JsonResponse404(w)
		return
	}
	if err != nil {
		JsonResponse400WithBody(w, RequestErr{Error: err.Error()})
		return
	}

	JsonResponse200WithBody(w, ResponseTask{ID: task.ID, Text: task.Text, CreatedAt: task.CreatedAt})
}
