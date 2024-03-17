package controllers

import(
  "time"
  "context"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/sajagsubedi/Restaurant-Management-Api/models"
)

func GetOrders() gin.HandlerFunc {  return func(c *gin.Context) {
    ctx,cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    orders,err:= models.GetOrdersDb(ctx)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "success":false,"message": "Failed to fetch orders",
      })
    }
    if orders == nil {
      c.JSON(http.StatusOK, gin.H {
        "orders": [0]models.Order {},
      })
      return
    }
    c.JSON(http.StatusOK, gin.H {
      "success": true,
      "message": "Fetch orders successfully",
      "orders": orders,
    })
  }

}

func GetOrder() gin.HandlerFunc {  return func(c *gin.Context) {
    ctx,
    cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    orderId:= c.Param("orderid")
    orderInfo,err:= models.GetOrderById(ctx, orderId)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "success":false,
        "message": err.Error(),
      })
      return
    }
   orderItems,err:= models.GetOrderItemsByOrderId(ctx, orderId)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
      "success":false,
        "message": err.Error(),
      })
      return
    }
    if orderItems == nil {
      c.JSON(http.StatusOK, gin.H {
        "success": true, "message": "Fetched order successfully", "orderInfo": orderInfo, "orderItems": [0]models.OrderItem {},})
    }
    c.JSON(http.StatusOK, gin.H {
      "success": true, "message": "Fetched order successfully", "orderInfo": orderInfo,
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