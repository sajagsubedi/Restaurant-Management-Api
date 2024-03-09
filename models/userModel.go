package models

import (
  "log"
  "fmt"
	"time"
	"context"
	"database/sql"
  database "github.com/sajagsubedi/Restaurant-Management-Api/database"
)

type User struct {
	ID         *int64     `json:"id"`
	First_name *string    `json:"first_name" validate:"required,min=2,max=100"`
	Last_name  *string    `json:"last_name" validate:"required,min=2,max=100"`
	Password   *string    `json:"password" validate:"required,min=6"`
	Email      *string    `json:"email" validate:"email,required"`
	Phone      *string    `json:"phone" validate:"required, eq=admin|eq=user`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
func GetUsersDb(ctx context.Context) ([]User, error) {
  db:= database.CreateConnection()
  defer db.Close()
  var users []User
  sqlStatement:= `SELECT * FROM users`
  rows,
  err:= db.QueryContext(ctx, sqlStatement)
  if err != nil {
    log.Fatalf("Unable to execute sql statement %v", err)
  }
  for rows.Next() {
    var user User
    err = rows.Scan(&user.ID,&user.First_name,&user.Last_name,&user.Password,&user.Email,&user.Phone,&user.CreatedAt,&user.UpdatedAt)
    if err != nil {
      log.Fatalf("Unable to scan row %v", err)
    }
    users = append(users, user)
  }
  return users,
  err
}

func AddUser(ctx context.Context,user User)(User,error){
   db:= database.CreateConnection()
  defer db.Close()
  sqlStatement:= `INSERT INTO users (first_name, last_name, password, email, phone,created_at, updated_at)
    VALUES ($1, $2, $3, $4, $5, NOW(), NOW()) 
    RETURNING *;`
    var createdUser User 
    err:=db.QueryRowContext(ctx, sqlStatement, user.First_name, user.Last_name, user.Password, user.Email, user.Phone).Scan(&createdUser.ID,&createdUser.First_name,&createdUser.Last_name,&createdUser.Password,&createdUser.Email,&createdUser.Phone,&createdUser.CreatedAt,&createdUser.UpdatedAt)
    if err != nil {
    log.Fatalf("Unable to execute query %v", err)
  }
  return createdUser,err
}

func CountUser(parameter string, value *string)(int,error){
   db:= database.CreateConnection()
   defer db.Close()
   sqlStatement:=fmt.Sprintf("SELECT COUNT(*) FROM users WHERE %s=$1;",parameter)
   count:=0
   err:=db.QueryRow(sqlStatement,value).Scan(&count)
   if err!=nil{
    log.Fatalf("Unable to execute query %v", err)
   }
   return count,err
  }
  
  func FilterUsers(ctx context.Context,parameter string ,value interface{})(User, error){
    db:= database.CreateConnection()
    defer db.Close()
    var foundUser User
   sqlStatement:=fmt.Sprintf("SELECT * FROM users WHERE %s=$1;",parameter)
   err:=db.QueryRowContext(ctx,sqlStatement,value).Scan(&foundUser.ID, &foundUser.First_name, &foundUser.Last_name, &foundUser.Password, &foundUser.Email, &foundUser.Phone, &foundUser.CreatedAt,&foundUser.UpdatedAt)
   if err!=nil{
    if err == sql.ErrNoRows {
      return User{},nil 
    }
    return User{},fmt.Errorf("Internal Server Error")
   }
   return foundUser,err
  
  }
  