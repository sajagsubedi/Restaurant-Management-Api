package routes

import (    
  controller "github.com/sajagsubedi/Restaurant-Management-Api/controllers"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
  invoiceRoutes:=incomingRoutes.Group("/api/v1/invoices")
	invoiceRoutes.GET("/", controller.GetInvoices())
	invoiceRoutes.GET("/:invoiceid", controller.GetInvoice())
	invoiceRoutes.POST("/add", controller.CreateInvoice())
	invoiceRoutes.PATCH("/update/:invoiceid", controller.UpdateInvoice())
}
