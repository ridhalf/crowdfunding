package auth

import (
	"crowdfunding/helper"
	"github.com/dgrijalva/jwt-go"
	"os"
)

type JwtServiceImpl struct {
}

func NewJwtService() JwtService {
	return &JwtServiceImpl{}
}

func (service JwtServiceImpl) GenerateToken(userID int) (string, error) {
	payload := jwt.MapClaims{}
	payload["user_id"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	SecretKey := []byte(os.Getenv("SECRET_KEY"))

	signedToken, err := token.SignedString(SecretKey)
	helper.PanicIfError(err)
	return signedToken, nil
}
