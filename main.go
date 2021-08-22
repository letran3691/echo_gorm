package main

import (
	"echo_gorm/handle"
	"echo_gorm/modle"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"log"
)

func main() {


	server := echo.New()
	db := handle.Connect()
	handle.Handlemigrate()
	db.LogMode(true)

	temp := &handle.Template{
		Templates: template.Must(template.ParseGlob("./temp/*")),
	}
	server.Renderer= temp
	islogin := middleware.JWT(modle.GetSecretKey())
	log.Printf("%v",islogin)

	server.Use(middleware.Logger())

	server.GET("/",handle.Home)
	//handle.TestHasOne()

	server.GET("/login",handle.Login_Get)
	server.POST("/login",handle.Login_Post)

	r := server.Group("/restricted")
	{
		config := middleware.JWTConfig{
			KeyFunc: handle.GetKey,
		}
		r.Use(middleware.JWTWithConfig(config))
		r.GET("", handle.Restricted)
	}

	server.GET("/register",handle.Register_Get)
	server.POST("/register",handle.Register_Post)

	server.GET("/listsong",handle.ListSong)


	server.GET("/convert",handle.Convert_Get)
	server.POST("/convert",handle.Convert_Post)

	server.GET("/download", handle.Download)
	server.DELETE("/delete", handle.Delete)

	server.GET("/listuser", handle.ListUer)

	server.GET("/permiss", handle.Permiss_Get)
	server.POST("/permiss", handle.Permiss_Post)

	server.Logger.Fatal(server.Start(":8000"))
}

