package routes

import (
  controller "github.com/sajagsubedi/Restaurant-Management-Api/controllers"
	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoutes *gin.Engine) {
  tableRoutes:=incomingRoutes.Group("/api/v1/tables")
  
	tableRoutes.GET("/", controller.GetTables())
	tableRoutes.GET("/:tableid", controller.GetTable())
	tableRoutes.POST("/add", controller.CreateTable())
	tableRoutes.PATCH("/update/:tableid", controller.UpdateTable())
}
