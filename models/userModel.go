package models
import(
// database "github.com/sajagsubedi/Restaurant-Management-Api/database"
  )
type User struct{
  First_Name *string
  Last_Name *string
  Age *int8
  Phone *string 
  Email *string
  Password *string
  Profile *string
}