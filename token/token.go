package token

import (
	"log"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/jwt"
)

var secret = "team"

func TokenValidator(tokenString string) error {

	token, err := jws.ParseJWT([]byte(tokenString))
	if err != nil {
		return err
	}

	validator := &jwt.Validator{}
	validator.SetIssuer("app kob")

	err = token.Validate([]byte(secret), crypto.SigningMethodHS256, validator)
	return err
}

func CreateToken(user string) string {
	claims := jws.Claims{}
	claims.SetIssuer("app kob")
	claims.SetAudience(user)

	tokenStruct := jws.NewJWT(claims, crypto.SigningMethodHS256)

	serialized, err := tokenStruct.Serialize([]byte(secret))
	if err != nil {
		log.Fatal("error : ", err.Error())
	}

	token := string(serialized)

	return token
}
