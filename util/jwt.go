package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/lyd2/live/pkg/setting"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(setting.JwtexpireTime)

	claims := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "mylive",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
