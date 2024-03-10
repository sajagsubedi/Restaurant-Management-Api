package main

import(
  "os"
  "github.com/gin-gonic/gin"
  "github.com/sajagsubedi/Restaurant-Management-Api/routes"
  "github.com/joho/godotenv"
  "github.com/gin-contrib/cors"
  "log"
)

func main() {
  err:= godotenv.Load(".env")
  if err != nil {
    log.Fatal("Error on loading .env file")
  }
  port:= os.Getenv("PORT")
  if port == "" {
    port = "8000"
  }
  router:= gin.New()

  config:= cors.DefaultConfig()
  config.AllowOrigins = []string {
    "https://mypostmaster.netlify.app",
  }
  config.AllowMethods = []string {
    "GET",
    "POST",
    "PUT",
    "DELETE",
    "PATCH",
  }
  config.AllowHeaders = []string {
    "Origin",
    "Content-Type",
    "auth_token",
  }

  //middlewares
  router.Use(cors.New(config))
  router.Use(gin.Logger())

  //routes
  routes.UserRoutes(router)
  routes.FoodRoutes(router)
  routes.MenuRoutes(router)
  routes.OrderRoutes(router)
  routes.OrderItemRoutes(router)
  routes.InvoiceRoutes(router)

  router.Run(":"+port)
}