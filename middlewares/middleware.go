package middlewares

import (
	"log"
	"net/http"
	"resume/models"

	"resume/utils/token"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		b := c.Request.Header.Get("Authorization")
		if b == "" {
			log.Println(models.Response{Code: "401", Message: "Missing Authorization Header"})
			c.JSON(http.StatusUnauthorized, models.Response{Code: "401", Message: "Missing Authorization Header"})
			c.Abort()
			return
		} else {
			err := token.TokenValidApi(c)
			if err != nil {
				c.JSON(http.StatusUnauthorized, models.Response{Code: "401", Message: err.Error()})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
