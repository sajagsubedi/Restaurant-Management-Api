package controllers

import(
  "fmt"
  "time"
  "strings"
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

func CreateTable() gin.HandlerFunc {  return func(c *gin.Context) {
    ctx,cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    var table models.Table
    if err:= c.BindJSON(&table); err != nil {
      c.JSON(http.StatusBadRequest, gin.H {
        "success":false,"message": err.Error()})
      return
    }
    validationErr:= validate.Struct(table)
    if validationErr != nil {
      c.JSON(http.StatusBadRequest, gin.H {
        "success":false,"message": validationErr.Error()})
      return
    }
   createdTable,err:= models.CreateTableDb(ctx, table)
      if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H {
          "success": false, "message": "Failed to add table",
        })
        return
      }
      c.JSON(http.StatusCreated, gin.H {
        "success": true,
        "message": "Created table successfully!",
        "table": createdTable,
      })
    }

}

func UpdateTable() gin.HandlerFunc {    return func(c *gin.Context) {
      ctx,cancel:= context.WithTimeout(context.Background(), 10*time.Second)
      defer cancel()

      var table models.Table

      tableId:= c.Param("tableid")

      if err:= c.BindJSON(&table); err != nil {
        c.JSON(http.StatusBadRequest, gin.H {
          "success":false,"message": err.Error()})
        return
      }

      var updateObj []string
      var values []interface {}

      if table.Guests != nil {        updateObj = append(updateObj, fmt.Sprintf("guests=$%d",len(values)+1))
        values = append(values, *table.Guests)
      }

      if table.TableNumber != nil {
        updateObj = append(updateObj, fmt.Sprintf("tablenumber=$%d",len(values)+1))
        values = append(values, *table.TableNumber)
      }

      values = append(values, tableId)

      setVal:= strings.Join(updateObj, ", ")

      err:= models.UpdateTableDb(ctx, setVal, values)
      if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H {
          "success":false,"message": err.Error()})
        return
      }

      c.JSON(http.StatusOK, gin.H {
        "success": true, "message": "Table  updated successfully",
      })
    }

}