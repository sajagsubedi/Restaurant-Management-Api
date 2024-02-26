package controllers

import(
  "github.com/gin-gonic/gin"
  "net/http"
)

func GetUsers() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "Get users",
    })
  }
}

func GetUser() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "Get user",
    })
  }
}

func Login() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "login",
    })
  }
}

func Signup() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "signup",
    })
  }
}

func UpdateProfile() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "update profile",
    })
  }
}