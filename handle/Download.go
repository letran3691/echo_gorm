package handle

import (
	"echo_gorm/modle"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"log"
)

func Download(c echo.Context)error  {
	db := Connect()
	db.LogMode(false)

	id := cast.ToInt64(c.QueryParam("id"))

	log.Printf("id %v",id)


	list_song := &modle.List_Song{
		Id: id,
	}


	//////download

	db.First(list_song)
	//log.Printf("file %v, path %v",list_song.Name_Song,list_song.Path)
	file := list_song.Name_Song
	path := list_song.Path

	dowload := path+file

	return c.Attachment(dowload, list_song.Name_Song)
}