package helpers
import(
  "fmt"
  "log"
  "os"
  "time"
  jwt "github.com/dgrijalva/jwt-go"
  "github.com/sajagsubedi/Restaurant-Management-Api/models"
)
type SignedDetails struct {
	Userid        *int64
	UserType      *string
	TokenType      string
	jwt.StandardClaims
}
var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(userDetails models.User)(string,string,int64,int64,error){
  tokenType:="accesstoken"
  accessTokenExpiration:=time.Now().Local().Add(time.Hour * time.Duration(6)).Unix()
  claims:=&SignedDetails{
    Userid: userDetails.ID,
    UserType:userDetails.UserType,
    TokenType:tokenType,
    StandardClaims:jwt.StandardClaims{
      ExpiresAt:accessTokenExpiration,
    },
  }
  tokenType="refreshtoken"
  refreshTokenExpiration:= time.Now().Local().Add(time.Hour * time.Duration(168)).Unix()
  refreshClaims := &SignedDetails{
     Userid: userDetails.ID,
    UserType:userDetails.UserType,
    TokenType:tokenType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:refreshTokenExpiration,
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
	}

	return accessToken, refreshToken,accessTokenExpiration,refreshTokenExpiration, err
}

func ValidateToken(signedToken string) (*SignedDetails,error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)
	if err !=nil{
	  return nil,err
	}
	claims, ok := token.Claims.(*SignedDetails)
  if !ok || !token.Valid {	
    return claims,fmt.Errorf("Invalid token!")
	}
	
if claims.ExpiresAt < time.Now().Local().Unix() {
	return claims,fmt.Errorf("Token is expired!")
}

	return claims, nil

}
