package routes

import (
  controller "github.com/sajagsubedi/Restaurant-Management-Api/controllers"
  "github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
  incomingRoutes.GET("/api/v1/menus", controller.GetMenus())
  incomingRoutes.GET("/api/v1/menus/:menuid", controller.GetMenu())
  incomingRoutes.POST("/api/v1/menus/add", controller.CreateMenu())
  incomingRoutes.PATCH("/api/v1/menus/update/:menuid", controller.UpdateMenu())
}