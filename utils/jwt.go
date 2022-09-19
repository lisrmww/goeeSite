package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenClaims struct {
	ID uint
	jwt.StandardClaims
}

var jwtSecret = []byte("123456")

//accessToken 过期时间
var accessTokenExpireTime = time.Now().Add(20 * time.Minute).Unix()

//refreshToken 过期时间
var refreshTokenExpireTime = time.Now().Add(2 * time.Hour).Unix()

// GenerateToken 颁发token
///**
// accessTokenString  访问token
// refreshTokenString 刷新token
func GenerateToken(userID uint) (accessTokenString string, refreshTokenString string, err error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessTokenExpireTime,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "wang",
			Subject:   "access_token",
		},
	})
	accessTokenString, err = accessToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: refreshTokenExpireTime,
		IssuedAt:  time.Now().Unix(),
		Issuer:    "wang",
		Subject:   "refresh_token",
	})
	refreshTokenString, err = refreshToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func ParseToken(token string) (*jwt.Token, *TokenClaims, error) {
	Claims := &TokenClaims{}
	tokenClaims, err := jwt.ParseWithClaims(token, Claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	return tokenClaims, Claims, err
}
