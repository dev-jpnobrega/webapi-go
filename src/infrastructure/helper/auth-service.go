package infrastructure

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/gommon/log"
)

var key = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func SignIn(data interface{}, duration time.Duration) (string, error) {
	claims := &Claims{
		Username: "creds.Username",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(key)
}

// VerifyAndDecode -
func VerifyAndDecode(token string) (jwt.Claims, error) {
	tokenDecode, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	log.Info(err)

	if err != nil {
		return nil, err
	}

	var claims = tokenDecode.Claims

	return claims, err
}
