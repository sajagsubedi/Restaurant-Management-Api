package routes

import (
  controller "github.com/sajagsubedi/Restaurant-Management-Api/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
  orderItemRoutes:=incomingRoutes.Group("")
	orderItemRoutes.GET("/", controller.GetOrderItems())
	orderItemRoutes.GET("/:orderitemId", controller.GetOrderItem())
	orderItemRoutes.GET("/order/:orderid", controller.GetOrderItemsByOrder())
	orderItemRoutes.POST("/create", controller.CreateOrderItem())
	orderItemRoutes.PATCH("/:orderitemId", controller.UpdateOrderItem())
}
