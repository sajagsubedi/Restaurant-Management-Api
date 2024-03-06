package models
import(
  "context"
  "log"
  database "github.com/sajagsubedi/Restaurant-Management-Api/database"
)
type Token struct {
  UserId *int64 `json:"userid"`
  Access_Token string `json:"access_token"`
  Refresh_Token string `json:"refresh_token"`
}

func AddToken(ctx context.Context, UserId *int64, Access_Token string, Refresh_Token string)error {
  db:= database.CreateConnection()
  defer db.Close()
  sqlStatement:=`INSERT INTO tokens (userid, access_token, refresh_token) VALUES ($1, $2, $3) RETURNING *;`
  var token Token
  err:=db.QueryRowContext(ctx,sqlStatement,UserId, Access_Token, Refresh_Token).Scan(&token.UserId,&token.Access_Token,&token.Refresh_Token)
  
  if err!=nil{
   log.Fatalf("Unable to execute query %v", err)
  }
  return err
}