package main

import (
	"crypto/rand"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func main() {
	mapClaims := jwt.MapClaims{
		"iss": "hxia",
		"sub": "hxia.io",
		"aud": "demo",
		"exp": time.Now().Add(time.Second * 10).UnixMilli(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)

	jwtKey := make([]byte, 32)
	if _, err := rand.Read(jwtKey); err != nil {
		panic(err)
	}

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}

	fmt.Println(tokenString)

	claims, err := ParseJwt(jwtKey, tokenString, jwt.WithExpirationRequired())
	if err != nil {
		panic(err)
	}

	fmt.Println(claims)
}

func ParseJwt(jwtKey any, tokenString string, opts ...jwt.ParserOption) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	}, opts...)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token.Claims, nil
}
