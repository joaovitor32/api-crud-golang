package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func JwtTokenCheck(c *gin.Context) {
	apiToken:=c.Request.Header["Token"][0]

	if apiToken != os.Getenv("API_TOKEN") {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "Error", "message": "Invalid API token"})
		c.Abort()
		return
	}
}
