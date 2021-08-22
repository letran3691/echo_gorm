package modle

import (
	"github.com/dgrijalva/jwt-go"
	_ "github.com/qor/validations"
)

type User struct {
	Id int64 `gorm:"primaryKey;autoIncrement:true"`
	Username string
	Password string
	Email string `json:"email"`
	RoleID int64
	jwt.StandardClaims
}


func GetSecretKey() []byte {
	return []byte("sadfaeg4t54yawfarwefef4taaff")
}
