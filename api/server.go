package api

import "github.com/gin-gonic/gin"

type Server struct {
	port           string
	userController UserController
}

func NewServer(port string, userController UserController) *Server {
	return &Server{port: port, userController: userController}
}

func (s *Server) Init() (err error) {
	r := gin.Default()

	s.userController.InitRoutes(r)

	return  r.Run(":" + s.port)
}
