package controllers

import(
  "strings"
  "context"
  "time"
  "github.com/gin-gonic/gin"
  "net/http"
  "github.com/sajagsubedi/Restaurant-Management-Api/models"
)

func GetMenus() gin.HandlerFunc {
  return func(c *gin.Context) {
    ctx,cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    menus,err:= models.GetMenusDb(ctx)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "error": "failed to fetch menus",
      })
    }
    if menus == nil {
      c.JSON(http.StatusOK, gin.H {
        "menus": [0]models.Menu {},
      })
      return
    }
    c.JSON(http.StatusOK, gin.H {
      "menus": menus,
    })
  }
}

func GetMenu() gin.HandlerFunc {
  return func(c *gin.Context) {
    ctx,
    cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    menuId:= c.Param("menuid")
    menuInfo,
    err:= models.GetMenuById(ctx, menuId)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "message": err.Error(),
      })
      return
    }
    foods,
    err:= models.GetFoodByMenuId(ctx, menuId)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "message": err.Error(),
      })
      return
    }
    if foods == nil {
      c.JSON(http.StatusOK, gin.H {
        "success": true, "message": "Fetched menu", "menuInfo": menuInfo, "foods": [0]models.Food {},})
    }
    c.JSON(http.StatusOK, gin.H {
      "success": true, "message": "Fetched menu", "menuInfo": menuInfo, "foods": foods,
    })
  }
}

func CreateMenu() gin.HandlerFunc {
  return func(c *gin.Context) {
    ctx,
    cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    var menu models.Menu
    if err:= c.BindJSON(&menu); err != nil {
      c.JSON(http.StatusBadRequest, gin.H {
        "error": err.Error()})
      return
    }
    validationErr:= validate.Struct(menu)
    if validationErr != nil {
      c.JSON(http.StatusBadRequest, gin.H {
        "error": validationErr.Error(),})
      return
    }
    createdMenu,
    err:= models.CreateMenuDB(ctx, menu)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "success": false, "message": "Failed to add menu",
      })
      return
    }
    c.JSON(http.StatusCreated, gin.H {
      "success": true,
      "message": "Created menu successfully!",
      "menu": createdMenu,
    })
  }

}

func UpdateMenu() gin.HandlerFunc {
  return func(c *gin.Context) {
    ctx,
    cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var menu models.Menu

    menuId:= c.Param("menuid")

    if err:= c.BindJSON(&menu); err != nil {
      c.JSON(http.StatusBadRequest, gin.H {
        "error": err.Error(),})
      return
    }

    var updateObj []string
    var values []interface {}

    if menu.Name != nil {
      updateObj = append(updateObj, "name=$1")
      values = append(values, *menu.Name)
    }

    if menu.Category != nil {
      updateObj = append(updateObj, "category=$2")
      values = append(values, *menu.Category)
    }

    if menu.StartDate != nil {
      updateObj = append(updateObj, "start_date=$3")
      values = append(values, *menu.StartDate)
    }
    if menu.EndDate != nil {
      updateObj = append(updateObj, "end_date=$3")
      values = append(values, *menu.StartDate)
    }
    
    values = append(values, menuId)

    setVal:= strings.Join(updateObj, ", ")

    err:= models.UpdateMenuDb(ctx, setVal, values)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "error": err.Error(),})
      return
    }

    c.JSON(http.StatusOK, gin.H {
      "success": true, "message": "Menu updated successfully",
    })
  }
}