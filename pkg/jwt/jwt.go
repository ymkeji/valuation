package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	mySigningKey = []byte("valuation")
)

// EncodeToken 生成
func EncodeToken(id string) (string, error) {
	o := &struct {
		signingMethod jwt.SigningMethod
		claims        jwt.Claims
	}{
		signingMethod: jwt.SigningMethodHS256,
		claims: jwt.RegisteredClaims{
			ID:        id,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 24)),
		},
	}
	token := jwt.NewWithClaims(o.signingMethod, o.claims)
	return token.SignedString(mySigningKey)
}

// DecodeToken 解析token
func DecodeToken(token string) (string, error) {
	tokenInfo, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return "", fmt.Errorf("token is invalid")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return "", fmt.Errorf("JWT token has expired")
			} else {
				return "", fmt.Errorf("fail to parse JWT token")
			}
		}
		return "", fmt.Errorf("unauthorized")
	}

	if claims, ok := tokenInfo.Claims.(jwt.MapClaims); ok && tokenInfo.Valid {
		return claims["jti"].(string), nil
	} else {
		return "", fmt.Errorf("token is invalid")
	}
}
