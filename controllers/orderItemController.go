package controllers

import(
  "time"
  "strconv"
  "context"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/sajagsubedi/Restaurant-Management-Api/models"
)

func GetOrderItems() gin.HandlerFunc {
  return func(c *gin.Context) {
    ctx,
    cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    orderItems,
    err:= models.GetOrderItemsDb(ctx)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "success": false, "message": "Failed to fetch order items",
      })
    }
    if orderItems == nil {
      c.JSON(http.StatusOK, gin.H {
        "success": false,
        "orderitems": [0]models.OrderItem {},
      })
      return
    }
    c.JSON(http.StatusOK, gin.H {
      "success": true,
      "message": "Fetch order items  successfully",
      "orderitems": orderItems,
    })
  }

}

func GetOrderItem() gin.HandlerFunc {  return func(c *gin.Context) {
    ctx,
    cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    orderItemId:= c.Param("orderitemid")
    orderItem,err:= models.GetOrderItemById(ctx, orderItemId)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "success":false,
        "message": err.Error(),
      })
      return
    }
    food,err:= models.GetFoodById(ctx, strconv.FormatInt(*orderItem.Food_id,10))
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "success":false,
        "message": err.Error(),
      })
      return
    }
    c.JSON(http.StatusOK, gin.H {
        "success": true, "message": "Fetched orderitem", "orderitem": orderItem, "food":food,})
  }
}


func CreateOrderItem() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H {
      "message": "create orderItem",
    })
  }
}

func UpdateOrderItem() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H {
      "message": "update orderitem",
    })
  }
}