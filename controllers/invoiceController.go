package controllers

import(
  "github.com/gin-gonic/gin"
  "net/http"

)

func GetInvoices() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "Get invoices",
    })
  }
}

func GetInvoice() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "Get invoice by id",
    })
  }
}

func CreateInvoice() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "Create invoice",
    })
  }
}

func UpdateInvoice() gin.HandlerFunc {
  return func(c *gin.Context) {}
}