package server

import (
	"encoding/json"
	"iSpringTest/models"
	"io"
	"log"
	"net/http"
)

func (s *Server) addTodo(w http.ResponseWriter, r *http.Request) {
	var payload models.AddTodoType
	err := getJsonBody(r, &payload)

	if err != nil {
		JsonResponse404WithBody(w, err.Error())
		return
	}

	err = models.AddTodo(payload, s.database)
	if err != nil {
		JsonResponse404WithBody(w, err.Error())
		return
	}

	JsonResponse201(w)
}

func (s *Server) completeTodo(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) getUncompletedTodo(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) deleteTodo(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) getCompletedTodo(w http.ResponseWriter, r *http.Request) {

}

func getJsonBody(r *http.Request, payload interface{}) error {
	body, err := io.ReadAll(r.Body)

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Println(err)
		}
	}(r.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(body, &payload)
}
