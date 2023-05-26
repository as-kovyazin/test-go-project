package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"iSpringTest/database"
	"io"
	"log"
	"net/http"
	"strconv"
)

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

func getRequestIntVal(param string, r *http.Request) (int64, error) {
	str := mux.Vars(r)[param]
	return strconv.ParseInt(str, 10, 64)
}

func getResponseTaskList(tasks []database.Task) ResponseTaskList {
	responseTasks := make([]ResponseTask, 0)
	for _, task := range tasks {
		responseTasks = append(responseTasks, ResponseTask{ID: task.ID, Text: task.Text, CreatedAt: task.CreatedAt})
	}
	return ResponseTaskList{Tasks: responseTasks}
}
