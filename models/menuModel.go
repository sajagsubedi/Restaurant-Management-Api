package models
import(
  "fmt"
  "database/sql"
  "log"
  "time"
  "context"
  database "github.com/sajagsubedi/Restaurant-Management-Api/database"
)
type Menu struct {
	ID        *int64     `json:"id"`
	Name      *string     `json:"name" validate:"required"`
	Category  *string     `json:"category" validate:"required"`
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
func GetMenuById(ctx context.Context, menuid string) (Menu, error) {
    db := database.CreateConnection()
    defer db.Close()
    var foundMenu Menu
    sqlStatement := `SELECT * FROM menus WHERE id=$1`
    err := db.QueryRowContext(ctx, sqlStatement, menuid).Scan(&foundMenu.ID, &foundMenu.Name, &foundMenu.Category, &foundMenu.StartDate, &foundMenu.EndDate, &foundMenu.CreatedAt, &foundMenu.UpdatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return Menu{}, fmt.Errorf("Menu with id %s not found", menuid)
        }
        return Menu{}, fmt.Errorf("error executing query: %w", err)
    }
    return foundMenu, nil
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

func UpdateMenuDb(ctx context.Context, setVal string, values []interface {}) error {
  db:= database.CreateConnection()
  defer db.Close()
  query:= fmt.Sprintf("UPDATE menus SET %s, updated_at=NOW() WHERE id=$%d", setVal, len(values))
  _,
  err:= db.ExecContext(ctx, query, values...)
  if err != nil { 
    if err == sql.ErrNoRows {
      return fmt.Errorf("Menu with id %s not found", values[len(values)-1])
    }

    return fmt.Errorf("Error while executing query: %w", err)
  }

  return err
}
