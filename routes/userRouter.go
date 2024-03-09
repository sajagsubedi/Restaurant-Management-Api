package routes 

import(
    "github.com/gin-gonic/gin"
    controller "github.com/sajagsubedi/Restaurant-Management-Api/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
  incomingRoutes.GET("/api/v1/users", controller.GetUsers())
  incomingRoutes.GET("/api/v1/users/:userid", controller.GetUser())
  incomingRoutes.POST("/api/v1/users/signup", controller.Signup())
  incomingRoutes.POST("/api/v1/users/signin", controller.Login())
  incomingRoutes.PATCH("/api/v1/users/update",controller.UpdateProfile())
}