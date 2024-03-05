package models

import (
  "log"
  "fmt"
	"time"
	"context"
  database "github.com/sajagsubedi/Restaurant-Management-Api/database"
)

type User struct {
	ID         *int64     `json:"id"`
	First_name *string    `json:"first_name" validate:"required,min=2,max=100"`
	Last_name  *string    `json:"last_name" validate:"required,min=2,max=100"`
	Password   *string    `json:"password" validate:"required,min=6"`
	Email      *string    `json:"email" validate:"email,required"`
	Phone      *string    `json:"phone" validate:"required"`
	Created_at time.Time  `json:"created_at"`
	Updated_at time.Time  `json:"updated_at"`
}

func AddUser(ctx context.Context,user User)(User,error){
   db:= database.CreateConnection()
  defer db.Close()
  sqlStatement:= `INSERT INTO users (first_name, last_name, password, email, phone,created_at, updated_at)
    VALUES ($1, $2, $3, $4, $5, NOW(), NOW()) 
    RETURNING *;`
    var createdUser User 
    err:=db.QueryRowContext(ctx,sqlStatement,user.First_name,user.Last_name,user.Password,user.Email,user.Phone,user.Created_at,user.Updated_at).Scan(&createdUser.First_name,&createdUser.Last_name,&createdUser.Password,&createdUser.Email,&createdUser.Phone,&createdUser.Created_at,&createdUser.Updated_at)
    if err != nil {
    log.Fatalf("Unable to execute query %v", err)
  }
  return createdUser,err
}

func CountUser(parameter string, value *string)(int,error){
   db:= database.CreateConnection()
   defer db.Close()
   sqlStatement:=fmt.Sprintf("SELECT COUNT(*) FROM users WHERE %s=$1",parameter)
   count:=0
   err:=db.QueryRow(sqlStatement,value).Scan(&count)
   if err!=nil{
    log.Fatalf("Unable to execute query %v", err)
   }
   return count,err
  }