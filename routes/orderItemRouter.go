package routes

import (
  controller "github.com/sajagsubedi/Restaurant-Management-Api/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/api/v1/orderitems", controller.GetOrderItems())
	incomingRoutes.GET("/api/v1/orderitems/:orderitemId", controller.GetOrderItem())
	incomingRoutes.GET("/api/v1/orderitems/order/:orderid", controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/api/v1/orderitems/create", controller.CreateOrderItem())
	incomingRoutes.PATCH("/api/v1/orderitems/:orderitemId", controller.UpdateOrderItem())
}
