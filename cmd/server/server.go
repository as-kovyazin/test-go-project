package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"iSpringTest/config"
	"iSpringTest/database"
	"iSpringTest/repositories"
	"log"
	"net/http"
)

type Server struct {
	config         *config.Config
	router         *mux.Router
	taskRepository *repositories.Task
}

func MakeServer(confPath string) *Server {
	conf := config.Load(confPath)
	db := database.Init(conf.PostgresURL, conf.DebugDb)
	s := &Server{
		config:         conf,
		taskRepository: repositories.CreateTaskRepository(db),
	}
	s.router = s.NewRouter()

	http.Handle("/", s.router)

	return s
}

func (s *Server) Run() {
	addr := fmt.Sprint("[::]:", s.config.ApiPort)
	log.Println("server up")
	err := http.ListenAndServe(addr, s.router)
	if err != nil {
		log.Fatal("Error to load server: ", err)
	}
}
