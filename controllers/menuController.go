package controllers

import(
  "context"
  "time"
  "github.com/gin-gonic/gin"
  "net/http"
  "github.com/sajagsubedi/Restaurant-Management-Api/models"
)

func GetMenus() gin.HandlerFunc {
  return func(c *gin.Context) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() 
    menuId:=c.Param("menuid")
    menuInfo,err:=models.GetMenuById(ctx,menuId)
    if err!=nil{
      c.JSON(http.StatusInternalServerError,gin.H{
        "message":err.Error(),})
        return
    }
    foods,err:=models.GetFoodByMenuId(ctx,menuId)
        if err!=nil{
      c.JSON(http.StatusInternalServerError,gin.H{
        "message":err.Error(),
      })
      return
    }
    if foods==nil{
    c.JSON(http.StatusOK,gin.H{"success":true,"message":"Fetched menu","menuInfo":menuInfo,"foods":[0]models.Food{}})
    }
    c.JSON(http.StatusOK,gin.H{"success":true,"message":"Fetched menu","menuInfo":menuInfo,"foods":foods})
  }
}

func CreateMenu() gin.HandlerFunc {  return func(c *gin.Context) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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
        "error": validationErr.Error()})
      return
    }
    createdMenu,
    err:= models.CreateMenuDB(ctx,menu)
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
    c.JSON(http.StatusOK, gin.H {
      "message": "Update menu",
    })
  }
}