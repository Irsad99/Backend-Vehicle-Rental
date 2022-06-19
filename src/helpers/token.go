package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var myScretKeys = []byte(os.Getenv("JWT_KETS"))

type Claims struct {
	User_Id uint `json:"user_id"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	jwt.StandardClaims
}

func NewToken(id uint, email, role string) *Claims {
	return &Claims{
		User_Id: id,
		Email:   email,
		Role:    role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
	}
}

func (c *Claims) Create() (string, error) {
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return tokens.SignedString(myScretKeys)
}

func CheckToken(token, role string) (bool, error) {
	tokens, err := jwt.ParseWithClaims(token, &Claims{Role: role}, func(t *jwt.Token) (interface{}, error) {
		return []byte(myScretKeys), nil
	})

	if err != nil {
		return false, err
	}

	claims := tokens.Claims.(*Claims)

	if claims.Role == role {
		return tokens.Valid, nil
	} else {
		if claims.Role == "admin" {
			return tokens.Valid, nil
		} else {
			return false, err
		}
	}
}

func EksToken(token string) (*Claims, error) {
	tokens, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(myScretKeys), nil
	})

	if err != nil {
		return nil, err
	}

	claims := tokens.Claims.(*Claims)

	return claims, nil
}
