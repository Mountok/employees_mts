package handler

import (
	"net/http"
	"rest_api_learn/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateEmployer(c *gin.Context) {}

func (h *Handler) ReadEmployers(c *gin.Context) {
	var input [][]models.EmployersResponse

	h.service.Employees.ReadEmployers()

	c.JSON(http.StatusOK,input)
	
}

func (h *Handler) ReadEmployer(c *gin.Context) {
	var input models.Employers

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}

	result, err := h.service.Employees.ReadEmployer(input)
	if err != nil {
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}

	c.JSON(http.StatusOK,map[string]interface{}{
		"data": result,
	})

}
func (h *Handler) UpadateEmployer(c *gin.Context) {}
func (h *Handler) DeleteEmployer(c *gin.Context) {}