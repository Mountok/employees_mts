package handler

import (
	"rest_api_learn/pgk/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
    service *service.Service
}

func NewHandler(service *service.Service) *Handler {
    return &Handler{
        service: service,
    }
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-up", h.SignUp)
	}
	api := router.Group("/api")
	{
		api.POST("/employer", h.CreateEmployer)
		api.POST("/employers", h.ReadEmployer)
		api.PUT("/employers", h.UpadateEmployer)
		api.DELETE("/employers", h.DeleteEmployer)
	}

	return router
}
