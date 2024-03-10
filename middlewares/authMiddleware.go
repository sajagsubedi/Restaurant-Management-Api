package middlewares

import (
	"net/http"
	"github.com/sajagsubedi/Restaurant-Management-Api/helpers"
	"github.com/gin-gonic/gin"
)

func CheckUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.Request.Header.Get("auth_token")
		if authToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "No Authorization header provided",
			})
			c.Abort()
			return
		}
		claims, err := helpers.ValidateToken(authToken)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Please provide correct authorization header!",
			})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("userid", claims.Userid)

		c.Next()
	}
}