package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ReadAllFiltersDate(c *gin.Context)  {

	filters, err := h.service.ReadAllFiltersDate()
	if err != nil {
		newErrorResponse(c,http.StatusOK, err.Error())
		return
	}

	c.JSON(http.StatusOK,filters)



}
