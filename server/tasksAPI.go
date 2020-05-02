package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sato11/vscode-remote-development-go/model/tasks"
)

func (s *Server) list(c *gin.Context) {
	tasks := s.repository.List()
	c.JSON(http.StatusOK, tasks)
}

func (s *Server) create(c *gin.Context) {
	var task tasks.Task
	c.ShouldBindJSON(&task)
	task.ID = s.repository.Add(task)
	c.JSON(200, &task)
}

func (s *Server) done(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(400)
	}
	err = s.repository.Done(id)
	if err != nil {
		c.Status(404)
	}
	c.Status(200)
}
