package routes

import (
  controller "github.com/sajagsubedi/Restaurant-Management-Api/controllers"
  "github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
  menuRoutes:=incomingRoutes.Group("/api/v1/menus")
  menuRoutes.GET("", controller.GetMenus())
  menuRoutes.GET("/:menuid", controller.GetMenu())
  menuRoutes.POST("/add", controller.CreateMenu())
  menuRoutes.PATCH("/update/:menuid", controller.UpdateMenu())
}