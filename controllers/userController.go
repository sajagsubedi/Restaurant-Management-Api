package controllers

import(
  "fmt"
  "log"
  "time"
  "strings"
  "context"
  "net/http"
  "github.com/gin-gonic/gin"
  "golang.org/x/crypto/bcrypt"
  "github.com/sajagsubedi/Restaurant-Management-Api/models"
  "github.com/sajagsubedi/Restaurant-Management-Api/helpers"
)

func GetUsers() gin.HandlerFunc {
  return func(c *gin.Context) {
    ctx,cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    users,err:= models.GetUsersDb(ctx)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "success":false,"message": "Failed to fetch users",
      })
    }
    if users == nil {
      c.JSON(http.StatusOK, gin.H {
        "users": [0]models.User {},
      })
      return
    }
    c.JSON(http.StatusOK, gin.H {
      "success": true,
      "message": "Fetch users successfully",
      "users": users,
    })
  }
}


func GetUser() gin.HandlerFunc {  return func(c *gin.Context) {
    ctx,cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    userId,_:= c.Get("userid")
    parameter:="id"
    user,err:= models.FilterUsers(ctx, parameter,userId)
    if user.ID==nil{
      msg:=fmt.Sprintf("User with id %s doesn't exist!",userId)
    c.JSON(http.StatusNotFound, gin.H {
        "success":false,"message":msg,
      })
     return
    }
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
        "success":false,"message": err.Error(),
      })
      return
    }
    c.JSON(http.StatusOK, gin.H {
      "success": true,
      "message": "Fetched profile successfully",
      "user": user,
    })
  }

}

func Signin() gin.HandlerFunc {  
  return func(c *gin.Context) {
    ctx,cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    var user models.User
    var foundUser models.User
    if err:=c.BindJSON(&user);err!=nil{
      c.JSON(http.StatusBadRequest, gin.H {"success":false,"message": err.Error()})
        return
    }
    if user.Email==nil || user.Password==nil{
      c.JSON(http.StatusBadRequest, gin.H {"success":false,"message":"Please enter correct credentials to login!"})
      return 
    }
    parameter:="email"
    foundUser,err:=models.FilterUsers(ctx,parameter,user.Email)
    if foundUser.ID==nil{
      c.JSON(http.StatusBadRequest, gin.H {"success":false,"message":"Please enter correct credentials to login!"})
      return
    }
    if err !=nil{
        c.JSON(http.StatusInternalServerError, gin.H {"success":false,"message":err.Error()})
      return
    }
    isCorrectPassword:=VerifyPassword(*user.Password,*foundUser.Password)
    if isCorrectPassword!=true{  
      c.JSON(http.StatusBadRequest, gin.H {"success":false,"message":"Please enter correct credentials to login!"})
     return
    }
   accessToken, refreshToken, _ :=helpers.GenerateAllTokens(foundUser)
   c.JSON(http.StatusOK,gin.H{"success":true,"message":"Signed in successfully!","access_token":accessToken,"refresh_token":refreshToken})
  }
}

func Signup() gin.HandlerFunc {
  return func(c *gin.Context) { 
    ctx,cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    var user models.User
    if err:=c.BindJSON(&user);err!=nil{
      c.JSON(http.StatusInternalServerError,gin.H{
        "success":false,
        "message":err.Error()})
      return
    }
    validationErr:=validate.Struct(user)
    if validationErr!=nil{
    c.JSON(http.StatusBadRequest, gin.H{"success":false,"message": validationErr.Error()})
    return
    }
    parameter:="email"
    count,err:=models.CountUser(parameter,user.Email)
    if err!=nil{
      c.JSON(http.StatusInternalServerError,gin.H{"success":false,"message":err.Error()})
      return 
    }
    if count >0{
      c.JSON(http.StatusBadRequest,gin.H{"message":"The user with the given email already exists"})
      return 
    }
    parameter="phone"
    count,err=models.CountUser(parameter,user.Phone)
    if err !=nil{
    c.JSON(http.StatusInternalServerError,gin.H{"success":false,"message":err.Error()})
    return 
    }
    if count >0{
      c.JSON(http.StatusBadRequest,gin.H{"message":"The user with the given phone number already exists"})
    return 
  }
  password:=HashPassword(*user.Password)
  user.Password=&password
  insertedUser,err:=models.AddUser(ctx,user)
  accessToken,refreshToken,err:=helpers.GenerateAllTokens(insertedUser)
  if err!=nil{
    c.JSON(http.StatusInternalServerError,gin.H{"message":"Internal Server Error occurred!. Try again later"})
    return
  }

  c.JSON(http.StatusOK,gin.H{"success":true,"message":"New account created Successfully!","access_token":accessToken,"refresh_token":refreshToken})
}}

  func UpdateProfile() gin.HandlerFunc {
    return func(c *gin.Context) {
      ctx,cancel:= context.WithTimeout(context.Background(), 10*time.Second)
      defer cancel()
      userid,_:=c.Get("userid")
      var user models.User
      if err:= c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H {
          "success":false,"message": err.Error()})
        return
      }

      var updateObj []string
      var values []interface {}
     
      if user.First_name != nil {
        updateObj = append(updateObj, fmt.Sprintf("first_name=$%d",len(values)+1))
        values = append(values, *user.First_name)
      }
      
      if user.Last_name != nil {
        updateObj = append(updateObj, fmt.Sprintf("last_name=$%d",len(values)+1))
        values = append(values, *user.Last_name)
      }
           
        if user.Password != nil {
        password:=HashPassword(*user.Password)
        updateObj = append(updateObj, fmt.Sprintf("password=$%d",len(values)+1))
        values = append(values,password)
      }
      values = append(values,userid)
      if len(values)<2{
        c.JSON(http.StatusBadRequest,gin.H{"success":false,"message":"Please enter some fields to update"})
        return
      }
      setVal:= strings.Join(updateObj, ", ")

      err:= models.UpdateUser(ctx, setVal, values)
      if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H {
          "success":false,"message": err.Error()})
        return
      }

      c.JSON(http.StatusOK, gin.H {
        "success": true, "message": "Your profile updated successfully",
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

func VerifyPassword(userPassword string, providedPassword string)bool{
  err:=bcrypt.CompareHashAndPassword([]byte(providedPassword),[]byte(userPassword))
  if err!=nil{
    return false
  }
  return true
}