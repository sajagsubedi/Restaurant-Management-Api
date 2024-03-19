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

func CreateOrder() gin.HandlerFunc {  return func(c *gin.Context) {
    ctx,
    cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    var order models.Order
    if err:= c.BindJSON(&order); err != nil {
      c.JSON(http.StatusBadRequest, gin.H {
        "success":false,"message": err.Error()})
      return
    }
    validationErr:= validate.Struct(order)
    if validationErr != nil {
      c.JSON(http.StatusBadRequest, gin.H {
        "success":false,"message": validationErr.Error(),})
      return
    }
    _,err:=models.GetTableById(ctx,strconv.FormatInt(*order.TableId,10))
    if err!=nil{
    c.JSON(http.StatusBadRequest, gin.H {
        "success":false,"message": err.Error(),})
    return
    }
    createdOrder,err:= models.CreateOrderDb(ctx, order)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "success": false, "message": "Failed to add order",
      })
      return
    }
    c.JSON(http.StatusCreated, gin.H {
      "success": true,
      "message": "Created order successfully!",
      "order": createdOrder,
    })
  }

}

func UpdateOrder() gin.HandlerFunc {  
  return func(c *gin.Context) {
    ctx,
    cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var order models.Order

    orderId:= c.Param("orderid")

    if err:= c.BindJSON(&order); err != nil {
      c.JSON(http.StatusBadRequest, gin.H {
        "success":false,"message": err.Error(),})
      return
    }

    var updateObj []string
    var values []interface {}

    if order.TableId != nil {
    _,err:=models.GetTableById(ctx,strconv.FormatInt(*order.TableId,10))
    if err!=nil{
    c.JSON(http.StatusBadRequest, gin.H {
        "success":false,"message": err.Error(),})
    return
    }
  updateObj = append(updateObj, fmt.Sprintf("table_id=$%d",len(values)+1))
      values = append(values, *order.TableId)
    }

    values = append(values, orderId)

    setVal:= strings.Join(updateObj, ", ")

    err:= models.UpdateOrderDb(ctx, setVal, values)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "success":false,"message": err.Error(),})
      return
    }

    c.JSON(http.StatusOK, gin.H {
      "success": true, "message": "Order updated successfully",
    })
  }
}