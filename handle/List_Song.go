package handle

import (
	"echo_gorm/modle"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ListSong(c echo.Context)error  {
	db := Connect()
	db.LogMode(false)

	var list_song []modle.List_Song
	db.Find(&list_song)

	//log.Println(list_song)
	return c.Render(http.StatusOK,"ListSong", list_song)

}