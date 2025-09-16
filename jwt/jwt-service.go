package jwt

import (
	"errors"
	"fmt"
	"os"
	"github.com/golang-jwt/jwt/v5"
)

var (
  key []byte
)

func GenerateToken(id int) (string, error) {
  key = []byte(os.Getenv("JWT_SECRET"))
  t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
    Subject: fmt.Sprintf("%d", id),
  })
  return t.SignedString(key)
}

func ValidateToken(jws string) (string, error) {
  t, err := jwt.ParseWithClaims(jws, &jwt.RegisteredClaims{}, func(t *jwt.Token) (any, error) {
    return []byte(os.Getenv("JWT_SECRET")), nil
  })
  
  if err != nil {
    return "", err
  }

  claims, ok := t.Claims.(*jwt.RegisteredClaims)
  if !ok || !t.Valid {
    err := errors.New("Invalid token")
    return "", err
  }

  return claims.Subject, nil
}
