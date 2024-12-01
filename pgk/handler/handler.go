package handler

import (
	"rest_api_learn/pgk/service"
	"time"

	"github.com/gin-contrib/cors"
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

	config := cors.DefaultConfig()
    config.AllowOrigins = []string{"*"}
    config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
    config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
    config.ExposeHeaders = []string{"Content-Length"}
    config.AllowCredentials = true
	config.ExposeHeaders = []string{}
    config.MaxAge = 12 * time.Hour

    router.Use(cors.New(config))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-up", h.SignUp)
	}
	api := router.Group("/api")
	{
		api.GET("/employer", h.CreateEmployer)
		api.POST("/employers", h.ReadEmployer)
		// api.OPTIONS("/employers", h.ReadEmployer)
		api.PUT("/employers", h.UpadateEmployer)
		api.DELETE("/employers", h.DeleteEmployer)
		api.GET("/filters", h.ReadAllFiltersDate)
	}

	return router
}
