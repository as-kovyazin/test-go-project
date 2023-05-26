package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"iSpringTest/config"
	"log"
	"net/http"
)

type Server struct {
	config *config.Config
	router *mux.Router
}

func MakeServer(confPath string) *Server {
	conf := config.Load(confPath)

	s := &Server{
		config: conf,
	}
	s.router = s.NewRouter()

	return s
}

func (s *Server) Run() {
	addr := fmt.Sprint("[::]:", s.config.ApiPort)
	err := http.ListenAndServe(addr, s.router)
	if err != nil {
		log.Fatal("Error to load server: ", err)
	}
}
