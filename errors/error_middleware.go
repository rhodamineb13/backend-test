package customerrors

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware(c *gin.Context) {
	c.Next()
	for _, err := range c.Errors {
		var e Errors
		if errors.As(err, &e) {
			c.AbortWithStatusJSON(int(e.StatusCode), e.Message)
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, "unexpected server error")
		}
	}
}
