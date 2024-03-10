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
	Email         *string
	First_name    *string
	Last_name     *string
	UserType      *string
	jwt.StandardClaims
}
var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(userDetails models.User)(string,string,error){
  claims:=&SignedDetails{
    Userid: userDetails.ID,
    Email: userDetails.Email,
    First_name: userDetails.First_name,
    Last_name: userDetails.Last_name,
    UserType:userDetails.UserType,
    StandardClaims:jwt.StandardClaims{
      ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(5)).Unix(),
    },
  }
  refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
	}

	return accessToken, refreshToken, err
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
