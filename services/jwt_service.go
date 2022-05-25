package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtService struct {
	secret string
	issuer string
}

type Claim struct {
	Sub    int64 `json:"sub"`
	Claims jwt.StandardClaims
}

func NewJWTService() *jwtService {
	return &jwtService{
		secret: os.Getenv("SECRET_KEY"),
		issuer: "website-login",
	}
}

func (s *jwtService) GenerateToken(userid int64) (string, error) {
	claim := &Claim{
		Sub: userid,
		Claims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim.Claims)

	t, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(s.secret), nil
	})

	return err == nil
}
