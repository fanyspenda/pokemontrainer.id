package middlewares

import (
	"time"

	"github.com/spf13/viper"

	"github.com/golang-jwt/jwt"
)

//GenerateTokenJWT generate token
func GenerateTokenJWT(userID uint) (string, error) {

	//payload data
	claims := jwt.MapClaims{}
	claims["userId"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	//header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// generate enkripsi + key
	return token.SignedString(KeyToByte())
}

//KeyToByte to be signed to token
func KeyToByte() []byte {
	key := viper.GetString("jwt.secret")
	return []byte(key)
}
