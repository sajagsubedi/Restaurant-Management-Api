package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sajagsubedi/Restaurant-Management-Api/middlewares"
  controller "github.com/sajagsubedi/Restaurant-Management-Api/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
  orderItemRoutes:=incomingRoutes.Group("/api/v1/orderitems")
  
  adminRoutes:=orderItemRoutes.Group("")
	adminRoutes.Use(middlewares.CheckAdmin())
	
	adminRoutes.GET("/", controller.GetOrderItems())
	adminRoutes.GET("/:orderitemid", controller.GetOrderItem())
	
	orderItemRoutes.POST("/create", controller.CreateOrderItem())
	orderItemRoutes.PATCH("/:orderitemid", controller.UpdateOrderItem())
}
