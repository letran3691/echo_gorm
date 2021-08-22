package handle

import (
	"echo_gorm/modle"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	_"github.com/go-sql-driver/mysql"
)



func ListUer(c echo.Context) error  {
	db := Connect()
	db.LogMode(true)

	var User []modle.User
	db.Select("id, username, email").Find(&User)

	log.Println("ID %v",User)

	return c.Render(http.StatusOK,"ListUser",User)


}