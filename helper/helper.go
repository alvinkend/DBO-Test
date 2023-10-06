package helper

import (
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

var jwtMiddleware *jwtmiddleware.JWTMiddleware
var signingKey []byte

func GetLimitOffset(page, size int) (limit int, offset int) {
	if page == 0 || size == 0 {
		// using -1 to disable gorm size and offset in case page and size not set
		size = -1
		offset = -1
		return size, offset
	}
	offset = (page - 1) * size
	return size, offset
}

func GetAuthenticatedUser(r *http.Request) (int64, error) {
	userID, err := ExtractToken(r, "customer_id")
	if err != nil {
		return 0, err
	}
	return int64(userID.(float64)), nil
}

func ExtractToken(r *http.Request, key string) (interface{}, error) {
	tokenStr, err := jwtMiddleware.Options.Extractor(r)
	if err != nil {
		return "", err
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims[key], nil
	} else {
		return "", nil
	}
}
