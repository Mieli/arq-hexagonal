package services

import (
	"fmt"
	"time"

	pkgdl "delegacia.com.br/infra/dl"
	"github.com/dgrijalva/jwt-go"
)

type jwtService struct {
	secretKey string
	issure    string
}

func NewJwtService() *jwtService {
	return &jwtService{
		secretKey: pkgdl.GetEnv("SECRET_KEY"),
		issure:    "delegacia",
	}
}

type Claim struct {
	sum string `json:"sum"`
	jwt.StandardClaims
}

func (s *jwtService) GenerateToken(id string) (string, error) {
	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (s *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isvalid := t.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}
		return []byte(s.secretKey), nil
	})
	return err == nil
}
