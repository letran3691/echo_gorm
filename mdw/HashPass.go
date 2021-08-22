package mdw

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPass(Password string) (string, error)  {
	hashpass, err := bcrypt.GenerateFromPassword([]byte(Password),15)

	return string(hashpass),err
}

func CheckPass(Password,hashpass string)bool  {
	err := bcrypt.CompareHashAndPassword([]byte(hashpass),[]byte(Password))
	return err ==nil
}