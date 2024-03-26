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

func CreateOrderDb(ctx context.Context,newOrder Order)(Order, error) {
  db:= database.CreateConnection()
  defer db.Close()
  sqlStatement:= `INSERT INTO orders (order_date, created_at, updated_at,table_id,user_id) 
		VALUES ($1, NOW(), NOW(),$2,$3) 
		RETURNING *;`
  var createdOrder Order
  err:= db.QueryRowContext(ctx,sqlStatement, newOrder.OrderDate, newOrder.TableId, newOrder.UserId).Scan(&createdOrder.ID, &createdOrder.OrderDate, &createdOrder.CreatedAt, &createdOrder.UpdatedAt, &createdOrder.TableId,&createdOrder.UserId)
  if err != nil {
    log.Fatalf("Unable to execute query %v", err)
  }
  return createdOrder,
  err
}

func UpdateOrderDb(ctx context.Context, setVal string, values []interface {}) error {
  db:= database.CreateConnection()
  defer db.Close()
  query:= fmt.Sprintf("UPDATE orders SET %s, updated_at=NOW() WHERE id=$%d AND user_id=$%d", setVal, len(values)-1,len(values))
  _,
  err:= db.ExecContext(ctx, query, values...)
  if err != nil { 
    if err == sql.ErrNoRows {
      return fmt.Errorf("Order with id %s not found", values[len(values)-1])
    }

    return fmt.Errorf("Error while executing query: %w", err)
  }

  return err
}
