package modle

import (
	_ "github.com/qor/validations"
)

type Role struct {
	Id int64 `gorm:"primaryKey;autoIncrement:true"`
	Role_name string `json:"role_name"`
	Method string
}