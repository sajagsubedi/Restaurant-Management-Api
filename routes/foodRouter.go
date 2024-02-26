package routes

import (
  controller "github.com/sajagsubedi/Restaurant-Management-Api/controllers"
  "github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
  incomingRoutes.GET("/api/v1/foods", controller.GetFoods())
  incomingRoutes.GET("/api/v1/foods/:foodid", controller.GetFood())
  incomingRoutes.POST("/api/v1/foods/add", controller.CreateFood())
  incomingRoutes.PATCH("/api/v1/foods/update/:foodid", controller.UpdateFood())
}