package models

import(
  "log"
  database "github.com/sajagsubedi/Restaurant-Management-Api/database"
)
type Food struct {
  ID *string `json:"id"`
  Name *string `json:"name" validate:"required,min=2,max=100"`
  Price *float64 `json:"price" validate:"required"`
  Food_image *string `json:"food_image" validate:"required"`
}

func GetFoodsDb() ([]Food, error) {
  db:= database.CreateConnection()
  defer db.Close()
  var foods []Food
  sqlStatement:= `SELECT * FROM foods`
  rows,
  err:= db.Query(sqlStatement)
  if err != nil {
    log.Fatalf("Unable to execute sql statement %v", err)
  }
  for rows.Next() {
    var food Food
    err = rows.Scan(&food.ID, &food.Name, &food.Price, &food.Food_image)
    if err != nil {
      log.Fatalf("Unable to scan row %v", err)
    }
    foods = append(foods, food)
  }
  return foods,
  err
}
func CreateFoodDB(createFood Food)(Food, error) {
  db:= database.CreateConnection()
  defer db.Close()
  sqlStatement:= `INSERT INTO foods (id, name, price, food_image)
    VALUES (gen_random_uuid(), $1, $2, $3)
    RETURNING *;`
  var createdFood Food
  err:= db.QueryRow(sqlStatement, createFood.Name, createFood.Price, createFood.Food_image).Scan(&createdFood.ID,&createdFood.Name,&createdFood.Price,&createdFood.Food_image)
  if err != nil {
    log.Fatalf("Unable to execute query %v", err)
  }
  return createdFood,err
}