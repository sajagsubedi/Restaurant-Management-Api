package controllers

import(
  "github.com/gin-gonic/gin"
  "net/http"
    
)

func GetOrders() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "Get orders",
    })
  }
}

func GetOrder() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "Get order by id",
    })
  }
}

func CreateOrder() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "create orders",
    })
  }
}

func UpdateOrder() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "update orders",
    })
  }
}