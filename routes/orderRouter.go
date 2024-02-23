package routes

import (
  controller "github.com/sajagsubedi/Restaurant-Management-Api/controllers"
  "github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
  incomingRoutes.GET("/api/v1/orders", controller.GetOrders())
  incomingRoutes.GET("/api/v1/orders/:orderid", controller.GetOrder())
  incomingRoutes.POST("/api/v1/orders/add", controller.CreateOrder())
  incomingRoutes.PATCH("/api/v1/orders/update/:orderid", controller.UpdateOrder())
}