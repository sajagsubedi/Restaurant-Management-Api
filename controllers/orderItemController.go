package controllers

import(
  "fmt"
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


func CreateOrderItem() gin.HandlerFunc {  return func(c *gin.Context) {
    ctx,cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    var orderitem models.OrderItem
    if err:= c.BindJSON(&orderitem); err != nil {
      c.JSON(http.StatusBadRequest, gin.H {
        "success":false,"message": err.Error()})
      return
    }
    validationErr:= validate.Struct(orderitem)
    if validationErr != nil {
      c.JSON(http.StatusBadRequest, gin.H {
        "success":false,"message": validationErr.Error()})
      return
    }
  if  _,err:=models.GetFoodById(ctx,fmt.Sprintf("%v",*orderitem.Food_id));err!=nil{
     c.JSON(http.StatusBadRequest, gin.H {
        "success":false,"message": err.Error()})
      return
    
  }
  
    if  _,err:=models.GetOrderById(ctx,fmt.Sprintf("%d",*orderitem.Order_id));err!=nil{
     c.JSON(http.StatusBadRequest, gin.H {
        "success":false,"message": err.Error()})
      return
    
  }
  
    _, err:= models.GetOrderById(ctx, 
      fmt.Sprintf("%d", *orderitem.Order_id))
      if err != nil {
        c.JSON(http.StatusBadRequest, gin.H {
          "success": false, "message": err.Error(),
        })
        return
      }
      createdOrderItem,err:= models.CreateOrderItemDB(ctx, orderitem)
      if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H {
          "success": false, "message": "Failed to add orderitem",
        })
        return
      }
      c.JSON(http.StatusCreated, gin.H {
        "success": true,
        "message": "Created order item successfully!",
        "orderitem": createdOrderItem,
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