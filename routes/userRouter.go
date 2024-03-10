package routes 

import(
    "github.com/gin-gonic/gin"
    controller "github.com/sajagsubedi/Restaurant-Management-Api/controllers"
    "github.com/sajagsubedi/Restaurant-Management-Api/middlewares"
)

func UserRoutes(incomingRoutes *gin.Engine) {
  userRoutes:=incomingRoutes.Group("/api/v1/users")
  userRoutes.GET("/", controller.GetUsers())
  userRoutes.GET("/:userid", controller.GetUser())
  userRoutes.POST("/signup", controller.Signup())
  userRoutes.POST("/signin", controller.Login())
  userRoutes.PATCH("/update",controller.UpdateProfile()).Use(middlewares.CheckUser())
}