package middleware

import (
	error2 "channel-contents-collector/api/v1/error"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			var e error2.Http
			switch {
			case errors.As(err.Err, &e):
				c.AbortWithStatusJSON(e.StatusCode, e)
				return
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": "Service Unavailable"})
				return
			}
		}
	}
}
