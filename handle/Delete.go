package handle

import (
	"echo_gorm/modle"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Delete(c echo.Context) error  {
	id := cast.ToInt64(c.QueryParam("id"))
	//log.Printf("iddsaaaaaaaaaaaaaa %v",id)

	db := Connect()
	db.LogMode(true)

	list_song := &modle.List_Song{
		Id: id,
	}

	db.Find(list_song)
	log.Printf("query", list_song.Name_Song)

	//remove file
	file := list_song.Name_Song
	path := "./download/"

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files{
		err := os.Remove(path+file.Name())
		log.Println(err)
	}
	db.Delete(list_song)
	return c.String(http.StatusOK, "Delete '"+file+ "'success")
}

