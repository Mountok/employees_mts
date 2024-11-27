package restapi

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr: ":" + os.Getenv("PORT"),
		Handler: handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logrus.Info(s.httpServer.Addr)

	return s.httpServer.ListenAndServe()
}


func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}