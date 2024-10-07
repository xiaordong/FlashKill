package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type MyClaims struct {
	ID int64 `json:"id"`
	jwt.StandardClaims
}

const aTokenTime = 3 * time.Hour

func GenToken(id int64) (aToken string, err error) {
	claims := MyClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(aTokenTime).Unix(),
			Issuer:    "user Login",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	aToken, err = token.SignedString([]byte("天王打鼹鼠"))
	if err != nil {
		log.Fatalln(err)
		return "", err
	}
	return aToken, nil
}
func ParseToken(aToken string) (claims *MyClaims, err error) {
	var token *jwt.Token
	claims = new(MyClaims)
	token, err = jwt.ParseWithClaims(aToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("天王打鼹鼠"), nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		err = errors.New("invalid token")
		return nil, err
	}
	return claims, nil
}
