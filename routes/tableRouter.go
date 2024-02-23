package routes

import (
  controller "github.com/sajagsubedi/Restaurant-Management-Api/controllers"
	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controller.GetTables())
	incomingRoutes.GET("/tables/:tableid", controller.GetTable())
	incomingRoutes.POST("/tables/add", controller.CreateTable())
	incomingRoutes.PATCH("/tables/update/:tableid", controller.UpdateTable())
}
