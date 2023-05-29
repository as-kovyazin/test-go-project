package server

import (
	"context"
	"iSpringTest/services"
	"net/http"
)

func (s *Server) addTask(w http.ResponseWriter, r *http.Request) {
	var payload services.RequestTask
	err := getJsonBody(r, &payload)

	if err != nil {
		JsonResponse400WithBody(w, RequestErr{Error: err.Error()})
		return
	}

	service := services.Task{
		Repository: s.taskRepository,
	}

	task, err := service.AddTask(payload)
	if err == services.DatabaseError {
		JsonResponse500(w)
		return
	}
	if err != nil {
		JsonResponse400WithBody(w, RequestErr{Error: err.Error()})
		return
	}

	JsonResponse200WithBody(w, GetResponseTask(task))
}

func (s *Server) updateTask(w http.ResponseWriter, r *http.Request) {
	id, err := getRequestIntVal("id", r)
	if err != nil {
		JsonResponse400WithBody(w, RequestErr{Error: err.Error()})
		return
	}

	service := services.Task{
		Repository: s.taskRepository,
	}

	err = service.CompleteTask(id)
	if err == services.NotFoundByIdError {
		JsonResponse404WithBody(w, RequestErr{Error: err.Error()})
		return
	}
	if err == services.DatabaseError {
		JsonResponse500(w)
		return
	}
	if err != nil {
		JsonResponse400WithBody(w, RequestErr{Error: err.Error()})
		return
	}

	JsonResponse200(w)
}

func (s *Server) getUncompletedTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := s.taskRepository.FindUncompletedTasks(context.Background())
	if err != nil {
		JsonResponse500(w)
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

	service := services.Task{
		Repository: s.taskRepository,
	}

	err = service.DeleteTask(id)
	if err == services.NotFoundByIdError {
		JsonResponse404WithBody(w, RequestErr{Error: err.Error()})
		return
	}
	if err == services.DatabaseError {
		JsonResponse500(w)
		return
	}
	if err != nil {
		JsonResponse400WithBody(w, RequestErr{Error: err.Error()})
		return
	}

	JsonResponse200(w)
}

func (s *Server) getCompletedTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := s.taskRepository.FindCompletedTasks(context.Background())
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

	service := services.Task{
		Repository: s.taskRepository,
	}

	task, err := service.GetTask(id)
	if err == services.NotFoundByIdError {
		JsonResponse404WithBody(w, RequestErr{Error: err.Error()})
		return
	}
	if err != nil {
		JsonResponse400WithBody(w, RequestErr{Error: err.Error()})
		return
	}

	JsonResponse200WithBody(w, GetResponseTask(task))
}
