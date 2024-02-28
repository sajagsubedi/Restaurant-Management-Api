package controllers

import(
  "github.com/gin-gonic/gin"
  "net/http"
  "github.com/sajagsubedi/Restaurant-Management-Api/models"
  "github.com/go-playground/validator/v10"
)
var validate = validator.New()

func GetFoods() gin.HandlerFunc {
  return func(c *gin.Context) {
    foods,
    err:= models.GetFoodsDb()
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "error": "failed to fetch foods",
      })
    }
    if foods == nil {
      c.JSON(http.StatusOK, gin.H {
        "foods": [0]models.Food {},
      })
      return
    }
    c.JSON(http.StatusOK, gin.H {
      "foods": foods,
    })
  }
}

func GetFood() gin.HandlerFunc {
  return func(c *gin.Context) {
    foodId:= c.Param("foodid")
    food,
    err:= models.GetFoodById(foodId)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "error": "Failed to fetch food",
      })

    }
    c.JSON(http.StatusOK, gin.H {
      "success": true,
      "msg": "Fetched Food successfully",
      "food": food,
    })
  }
}

func CreateFood() gin.HandlerFunc {
  return func(c *gin.Context) {
    var food models.Food
    if err:= c.BindJSON(&food); err != nil {
      c.JSON(http.StatusBadRequest, gin.H {
        "error": err.Error()})
      return
    }
    validationErr:= validate.Struct(food)
    if validationErr != nil {
      c.JSON(http.StatusBadRequest, gin.H {
        "error": validationErr.Error()})
      return
    }
    createdFood,
    err:= models.CreateFoodDB(food)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "success": false, "msg": "Failed to add food",
      })
      return
    }
    c.JSON(http.StatusCreated, gin.H {
      "success": true,
      "msg": "Created food successfully!",
      "food": createdFood,
    })
  }
}

func UpdateFood() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H {
      "message": "Update food",
    })
  }
}

func DeleteFood() gin.HandlerFunc {
  return func(c *gin.Context) {
    foodId:= c.Param("foodid")
    err:= models.DeleteFoodById(foodId)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "success": false, "msg": "Failed to delete food",
      })
      return
    }
  c.JSON(http.StatusOK, gin.H {
    "success": true, "msg": "Delete food successfully!",
  })
}
}