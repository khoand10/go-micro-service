package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"go-micro-service/services/user_management/domain/entity"
	"time"
)

type (
	Claims struct {
		UserId uint64
		Email  string
		Name   string
	}

	TokenInfo struct {
		Token        string    `json:"token"`
		RefreshToken string    `json:"refresh_token"`
		ExpireAt     time.Time `json:"expire_at"`
	}
)

func (c Claims) Valid() error {
	return nil
}

func CreateJWT(user *entity.User, secretKey string, jwtExpirationHour int) (string, error) {
	claims := jwt.MapClaims{}
	expireAt := time.Now().Add(time.Hour * time.Duration(jwtExpirationHour))
	claims["userId"] = user.ID
	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["exp"] = expireAt.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(token string, secretKey string) (*Claims, error) {
	claims := &Claims{}

	jwtKey := []byte(secretKey)

	fn := func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	}
	tk, err := jwt.ParseWithClaims(token, claims, fn)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("ErrSignatureInvalid: ", err)
			return nil, err
		}
		fmt.Println("Err: ", err)
		return nil, err
	}
	if !tk.Valid {
		fmt.Println("tk.InValid: ", err)
		return nil, err
	}
	return claims, nil
}
