package routes

import (
  controller "github.com/sajagsubedi/Restaurant-Management-Api/controllers"
  "github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
  incomingRoutes.GET("/foods", controller.GetFoods())
  incomingRoutes.GET("/foods/:foodid", controller.GetFood())
  incomingRoutes.POST("/foods/add", controller.CreateFood())
  incomingRoutes.PATCH("/foods/update/:foodid", controller.UpdateFood())
}