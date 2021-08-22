package handle

import (
	"echo_gorm/modle"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"

	//"strings"

)

func Convert_Get(c echo.Context) error  {

	return c.Render(http.StatusMovedPermanently,"Convert", "")

}
func Convert_Post(c echo.Context) error  {
	db := Connect()

	db.LogMode(true)

	// convert video
	uri := c.FormValue("uri")

	// parse uri
	link, err := url.Parse(uri)
	if err != nil {
		log.Fatal(err)
	}
	link_id := link.Query()

	//log.Printf("uri", uri)
	cmdArgs := []string{"-x", "--audio-format", "mp3", "--audio-quality","0" , "--output", "./file/%(title)s.%(ext)s",uri }

	//_, err_c:= exec.Command ("/usr/local/bin/youtube-dl",cmdArgs...).Output()
	//
	//if err_c != nil {
	//	log.Printf("error %v",err_c)
	//}

	// list file
	Path := "./file/"
	Path_Download := "./download/"
	song_id := link_id.Get("v")
	files, err := ioutil.ReadDir(Path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		list_song := &modle.List_Song{
			Name_Song: file.Name(),
			Path: Path_Download,
			Song_id: song_id,
		}
			list := 0
			db.Where("song_id = ?",list_song.Song_id).First(&list_song).Count(&list)
			if list > 0{
				//
				//err := os.Remove(Path+file.Name())
				//if err != nil {
				//	log.Println(err)
				//}
				return c.String(http.StatusBadRequest,fmt.Sprintf("file existed!!!"))

			}else {
				_, err_c:= exec.Command ("/usr/local/bin/youtube-dl",cmdArgs...).Output()

				if err_c != nil {
					log.Printf("error %v",err_c)
				}

				db.Create(&list_song)
				err := os.Rename(Path+file.Name(), Path_Download+file.Name())
				//err := os.Rename(Path+file.Name(), uri+file.Name())
				if err != nil {
					fmt.Println(err)
				}
			}

	}
	return c.String(http.StatusOK,"Download success")
}