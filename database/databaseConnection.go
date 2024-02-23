package database

import(
  "database/sql"
  "log"
  "github.com/joho/godotenv"
  "github.com/"
)
func CreateConnection() *sql.DB {
  err: = godotenv.Load(".env")
  if err != nil {
    log.Fatal("Error on loading .env file")
  }
  //connect to db
  db,err:= sql.Open("postgres", os.Getenv("POSTGRES_URL"))

  if err != nil {
    panic(err)
  }

  //checking connection
  err = d.Ping()

  if err != nil {
    panic(err)
  }
  return db
}

