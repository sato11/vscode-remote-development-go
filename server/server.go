package server

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/sato11/vscode-remote-development-go/model/tasks"
	"github.com/sato11/vscode-remote-development-go/repository"
)

// Server defines struct
type Server struct {
	server     http.Server
	repository tasks.Repository
}

// New initializes instance
func New(addr string, webroot string) *Server {
	s := &Server{
		server: http.Server{
			Addr: addr,
		},
		repository: repository.New(),
	}

	s.setRouter(webroot)

	return s
}

func (s *Server) setRouter(webroot string) {
	router := gin.Default()
	api := router.Group("/api")
	api.GET("/tasks", s.list)
	api.POST("/tasks", s.create)
	api.POST("/tasks/:id/done", s.done)

	router.StaticFile("/", filepath.Join(webroot, "index.html"))
	router.Static("/js", filepath.Join(webroot, "js"))
	s.server.Handler = router
}

// Serve starts http server
func (s *Server) Serve() {
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("could not start server: %s", err.Error())
	}
}
