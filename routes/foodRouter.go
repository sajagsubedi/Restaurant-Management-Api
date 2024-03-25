package routes

import (
  "github.com/gin-gonic/gin"
  "github.com/sajagsubedi/Restaurant-Management-Api/middlewares"
  controller "github.com/sajagsubedi/Restaurant-Management-Api/controllers"
  	
)

func FoodRoutes(incomingRoutes *gin.Engine) {
  foodRoutes:=incomingRoutes.Group("api/v1/foods")
  foodRoutes.GET("", controller.GetFoods())
  foodRoutes.GET("/:foodid", controller.GetFood())
  
	adminRoutes := foodRoutes.Group("")
  adminRoutes.Use(middlewares.CheckAdmin())
  adminRoutes.POST("/add", controller.CreateFood())
  adminRoutes.PATCH("/update/:foodid", controller.UpdateFood())
  adminRoutes.DELETE("/delete/:foodid", controller.DeleteFood())
}