package controllers

import(
  "time"
  "context"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/sajagsubedi/Restaurant-Management-Api/models"
)

func GetTables() gin.HandlerFunc {  
  return func(c *gin.Context) {
    ctx,cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    tables,err:= models.GetTablesDb(ctx)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "success":false,"message": "Failed to Fetch tables",
      })
    }
    if tables == nil {
      c.JSON(http.StatusOK, gin.H {
        "tables": [0]models.Table {},
      })
      return
    }
    c.JSON(http.StatusOK, gin.H {
      "success": true,
      "message": "Fetch tables successfully",
      "tables": tables,
    })
  }

}

func GetTable() gin.HandlerFunc {  return func(c *gin.Context) {
    ctx,cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    tableId:= c.Param("tableid")
    table,err:= models.GetTableById(ctx, tableId)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "success":false,"message": err.Error(),
      })
      return
    }
    c.JSON(http.StatusOK, gin.H {
      "success": true,
      "message": "Fetched table successfully",
      "table": table,
    })
  }

}

func CreateTable() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "create table",
    })
  }
}

func UpdateTable() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "update table",
    })
  }
}