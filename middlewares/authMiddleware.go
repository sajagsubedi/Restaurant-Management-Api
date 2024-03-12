package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sajagsubedi/Restaurant-Management-Api/helpers"
)

func DecodeJwt(c *gin.Context) (*helpers.SignedDetails, int, string) {
	authToken := c.Request.Header.Get("auth_token")
	if authToken == "" {
		msg := "No Authorization header provided"
		return nil, http.StatusUnauthorized, msg
	}

	claims, err := helpers.ValidateToken(authToken)
	if err != nil {
		msg := "Please provide correct authorization header!"
		return nil, http.StatusUnauthorized, msg
	}

	return claims, 0, ""
}

func CheckUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, status, msg := DecodeJwt(c)
		if msg != "" {
			c.JSON(status, gin.H{
				"success": false, "message": msg,
			})
			c.Abort()
			return
		}

		c.Set("email", *claims.Email)
		c.Set("first_name", *claims.First_name)
		c.Set("last_name", *claims.Last_name)
		c.Set("userid", *claims.Userid)
		c.Set("usertype", *claims.UserType)
		c.Next()
	}
}

func CheckAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, status, msg := DecodeJwt(c)
		if msg != "" {
			c.JSON(status, gin.H{
				"success": false, "message": msg,
			})
			c.Abort()
			return
		}
		if *claims.UserType != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "You are not authorized!",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func CheckAdminAndSetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, status, msg := DecodeJwt(c)
		if msg != "" {
			c.JSON(status, gin.H{
				"success": false, "message": msg,
			})
			c.Abort()
			return
		}

		if *claims.UserType != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "You are not authorized!",
			})
			c.Abort()
			return
		}

		c.Set("userid", c.Param("userid"))
		c.Next()
	}
}
