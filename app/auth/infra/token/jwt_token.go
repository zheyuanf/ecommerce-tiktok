package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	TokenAuthenticator *JWTTokenAuthenticator
)

func InitAuthenticator(signKey string) {
	TokenAuthenticator = NewJWTTokenAuthenticator(
		signKey,
	)
}

type MyClaims struct {
	UserId int32 `json:"user_id"`
	jwt.RegisteredClaims
}

// JWTTokenAuthenticator JWT token 鉴权器
type JWTTokenAuthenticator struct {
	signkey    string
	signMethod jwt.SigningMethod
}

func NewJWTTokenAuthenticator(signkey string) *JWTTokenAuthenticator {
	return &JWTTokenAuthenticator{
		signkey:    signkey,
		signMethod: jwt.SigningMethodHS256,
	}
}

func (j *JWTTokenAuthenticator) GenerateToken(userId int32, expireAt time.Time) (string, error) {
	token := jwt.NewWithClaims(j.signMethod, MyClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireAt),
		},
	})
	ss, err := token.SignedString([]byte(j.signkey))
	return ss, err
}

func (j *JWTTokenAuthenticator) ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.signkey), nil
	})
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
