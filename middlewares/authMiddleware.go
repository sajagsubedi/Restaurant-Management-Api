package middlewares

import (
  "strings"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/sajagsubedi/Restaurant-Management-Api/helpers"
  "github.com/go-playground/validator/v10"
)
var validate = validator.New()

type RefreshToken struct {
  RefreshToken string `json:"refresh_token" validate:"required"`
}

func DecodeJwt(c *gin.Context) (*helpers.SignedDetails, int, string) {
  authToken:= c.Request.Header.Get("auth_token")
  if authToken == "" {
    msg:= "No Authorization header provided"
    return nil,
    http.StatusUnauthorized,
    msg
  }
  if strings.HasPrefix(authToken, "Bearer ") {
    authToken = strings.TrimPrefix(authToken, "Bearer ")
  } else {
    msg:= "Invalid authentication format. Include 'Bearer' prefix"
    return nil,
    http.StatusUnauthorized,
    msg
  }
  claims,
  err:= helpers.ValidateToken(authToken)

  if err != nil {
    msg:= "Please provide correct authorization header!"
    return nil,
    http.StatusUnauthorized,
    msg
  }

  if claims.TokenType != "accesstoken" {
    msg:= "Please use access token to access resources!"
    return nil,
    http.StatusUnauthorized,
    msg
  }
  return claims,
  0,
  ""
}

func CheckUser() gin.HandlerFunc {
  return func(c *gin.Context) {
    claims,
    status,
    msg:= DecodeJwt(c)
    if msg != "" {
      c.JSON(status, gin.H {
        "success": false, "message": msg,
      })
      c.Abort()
      return
    }

    c.Set("userid", claims.Userid)
    c.Set("usertype",claims.UserType)
    c.Next()
  }
}

func CheckAdmin() gin.HandlerFunc {
  return func(c *gin.Context) {
    claims,
    status,
    msg:= DecodeJwt(c)
    if msg != "" {
      c.JSON(status, gin.H {
        "success": false, "message": msg,
      })
      c.Abort()
      return
    }
    if claims.UserType != "admin" {
      c.JSON(http.StatusUnauthorized, gin.H {
        "success": false,
        "message": "You are not authorized!",
      })
      c.Abort()
      return
    }

    c.Next()
  }
}

func CheckAdminAndSetUser() gin.HandlerFunc {
  return func(c *gin.Context) {
    claims,
    status,
    msg:= DecodeJwt(c)
    if msg != "" {
      c.JSON(status, gin.H {
        "success": false, "message": msg,
      })
      c.Abort()
      return
    }

    if claims.UserType != "admin" {
      c.JSON(http.StatusUnauthorized, gin.H {
        "success": false,
        "message": "You are not authorized!",
      })
      c.Abort()
      return
    }

    c.Set("userid", c.Param("userid"))
    c.Next()
  }
}

func ValidateRefreshToken() gin.HandlerFunc{
  return func(c *gin.Context) {
    var tokenBody RefreshToken
    if err:= c.BindJSON(&tokenBody); err != nil {
      c.JSON(http.StatusUnauthorized, gin.H {
        "success": false,
        "message": "Provide refresh token to proceed!",
      })
      c.Abort()
      return
    }

    validationErr:= validate.Struct(tokenBody)
    if validationErr != nil {
      c.JSON(http.StatusBadRequest, gin.H {
        "success": false, "message": validationErr.Error()})
      return
    }
  claims,err:= helpers.ValidateToken(tokenBody.RefreshToken)
    if err!=nil{
       c.JSON(http.StatusUnauthorized, gin.H {
        "success": false, "message": "Please provide valid refresh token!",
      })
      c.Abort()
      return
   
    }
      if claims.TokenType != "refreshtoken" {
      c.JSON(http.StatusUnauthorized, gin.H {
        "success": false, "message": "Please provide refresh token to refetch access token!",
      })
      c.Abort()
      return
    }
    c.Set("userid",claims.Userid)
    c.Set("usertype",claims.UserType)
    c.Next()
  }
}