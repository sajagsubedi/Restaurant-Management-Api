package models
import(
  "log"
  "time"
  "context"
  database "github.com/sajagsubedi/Restaurant-Management-Api/database"
)
type Menu struct {
	ID        *int64     `json:"id"`
	Name      string     `json:"name" validate:"required"`
	Category  string     `json:"category" validate:"required"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func GetMenusDb(ctx context.Context) ([]Menu, error) {
  db:= database.CreateConnection()
  defer db.Close()
  var menus []Menu
  sqlStatement:= `SELECT * FROM menus`
  rows,
  err:= db.QueryContext(ctx,sqlStatement)
  if err != nil {
    log.Fatalf("Unable to execute sql statement %v", err)
  }
  for rows.Next() {
    var menu Menu
    err = rows.Scan(&menu.ID, &menu.Name, &menu.Category, &menu.StartDate, &menu.EndDate,&menu.CreatedAt,&menu.UpdatedAt)
    if err != nil {
      log.Fatalf("Unable to scan row %v", err)
    }
    menus = append(menus, menu)
  }
  return menus,
  err
}
func CreateMenuDB(ctx context.Context,createMenu Menu)(Menu, error) {
  db:= database.CreateConnection()
  defer db.Close()
  sqlStatement:= `INSERT INTO menus (name, category, start_date, end_date, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, NOW(), NOW()) 
		RETURNING *;`
  var createdMenu Menu
  err:= db.QueryRowContext(ctx,sqlStatement, createMenu.Name, createMenu.Category, createMenu.StartDate, createMenu.EndDate).Scan(&createdMenu.ID, &createdMenu.Name, &createdMenu.Category, &createdMenu.StartDate, &createdMenu.EndDate,&createdMenu.CreatedAt,&createdMenu.UpdatedAt)
  if err != nil {
    log.Fatalf("Unable to execute query %v", err)
  }
  return createdMenu,
  err
}
