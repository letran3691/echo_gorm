package handle

import (
	"echo_gorm/modle"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(
		"mysql",
		"trunglv:123456@tcp(10.10.0.13:3306)/user_permiss?charset=utf8mb4&parseTime=True&loc=Local",
	)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Get(c echo.Context) error {
	db := Connect()
	db.LogMode(true)

	var Id int64
	var user = &modle.User{
		Id: Id,
	}
	db.Find(&user)
	log.Printf("ID %v", user)

	return c.Render(http.StatusOK, "Permiss", user)

}

func Handlemigrate() {
	db := Connect()
	db.AutoMigrate(modle.User{}).AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(modle.Role{}, modle.List_Song{})
}
