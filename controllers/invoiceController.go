package controllers

import(
  "fmt"
  "time"
  "context"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/sajagsubedi/Restaurant-Management-Api/models"
)

func GetInvoices() gin.HandlerFunc {  return func(c *gin.Context) {
    ctx,cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    invoices,err:= models.GetInvoicesDb(ctx)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "success":false,"message": "Failed to fetch invoices",
      })
    }
    if invoices == nil {
      c.JSON(http.StatusOK, gin.H {
        "success":true,
        "invoices": [0]models.Invoice {},
      })
      return
    }
    c.JSON(http.StatusOK, gin.H {
      "success": true,
      "message": "Fetch invoices successfully",
      "invoices": invoices,
    })
  }

}

func GetInvoice() gin.HandlerFunc {  return func(c *gin.Context) {
    ctx,cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    invoiceId:= c.Param("invoiceid")
    invoice,err:= models.GetInvoiceById(ctx, invoiceId)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "success":false,"message": err.Error(),
      })
      return
    }
    c.JSON(http.StatusOK, gin.H {
      "success": true,
      "message": "Fetched invoice successfully",
      "invoice": invoice,
    })
  }

}

func CreateInvoice() gin.HandlerFunc {  return func(c *gin.Context) {
    ctx,cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    var invoice models.Invoice
    if err:= c.BindJSON(&invoice); err != nil {
      c.JSON(http.StatusBadRequest, gin.H {
        "success":false,"message": err.Error()})
      return
    }
    validationErr:= validate.Struct(invoice)
    if validationErr != nil {
      c.JSON(http.StatusBadRequest, gin.H {
        "success":false,"message": validationErr.Error()})
      return
    }
    _, err:= models.GetOrderById(ctx, 
      fmt.Sprintf("%d", *invoice.Order_id))
      if err != nil {
        c.JSON(http.StatusBadRequest, gin.H {
          "success": false, "message": err.Error(),
        })
        return
      }
      createdInvoice,err:= models.CreateInvoiceDB(ctx, invoice)
      if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H {
          "success": false, "message": "Failed to add invoice",
        })
        return
      }
      c.JSON(http.StatusCreated, gin.H {
        "success": true,
        "message": "Created invoice successfully!",
        "invoice": createdInvoice,
      })
    }

}

func UpdateInvoice() gin.HandlerFunc {
  return func(c *gin.Context) {}
}