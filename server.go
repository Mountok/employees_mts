package restapi

import (
	"context"
	"net/http"
	// "os"
	"time"

	"github.com/gorilla/handlers"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr: ":" + "8080",
		// Addr: ":" + os.Getenv("PORT"),
		Handler: handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{
				"POST",
				"GET",
				"UPDATE",
				"DELETE",
			}),
			handlers.AllowedHeaders([]string{
				"X-Requested-With",
				"Content-Type",
				"Authorization",
			}),
		)(handler),
		MaxHeaderBytes: 1 << 20,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}


func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}