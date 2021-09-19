package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(name string, user bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtClaim struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "pinku",
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if len(secret) == 0 {
		secret = "diptochuck"
	}
	return secret
}

func (jwtsvc *jwtService) GenerateToken(username string, user bool) string {
	claims := &jwtClaim{
		username,
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    jwtsvc.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenSign, err := token.SignedString([]byte(jwtsvc.secretKey))
	if err != nil {
		panic(err)
	}
	return tokenSign
}

func (jwtsvc *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("incorrect SignIn Method: %v", t.Header["alg"])
		}
		return []byte(jwtsvc.secretKey), nil
	})
}
