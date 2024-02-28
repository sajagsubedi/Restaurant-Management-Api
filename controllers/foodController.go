package controllers

import(
  "strings"
  "net/http"
  "github.com/gin-gonic/gin"
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
      "message": "Fetched Food successfully",
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
        "success": false, "message": "Failed to add food",
      })
      return
    }
    c.JSON(http.StatusCreated, gin.H {
      "success": true,
      "message": "Created food successfully!",
      "food": createdFood,
    })
  }
}
func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var food models.Food

		foodId := c.Param("foodid")

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var updateObj []string
		var values []interface{}

		if food.Name != nil {
			updateObj = append(updateObj, "name=$1")
			values = append(values, *food.Name)
		}

		if food.Price != nil {
			updateObj = append(updateObj, "price=$2")
			values = append(values, *food.Price)
		}

		if food.Food_image != nil {
			updateObj = append(updateObj, "food_image=$3")
			values = append(values, *food.Food_image)
		}
		values = append(values, foodId)

		setVal:=strings.Join(updateObj,", ")
		
    err:=models.UpdateFoodDb(setVal,values)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update food"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Food item updated successfully"})
	}
}


func DeleteFood() gin.HandlerFunc {
  return func(c *gin.Context) {
    foodId:= c.Param("foodid")
    err:= models.DeleteFoodById(foodId)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "success": false, "message": "Failed to delete food",
      })
      return
    }
    c.JSON(http.StatusOK, gin.H {
      "success": true, "message": "Delete food successfully!",
    })
  }
}