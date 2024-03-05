package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRequest(c *gin.Context, res *Response) {
	if res.Err == nil {
		statusCode := res.StatusCode
		if statusCode == 0 {
			statusCode = http.StatusOK
		}
		if res.Data != nil {
			c.JSON(res.StatusCode, res.Data)
		} else {
			c.Status(res.StatusCode)
		}
		return
	}

	var err *ErrorResponse
	err, ok := res.Err.(*ErrorResponse)
	if !ok {
		res.StatusCode = http.StatusInternalServerError
		err = &ErrorResponse{Code: InternalServerError, Message: "An error has occurred, please try again later"}
	}
	c.AbortWithStatusJSON(res.StatusCode, err)
}
