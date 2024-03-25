package routes

import (
	"github.com/gin-gonic/gin"
  controller "github.com/sajagsubedi/Restaurant-Management-Api/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
  orderItemRoutes:=incomingRoutes.Group("/api/v1/orderitems")
	orderItemRoutes.GET("/", controller.GetOrderItems())
	orderItemRoutes.GET("/:orderitemid", controller.GetOrderItem())
	orderItemRoutes.POST("/create", controller.CreateOrderItem())
	orderItemRoutes.PATCH("/:orderitemid", controller.UpdateOrderItem())
}
