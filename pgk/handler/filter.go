package handler

import (
	"net/http"
	"rest_api_learn/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ReadAllFiltersDate(c *gin.Context)  {



	allPosiotions, err := ReadAllPositions(c)
	if err != nil {
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}

	c.JSON(http.StatusOK, models.FiltersResponse{
		Position: allPosiotions,
	}	)



}



func ReadAllPositions(c *gin.Context) (result map[string]string, err error) {
	result = map[string]string {
		"value": "1",
		"label": "Директор",
	}
	return result, nil
}