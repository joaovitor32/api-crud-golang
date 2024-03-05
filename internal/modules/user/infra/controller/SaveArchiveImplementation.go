package controller

import (
	"api-crud-golang/internal/shared/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveArchiveImplementation(c *gin.Context) types.RequestBody {
	var body types.RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request in body",
		})
		c.Abort()
		return body
	}

	return body
}
