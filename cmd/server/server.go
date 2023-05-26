package server

import (
	"fmt"
	"iSpringTest/config"
	"log"
	"net/http"
)

type Server struct {
	config *config.Config
}

func MakeServer(confPath string) *Server {
	conf := config.Load(confPath)

	s := &Server{
		config: conf,
	}

	return s
}

func (s *Server) Run() {
	addr := fmt.Sprint("[::]:", s.config.ApiPort)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("Error to load server: ", err)
	}
}
