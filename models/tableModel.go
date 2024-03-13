package models

import (
  "log"
  "time"
  "context"
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
