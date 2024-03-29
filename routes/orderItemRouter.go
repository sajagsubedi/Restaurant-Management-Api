package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sajagsubedi/Restaurant-Management-Api/middlewares"
  controller "github.com/sajagsubedi/Restaurant-Management-Api/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
  orderItemRoutes:=incomingRoutes.Group("/api/v1/orderitems")
  authRoutes:=orderItemRoutes.Group("")
	authRoutes.Use(middlewares.CheckUser())
  
  authRoutes.POST("/create", controller.CreateOrderItem())
	authRoutes.PATCH("/:orderitemid", controller.UpdateOrderItem())
	
  adminRoutes:=orderItemRoutes.Group("")
	adminRoutes.Use(middlewares.CheckAdmin())
	
	adminRoutes.GET("/", controller.GetOrderItems())
	adminRoutes.GET("/:orderitemid", controller.GetOrderItem())
}
