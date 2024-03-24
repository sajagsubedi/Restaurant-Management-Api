package models

import(
  "fmt"
  "log"
  "time"
  "context"
  "database/sql"
  database "github.com/sajagsubedi/Restaurant-Management-Api/database"
)

type Food struct {
  ID         *int64 `json:"id"`
  Name       *string `json:"name" validate:"required,min=2,max=100"`
  Price      *float64 `json:"price" validate:"required"`
  Food_image *string `json:"food_image"`
  MenuId     *int64 `json:"menuid"`
  CreatedAt   time.Time `json:"created_at"`
  UpdatedAt   time.Time `json:"updated_at"`
}

func GetFoodsDb(ctx context.Context) ([]Food, error) {
  db:= database.CreateConnection()
  defer db.Close()
  var foods []Food
  sqlStatement:= `SELECT * FROM foods`
  rows,
  err:= db.QueryContext(ctx, sqlStatement)
  if err != nil {
    log.Fatalf("Unable to execute sql statement %v", err)
  }
  for rows.Next() {
    var food Food
    err = rows.Scan(&food.ID, &food.Name, &food.Price, &food.Food_image, &food.MenuId, &food.CreatedAt, &food.UpdatedAt)
    if err != nil {
      log.Fatalf("Unable to scan row %v", err)
    }
    foods = append(foods, food)
  }
  return foods,
  err
}

func CreateFoodDB(ctx context.Context, createFood Food)(Food, error) {
  db:= database.CreateConnection()
  defer db.Close()
  sqlStatement:= `INSERT INTO foods (name, price, food_image, menuid, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING *;`
  var createdFood Food
  err:= db.QueryRowContext(ctx, sqlStatement, createFood.Name, createFood.Price, createFood.Food_image, createFood.MenuId).Scan(&createdFood.ID, &createdFood.Name, &createdFood.Price, &createdFood.Food_image, &createdFood.MenuId, &createdFood.CreatedAt, &createdFood.UpdatedAt)
  if err != nil {
    log.Fatalf("Unable to execute query %v", err)
  }
  return createdFood,
  err
}

func GetFoodById(ctx context.Context, id string) (Food, error) {
  db:= database.CreateConnection()
  defer db.Close()
  var foundFood Food
  sqlStatement:= `SELECT * FROM foods WHERE id=$1`
  err:= db.QueryRowContext(ctx, sqlStatement, id).Scan(&foundFood.ID, &foundFood.Name, &foundFood.Price, &foundFood.Food_image, &foundFood.MenuId, &foundFood.CreatedAt, &foundFood.UpdatedAt)
  if err != nil {
    if err == sql.ErrNoRows {
      return Food {},
      fmt.Errorf("Food with id %s not found", id)
    }

    return Food {},
    fmt.Errorf("error executing query: %w", err)
  }
  return foundFood,
  nil
}



func UpdateFoodDb(ctx context.Context, setVal string, values []interface {}) error {
  db:= database.CreateConnection()
  defer db.Close()
  query:= fmt.Sprintf("UPDATE foods SET %s, updated_at=NOW() WHERE id=$%d", setVal, len(values))
  _,
  err:= db.ExecContext(ctx, query, values...)
  if err != nil { 
    if err == sql.ErrNoRows {
      return fmt.Errorf("Food with id %s not found", values[len(values)-1])
    }

    return fmt.Errorf("Error while executing query: %w", err)
  }

  return err
}

func DeleteFoodById(ctx context.Context, id string) error {
  db:= database.CreateConnection()
  defer db.Close()
  sqlStatement:= `DELETE FROM foods WHERE id=$1 RETURNING id;`
  var returnid string
  err:= db.QueryRowContext(ctx, sqlStatement, id).Scan(&returnid)
  if err != nil {
    if err == sql.ErrNoRows {
      return fmt.Errorf("Food with id %s not found", id)
    }

    return fmt.Errorf("Error while executing query: %w", err)
  }
  return err
}

func GetFoodByMenuId(ctx context.Context, menuid string)([]Food, error) {
  db:= database.CreateConnection()
  defer db.Close()
  var foods []Food
  sqlStatement:= `SELECT * FROM foods WHERE menuid=$1`
  rows,
  err:= db.QueryContext(ctx, sqlStatement, menuid)
  if err != nil {
    log.Fatalf("Unable to execute sql statement %v", err)
  }
  defer rows.Close()
  for rows.Next() {
    var food Food
    err = rows.Scan(&food.ID, &food.Name, &food.Price, &food.Food_image, &food.MenuId, &food.CreatedAt, &food.UpdatedAt)
    if err != nil {
      log.Fatalf("Unable to scan row %v", err)
    }
    foods = append(foods, food)
  }
  if err = rows.Err(); err != nil {
    log.Fatalf("Unable to iterate rows %v", err)
  }
  return foods,
  nil
}