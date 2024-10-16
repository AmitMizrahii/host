package api

import (
	"host/internal/api/controllers"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	port           string
	userController controllers.UserController
}

func NewServer(port string, userController controllers.UserController) *Server {
	return &Server{port: port, userController: userController}
}

func (s *Server) Init(logger *slog.Logger) (err error) {
	r := mux.NewRouter()

	s.userController.InitRoutes(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + s.port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("Server is running", "port", s.port)

	return srv.ListenAndServe()
}
