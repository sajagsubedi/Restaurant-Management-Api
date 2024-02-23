package routes

import (    
  controller "github.com/sajagsubedi/Restaurant-Management-Api/controllers"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/api/v1/invoices", controller.GetInvoices())
	incomingRoutes.GET("/api/v1/invoices/:invoiceid", controller.GetInvoice())
	incomingRoutes.POST("/api/v1/invoices/add", controller.CreateInvoice())
	incomingRoutes.PATCH("/api/v1/invoices/update/:invoiceid", controller.UpdateInvoice())
}
