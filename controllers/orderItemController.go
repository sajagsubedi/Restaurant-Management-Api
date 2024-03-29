package controllers

import(
  "fmt"
  "time"
  "strings"
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

func GetOrderItem() gin.HandlerFunc {
  return func(c *gin.Context) {
    ctx,
    cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    orderItemId:= c.Param("orderitemid")
    orderItem,
    err:= models.GetOrderItemById(ctx, orderItemId)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "success": false,
        "message": err.Error(),
      })
      return
    }
    food,
    err:= models.GetFoodById(ctx, strconv.FormatInt(*orderItem.Food_id, 10))
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "success": false,
        "message": err.Error(),
      })
      return
    }
    c.JSON(http.StatusOK, gin.H {
      "success": true, "message": "Fetched orderitem", "orderitem": orderItem, "food": food,
    })
  }
}


func CreateOrderItem() gin.HandlerFunc {
  return func(c *gin.Context) {
    ctx,
    cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    var orderitem models.OrderItem
    if err:= c.BindJSON(&orderitem); err != nil {
      c.JSON(http.StatusBadRequest, gin.H {
        "success": false, "message": err.Error()})
      return
    }
    validationErr:= validate.Struct(orderitem)
    if validationErr != nil {
      c.JSON(http.StatusBadRequest, gin.H {
        "success": false, "message": validationErr.Error()})
      return
    }
    if _,
    err:= models.GetFoodById(ctx, fmt.Sprintf("%v", *orderitem.Food_id)); err != nil {
      c.JSON(http.StatusBadRequest, gin.H {
        "success": false, "message": err.Error()})
      return

    }

     order,err:= models.GetOrderById(ctx, fmt.Sprintf("%d", *orderitem.Order_id))
      if err != nil {
      c.JSON(http.StatusBadRequest, gin.H {
        "success": false, "message": err.Error()})
      return
    }
     userid,_:=c.Get("userid")
   if *order.UserId!=userid.(int64){
     c.JSON(http.StatusUnauthorized,gin.H{
       "success":false,
       "message":"You are not authorized to add orderitem in the order wth given id! ",
     })
     return
   }
    createdOrderItem,
    err:= models.CreateOrderItemDB(ctx, orderitem)
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
    ctx,
    cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var orderitem models.OrderItem

    orderItemId:= c.Param("orderitemid")
    userid, _ := c.Get("userid")
        
    if err:= c.BindJSON(&orderitem); err != nil {
      c.JSON(http.StatusBadRequest, gin.H {
        "success": false, "message": err.Error(),
      })
      return
    }
    
    foundorderItem,err:=models.GetOrderItemById(ctx,orderItemId)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "success": false,
        "message": err.Error(),
      })
      return
    }
    
    order,_:=models.GetOrderById(ctx, strconv.FormatInt(*foundorderItem.Order_id,10))
    
    if *order.UserId!=userid.(int64){
      c.JSON(http.StatusUnauthorized,gin.H{
        "success":false,
        "message":"You are no authorized to update the orderitem!",
      })
      return
    }
    var updateObj []string
    var values []interface {}
    
    if orderitem.Quantity != nil {
      if *foundorderItem.Status !="Not_Started"{
        c.JSON(http.StatusBadRequest,gin.H{
          "success":false,
          "message":"Sorry, You can't update the quantity since the cooking of food has already been started or been fulfilled!",
        })
        return
      }
      updateObj = append(updateObj, fmt.Sprintf("quantity=$%d", len(values)+1))
      values = append(values, *orderitem.Quantity)
    }
    
    if orderitem.Status != nil {
      updateObj = append(updateObj, fmt.Sprintf("status=$%d", len(values)+1))
      values = append(values, *orderitem.Status)
    }

    values = append(values, orderItemId)

    setVal:= strings.Join(updateObj, ", ")

    err= models.UpdateOrderItemDb(ctx, setVal, values)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "success": false, "message": err.Error(),
      })
      return
    }

    c.JSON(http.StatusOK, gin.H {
      "success": true, "message": "OrderItem updated successfully",
    })
  }
}