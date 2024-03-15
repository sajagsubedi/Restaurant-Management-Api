package models

import (
  "log"
  "fmt"
  "time"
  "context"
  "database/sql"
  database "github.com/sajagsubedi/Restaurant-Management-Api/database"
  )
  type Table struct {
	ID              *int `json:"id"`
	Guests          *int               `json:"guests" validate:"required"`
	TableNumber     *int               `json:"tablenumber" validate:"required"`
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
}

func GetTablesDb(ctx context.Context) ([]Table, error) {
  db:= database.CreateConnection()
  defer db.Close()
  var tables []Table
  sqlStatement:= `SELECT * FROM tables`
  rows,
  err:= db.QueryContext(ctx, sqlStatement)
  if err != nil {
    log.Fatalf("Unable to execute sql statement %v", err)
  }
  for rows.Next() {
    var table Table
    err = rows.Scan(&table.ID, &table.Guests, &table.TableNumber, &table.CreatedAt, &table.UpdatedAt)
    if err != nil {
      log.Fatalf("Unable to scan row %v", err)
    }
    tables = append(tables, table)
  }
  return tables,
  err
}

func GetTableById(ctx context.Context, id string) (Table, error) {
  db:= database.CreateConnection()
  defer db.Close()
  var foundTable Table
  sqlStatement:= `SELECT * FROM tables WHERE id=$1`
  err:= db.QueryRowContext(ctx, sqlStatement, id).Scan(&foundTable.ID, &foundTable.Guests, &foundTable.TableNumber, &foundTable.CreatedAt, &foundTable.UpdatedAt) 
  if err !=nil{
    if err == sql.ErrNoRows {
      return Table {},
      fmt.Errorf("Table with id %s not found", id)
    }

    return Table {},
    fmt.Errorf("error executing query: %w", err)
  }
  return foundTable,
  nil
}

func CreateTableDb(ctx context.Context, createTable Table)(Table, error) {
  db:= database.CreateConnection()
  defer db.Close()
  sqlStatement:= `INSERT INTO tables (guests, tablenumber, created_at, updated_at)
		VALUES ($1, $2, NOW(), NOW())
		RETURNING *;`
  var createdTable Table
  err:= db.QueryRowContext(ctx, sqlStatement, createTable.Guests, createTable.TableNumber).Scan(&createdTable.ID, &createdTable.Guests, &createdTable.TableNumber, &createdTable.CreatedAt, &createdTable.UpdatedAt)
  if err != nil {
    log.Fatalf("Unable to execute query %v", err)
  }
  return createdTable,
  err
}

func UpdateTableDb(ctx context.Context, setVal string, values []interface {}) error {
  db:= database.CreateConnection()
  defer db.Close()
  query:= fmt.Sprintf("UPDATE tables SET %s, updated_at=NOW() WHERE id=$%d", setVal, len(values))
  _,
  err:= db.ExecContext(ctx, query, values...)
  if err != nil { 
    if err == sql.ErrNoRows {
      return fmt.Errorf("Table with id %s not found", values[len(values)-1])
    }

    return fmt.Errorf("Error while executing query: %w", err)
  }

  return err
}
