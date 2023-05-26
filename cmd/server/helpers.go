package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
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
