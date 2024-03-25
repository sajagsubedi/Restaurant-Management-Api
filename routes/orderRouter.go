package routes

import (
  "github.com/gin-gonic/gin"
  "github.com/sajagsubedi/Restaurant-Management-Api/middlewares"
  controller "github.com/sajagsubedi/Restaurant-Management-Api/controllers"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
  orderRoutes:=incomingRoutes.Group("/api/v1/orders")
  
  authRoutes:=incomingRoutes.Group("")
  authRoutes.Use(middlewares.CheckUser())
  authRoutes.POST("/add", controller.CreateOrder())
  authRoutes.PATCH("/update/:orderid", controller.UpdateOrder())

	adminRoutes := orderRoutes.Group("")
  adminRoutes.Use(middlewares.CheckAdmin())
  adminRoutes.GET("/", controller.GetOrders())
  adminRoutes.GET("/:orderid", controller.GetOrder())
  }