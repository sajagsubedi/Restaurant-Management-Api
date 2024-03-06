package helpers
import(
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
	jwt.StandardClaims
}
var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(userDetails models.User)(string,string,error){
  claims:=&SignedDetails{
    Userid: userDetails.ID,
    Email: userDetails.Email,
    First_name: userDetails.First_name,
    Last_name: userDetails.Last_name,
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