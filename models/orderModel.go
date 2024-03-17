package models

import(
"fmt"
"log"
"time"
"context"
"database/sql"
database "github.com/sajagsubedi/Restaurant-Management-Api/database"
)

type Order struct {
	ID           *int64             `json:"id"`
	OrderDate    time.Time          `json:"order_date" validate:"required"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	TableId     *int64            `json:"table_id" validate:"required"`
	UserId      *int64            `json:"user_id" validate:"required"`
}

func GetOrdersDb(ctx context.Context) ([]Order, error) {
  db:= database.CreateConnection()
  defer db.Close()
  var orders []Order
  sqlStatement:= `SELECT * FROM orders`
  rows,
  err:= db.QueryContext(ctx, sqlStatement)
  if err != nil {
    log.Fatalf("Unable to execute sql statement %v", err)
  }
  for rows.Next() {
    var order Order
    err = rows.Scan(&order.ID, &order.OrderDate, &order.CreatedAt, &order.UpdatedAt, &order.TableId, &order.UserId)
    if err != nil {
      log.Fatalf("Unable to scan row %v", err)
    }
    orders = append(orders, order)
  }
  return orders,
  err
}

func GetOrderById(ctx context.Context, orderid string) (Order, error) {
    db := database.CreateConnection()
    defer db.Close()
    var foundOrder Order
    sqlStatement := `SELECT * FROM orders WHERE id=$1`
    err := db.QueryRowContext(ctx, sqlStatement, orderid).Scan(&foundOrder.ID, &foundOrder.OrderDate, &foundOrder.CreatedAt, &foundOrder.UpdatedAt, &foundOrder.TableId, &foundOrder.UserId)
    if err != nil {
        if err == sql.ErrNoRows {
            return Order{}, fmt.Errorf("Order with id %s not found", orderid)
        }
        return Order{}, fmt.Errorf("error executing query: %w", err)
    }
    return foundOrder, nil
}
