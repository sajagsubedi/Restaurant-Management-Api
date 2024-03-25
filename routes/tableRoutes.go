package routes

import (
	"github.com/gin-gonic/gin"
  "github.com/sajagsubedi/Restaurant-Management-Api/middlewares"
  controller "github.com/sajagsubedi/Restaurant-Management-Api/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine) {
  tableRoutes:=incomingRoutes.Group("/api/v1/tables")
  
	authRoutes:=tableRoutes.Group("")
	authRoutes.Use(middlewares.CheckAdmin())
	authRoutes.POST("/add", controller.CreateTable())
	authRoutes.PATCH("/update/:tableid", controller.UpdateTable())
	
	
	adminRoutes:=tableRoutes.Group("")
	adminRoutes.Use(middlewares.CheckAdmin())
	adminRoutes.GET("/", controller.GetTables())
	adminRoutes.GET("/:tableid", controller.GetTable())
}
