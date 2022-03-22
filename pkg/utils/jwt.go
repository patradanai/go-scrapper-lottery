package utils

import (
	"fmt"
	"lottery-web-scrapping/configs"
	"time"

	"github.com/golang-jwt/jwt"
)

type IJWTService interface {
	GenerateToken(userId string) (string, error)
	ValidateToken(token string) (*Claims, error)
}

type Claims struct {
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issue     string
}

func NewJWTService() IJWTService {
	return &jwtService{
		secretKey: configs.LoadEnv("JWT_SECRET_KEY"),
		issue:     configs.LoadEnv("JWT_ISSUE"),
	}
}

func (s *jwtService) GenerateToken(userId string) (string, error) {
	expiredTime := time.Now().Add(time.Minute * 15).Unix()

	claims :=
		Claims{
			jwt.StandardClaims{Id: userId, Issuer: s.issue, ExpiresAt: expiredTime},
		}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(s.secretKey))

	return token, err
}

func (s *jwtService) ValidateToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err

}
