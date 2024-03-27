package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sajagsubedi/Restaurant-Management-Api/middlewares"
	controller "github.com/sajagsubedi/Restaurant-Management-Api/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	userRoutes := incomingRoutes.Group("/api/v1/users")

	userRoutes.POST("/signup", controller.Signup())
	userRoutes.POST("/signin", controller.Signin())
	userRoutes.POST("/getaccesstoken",middlewares.ValidateRefreshToken() ,controller.GetAccessToken())
	
	authRoutes:=userRoutes.Group("")
	authRoutes.Use(middlewares.CheckUser())
	authRoutes.GET("/profile", middlewares.CheckUser(), controller.GetUser())
	authRoutes.PATCH("/updateprofile", middlewares.CheckUser(), controller.UpdateProfile())
	adminRoutes := userRoutes.Group("")
	adminRoutes.Use(middlewares.CheckAdmin())
  adminRoutes.GET("", controller.GetUsers())
	adminRoutes.GET("/:userid", middlewares.CheckAdminAndSetUser(), controller.GetUser())
	adminRoutes.PATCH("/update/:userid", middlewares.CheckAdminAndSetUser(), controller.UpdateProfile())
}
