package controllers

import(
  "log"
  "time"
  "context"
  "net/http"
  "github.com/gin-gonic/gin"
  "golang.org/x/crypto/bcrypt"
  "github.com/sajagsubedi/Restaurant-Management-Api/models"
)

func GetUsers() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "Get users",
    })
  }
}

func GetUser() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "Get user",
    })
  }
}

func Login() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "login",
    })
  }
}

func Signup() gin.HandlerFunc {
  return func(c *gin.Context) { 
    ctx,cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    var user models.User
    if err:=c.BindJSON(&user);err!=nil{
      c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
      return
    }
    validationErr:=validate.Struct(user)
    if validationErr!=nil{
    c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
    return
    }
    parameter:="email"
    count,err:=models.CountUser(parameter,user.Email)
    if err!=nil{
      c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
      return 
    }
    if count >0{
      c.JSON(http.StatusBadRequest,gin.H{"message":"The user with the given email already exists"})
      return 
    }
    parameter="phone"
    count,err=models.CountUser(parameter,user.Phone)
    if err !=nil{
    c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
    return 
    }
    if count >0{
      c.JSON(http.StatusBadRequest,gin.H{"message":"The user with the given phone number already exists"})
    return 
  }
  password:=HashPassword(*user.Password)
  user.Password=&password
  insertedUser,err:=models.AddUser(ctx,user)
  
}}

func UpdateProfile() gin.HandlerFunc {
  return func(c *gin.Context) { 
    c.JSON(http.StatusOK,gin.H{
      "message": "update profile",
    })
  }
}
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}
