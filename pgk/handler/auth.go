package handler

import (
	"net/http"
	"rest_api_learn/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {

	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}
	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return 
	}

	c.JSON(http.StatusOK,map[string]interface{}{
		"data": id,
	})
	
} 

func (h *Handler) SignIn(c *gin.Context) {
	c.JSON(http.StatusOK,map[string]interface{}{
		"Status": "OK",
	})
} 