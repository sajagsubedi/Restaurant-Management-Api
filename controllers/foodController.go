package controllers

import(
  "github.com/gin-gonic/gin"
  "net/http"
  "github.com/sajagsubedi/Restaurant-Management-Api/models"
)

func GetFoods() gin.HandlerFunc {
  return func(c *gin.Context) { 
    foods,err:=models.GetFoodsDb()
    if err!=nil{
    c.JSON(http.StatusInternalServerError,gin.H{"error":"failed to fetch foods"})
    }
    c.JSON(http.StatusOK,foods)
  }
}

func GetFood() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "getFood by id",
    })
  }
}

func CreateFood() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "create food",
    })
  }
}

func UpdateFood() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "Update food",
    })
  }
}