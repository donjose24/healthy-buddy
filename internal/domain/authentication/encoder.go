package authentication

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jmramos02/healthy-buddy/internal/model"
)

type Claim struct {
	User    model.User
	Details interface{}
	jwt.StandardClaims
}

func encodeUserInfo(user model.User, details interface{}) string {
	appKey := os.Getenv("APPLICATION_KEY")
	expiration := time.Now().Add(1 * time.Hour)

	claim := Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(appKey))

	if err != nil {
		panic("Signing failed")
	}

	return tokenString
}
