package routes

import (
  "github.com/gin-gonic/gin"
  "github.com/sajagsubedi/Restaurant-Management-Api/middlewares"
   controller "github.com/sajagsubedi/Restaurant-Management-Api/controllers"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
  menuRoutes:=incomingRoutes.Group("/api/v1/menus")
  menuRoutes.GET("", controller.GetMenus())
  menuRoutes.GET("/:menuid", controller.GetMenu())
  
	adminRoutes := menuRoutes.Group("")
  adminRoutes.Use(middlewares.CheckAdmin())
  adminRoutes.POST("/add", controller.CreateMenu())
  adminRoutes.PATCH("/update/:menuid", controller.UpdateMenu())
}