package controllers

import(
  "github.com/gin-gonic/gin"
  "net/http"

)

func GetMenus() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "Get menus",
    })
  }
}

func GetMenu() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "get menu by id",
    })
  }
}

func CreateMenu() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "create menu",
    })
  }
}

func UpdateMenu() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "Update menu",
    })
  }
}