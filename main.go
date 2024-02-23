package main

import(
  "os"
  "github.com/gin-gonic/gin"
  "github.com/sajagsubedi/Restaurant-Management-Api/routes"
)

func main() {
  port:= os.Getenv("PORT")
  if port == "" {
    port = "8000"
  }
  router:= gin.New()
  
  //middlewares
  router.use(gin.Logger())

  //routes
  routes.UserRoutes(router)
  routes.FoodRoutes(router)
  routes.MenuRoutes(router)
  routes.TableRoutes(router)
  routes.OrderRoutes(router)
  routes.OrderItemRoutes(router)
  routes.InvoiceRoutes(router)
  
  router.Run(":"+port)
}