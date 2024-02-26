package controllers

import(
  "github.com/gin-gonic/gin"
  "net/http"
)

func GetTables() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "Get tables",
    })
  }
}

func GetTable() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "Get table by id",
    })
  }
}

func CreateTable() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "create table",
    })
  }
}

func UpdateTable() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "update table",
    })
  }
}