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
	ID           *int64           `json:"id"`
	Quantity     *int64            `json:"quantity" validate:"required"`
	Status       *string            `json:"status" validate:"required, eq=Not_Started | eq=Cooking | eq=Completed`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	Food_id      *int64            `json:"food_id" validate:"required"`
	Order_id     *int64             `json:"order_id" validate:"required"`
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
    err = rows.Scan(&orderitem.ID,&orderitem.Quantity,&orderitem.Status,&orderitem.CreatedAt,&orderitem.UpdatedAt,&orderitem.Food_id,&orderitem.Order_id)

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
    err = rows.Scan(&orderitem.ID,&orderitem.Quantity,&orderitem.Status,&orderitem.CreatedAt,&orderitem.UpdatedAt,&orderitem.Food_id,&orderitem.Order_id)
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
  err:= db.QueryRowContext(ctx, sqlStatement, id).Scan(&orderitem.ID,&orderitem.Quantity,&orderitem.Status, &orderitem.CreatedAt,&orderitem.UpdatedAt,&orderitem.Food_id,&orderitem.Order_id)

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

func CreateOrderItemDB(ctx context.Context, newOrderItem OrderItem)(OrderItem, error) {
  db:= database.CreateConnection()
  defer db.Close()
  sqlStatement:= `INSERT INTO orderitems (quantity,status, food_id, order_id, created_at,updated_at)
		VALUES ($1, $2, $3,$4, NOW(), NOW())
		RETURNING *;`
	initialStatus:="Not_Started"
  var createdOrderItem OrderItem
  err:= db.QueryRowContext(ctx, sqlStatement, newOrderItem.Quantity,initialStatus, newOrderItem.Food_id, newOrderItem.Order_id).Scan(&createdOrderItem.ID,&createdOrderItem.Quantity,&createdOrderItem.Status,&createdOrderItem.CreatedAt,&createdOrderItem.UpdatedAt,&createdOrderItem.Food_id,&createdOrderItem.Order_id)
  if err != nil {
    log.Fatalf("Unable to execute query %v", err)
  }
  return createdOrderItem,
  err
}

func UpdateOrderItemDb(ctx context.Context, setVal string, values []interface {}) error {
  db:= database.CreateConnection()
  defer db.Close()
  query:= fmt.Sprintf("UPDATE orderitems SET %s, updated_at=NOW() WHERE id=$%d", setVal, len(values))
  _,
  err:= db.ExecContext(ctx, query, values...)
  if err != nil { 
    if err == sql.ErrNoRows {
      return fmt.Errorf("OrderItem with id %s not found", values[len(values)-1])
    }

    return fmt.Errorf("Error while executing query: %w", err)
  }

  return err
}
