package helpers

import (
	"time"

	// "github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt"
	// "github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("@4d61aw23rw")

type Claims struct {
    UserID uint `json:"user_id"`
    jwt.StandardClaims
}

func GenerateToken(userID uint) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func ValidateToken(tokenStr string) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil {
        return nil, err
    }
    if !token.Valid {
        return nil, err
    }
    return claims, nil
}
