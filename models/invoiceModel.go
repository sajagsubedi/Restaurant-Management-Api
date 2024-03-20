package models

import (
  "log"
	"time"
  "context"
  database "github.com/sajagsubedi/Restaurant-Management-Api/database"
)

type Invoice struct {
	ID               *int64 `json:"id"`
	Order_id         *int64             `json:"order_id"`
	Payment_method   *string            `json:"payment_method" validate:"eq=CARD|eq=CASH"`
	Payment_status   *string            `json:"payment_status" validate:"required,eq=PENDING|eq=PAID"`
	Payment_due_date time.Time          `json:"Payment_due_date"`
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
}

func GetInvoicesDb(ctx context.Context) ([]Invoice, error) {
  db:= database.CreateConnection()
  defer db.Close()
  var invoices []Invoice
  sqlStatement:= `SELECT * FROM invoices`
  rows,
  err:= db.QueryContext(ctx, sqlStatement)
  if err != nil {
    log.Fatalf("Unable to execute sql statement %v", err)
  }
  for rows.Next() {
    var invoice Invoice
    err = rows.Scan(&invoice.ID,&invoice.Order_id,&invoice.Payment_method,&invoice.Payment_status,&invoice.CreatedAt,&invoice.UpdatedAt)
    if err != nil {
      log.Fatalf("Unable to scan row %v", err)
    }
    invoices = append(invoices, invoice)
  }
  return invoices,
  err
}
