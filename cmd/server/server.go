package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/uptrace/bun"
	"iSpringTest/config"
	"iSpringTest/database"
	"log"
	"net/http"
)

type Server struct {
	config   *config.Config
	router   *mux.Router
	database *bun.DB
}

func MakeServer(confPath string) *Server {
	conf := config.Load(confPath)

	s := &Server{
		config:   conf,
		database: database.Init(conf.PostgresURL, conf.DebugDb),
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
