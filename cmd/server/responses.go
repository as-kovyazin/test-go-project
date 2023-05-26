package server

import (
	"encoding/json"
	"net/http"
)

type RequestErr struct {
	Error string `json:"error"`
}

type ResponseTask struct {
	ID        int64  `json:"id"`
	Text      string `json:"text"`
	CreatedAt int64  `json:"createdAt"`
}

type ResponseTaskList struct {
	Tasks []ResponseTask `json:"tasks"`
}

func JsonResponse200(w http.ResponseWriter) {
	JsonResponse(w, http.StatusOK, nil)
}

func JsonResponse200WithBody(w http.ResponseWriter, rawJson any) {
	JsonResponse(w, http.StatusOK, rawJson)
}

func JsonResponse201(w http.ResponseWriter) {
	JsonResponse(w, http.StatusCreated, nil)
}

func JsonResponse400WithBody(w http.ResponseWriter, rawJson any) {
	JsonResponse(w, http.StatusBadRequest, rawJson)
}

func JsonResponse404(w http.ResponseWriter) {
	JsonResponse(w, http.StatusNotFound, nil)
}

func JsonResponse404WithBody(w http.ResponseWriter, rawJson any) {
	JsonResponse(w, http.StatusNotFound, rawJson)
}

func JsonResponse(w http.ResponseWriter, statusCode int, rawJson any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if rawJson != nil {
		if err := json.NewEncoder(w).Encode(rawJson); err != nil {
			return
		}
	}
	return
}
