package controllers

import(
  "github.com/gin-gonic/gin"
  "net/http"
    
)

func GetOrderItems() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "Get order items",
    })
  }
}

func GetOrderItem() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "Get orders by id ",
    })
  }
}

func GetOrderItemsByOrder() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "Get orders by item",
    })
  }
}

func CreateOrderItem() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "create orderItem",
    })
  }
}

func UpdateOrderItem() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "update orderitem",
    })
  }
}