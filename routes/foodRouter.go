package routes

import (
  controller "github.com/sajagsubedi/Restaurant-Management-Api/controllers"
  "github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
  foodRoutes:=incomingRoutes.Group("api/v1/foods")
  foodRoutes.GET("", controller.GetFoods())
  foodRoutes.GET("/:foodid", controller.GetFood())
  foodRoutes.POST("/add", controller.CreateFood())
  foodRoutes.PATCH("/update/:foodid", controller.UpdateFood())
  foodRoutes.DELETE("/delete/:foodid", controller.DeleteFood())
}