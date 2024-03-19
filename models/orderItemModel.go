package models

import (
  "fmt"
  "log"
	"time"
	"context"
	"database/sql"
	database "github.com/sajagsubedi/Restaurant-Management-Api/database"
)

type OrderItem struct {
	ID            *int64 `json:"id"`
	Quantity      *int64            `json:"quantity" validate:"required"`
	Unit_price    *float64           `json:"unit_price" validate:"required"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	Food_id       *int64            `json:"food_id" validate:"required"`
	Order_id      *int64             `json:"order_id" validate:"required"`
}

func GetOrderItemsDb(ctx context.Context) ([]OrderItem, error) {
  db:= database.CreateConnection()
  defer db.Close()
  var orderitems []OrderItem
  sqlStatement:= `SELECT * FROM orderitems`
  rows,
  err:= db.QueryContext(ctx, sqlStatement)
  if err != nil {
    log.Fatalf("Unable to execute sql statement %v", err)
  }
  for rows.Next() {
    var orderitem OrderItem   
    err = rows.Scan(&orderitem.ID,&orderitem.Quantity,&orderitem.Unit_price,&orderitem.CreatedAt,&orderitem.UpdatedAt,&orderitem.Food_id,&orderitem.Order_id)

    if err != nil {
      log.Fatalf("Unable to scan row %v", err)
    }
    orderitems = append(orderitems, orderitem)
  }
  return orderitems,
  err
}

func GetOrderItemsByOrderId(ctx context.Context,orderid string) ([]OrderItem, error) {
  db:= database.CreateConnection()
  defer db.Close()
  var orderitems []OrderItem
  sqlStatement:= `SELECT * FROM orderitems WHERE order_id=$1`
  rows,
  err:= db.QueryContext(ctx, sqlStatement,orderid)
  if err != nil {
    log.Fatalf("Unable to execute sql statement %v", err)
  }
  for rows.Next() {
    var orderitem OrderItem
    err = rows.Scan(&orderitem.ID,&orderitem.Quantity,&orderitem.Unit_price,&orderitem.CreatedAt,&orderitem.UpdatedAt,&orderitem.Food_id,&orderitem.Order_id)
    if err != nil {
      log.Fatalf("Unable to scan row %v", err)
    }
    orderitems = append(orderitems, orderitem)
  }
  return orderitems,
  err
}
func GetOrderItemById(ctx context.Context, id string) (OrderItem, error) {
  db:= database.CreateConnection()
  defer db.Close()
  var orderitem OrderItem
  sqlStatement:= `SELECT * FROM orderitems WHERE id=$1`
  err:= db.QueryRowContext(ctx, sqlStatement, id).Scan(&orderitem.ID,&orderitem.Quantity,&orderitem.Unit_price,&orderitem.CreatedAt,&orderitem.UpdatedAt,&orderitem.Food_id,&orderitem.Order_id)

  if err != nil {
    if err == sql.ErrNoRows {
      return OrderItem {},
      fmt.Errorf("Order Item with id %s not found", id)
    }

    return OrderItem {},
    fmt.Errorf("error executing the query: %w", err)
  }
  return orderitem,
  nil
}
