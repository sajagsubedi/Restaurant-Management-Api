package routes

import (
  controller "github.com/sajagsubedi/Restaurant-Management-Api/controllers"
  "github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
  orderRoutes:=incomingRoutes.Group("/api/v1/orders")
  orderRoutes.GET("/", controller.GetOrders())
  orderRoutes.GET("/:orderid", controller.GetOrder())
  orderRoutes.POST("/add", controller.CreateOrder())
  orderRoutes.PATCH("/update/:orderid", controller.UpdateOrder())
}