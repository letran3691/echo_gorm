package handle

import (
	"echo_gorm/mdw"
	"echo_gorm/modle"
	"fmt"
	"github.com/badoux/checkmail"
	"github.com/labstack/echo/v4"
	"log"

	"net/http"
)

func Register_Get(c echo.Context)error  {
	return c.Render(http.StatusOK,"Register","")
}


func Register_Post(c echo.Context)error  {
	db := Connect()
	db.LogMode(true)
	Username := c.FormValue("username")
	Password := c.FormValue("password")
	Email 	 := c.FormValue("email")

	log.Printf("Pass %v user %v email %v",Password,Username,Email)

	hashpass,_ := mdw.HashPass(Password)
	role := &modle.Role{
		Role_name: "null",
		Method: "null",
	}
	db.Create(&role)

	Users := &modle.User{
		Username: Username,
		Password: hashpass,
		Email: Email,
		RoleID: role.Id,
	}
	checkemail := checkmail.ValidateFormat(Users.Email)

	if checkemail != nil {
		return c.String(http.StatusBadRequest,fmt.Sprintf("email format invalid!!!"))
	}

	name := 0
	db.Where("username = ?",Username).First(&Users).Count(&name)
	if name > 0{
		log.Println("user existed")
		return c.String(http.StatusBadRequest,fmt.Sprintf("Username existed!!!"))
	}
	email := 0
	db.Where("email = ?",Email).First(&Users).Count(&email)
	if email > 0{
		log.Println("Email existed")
		return c.String(http.StatusBadRequest,fmt.Sprintf("Email existed!!!"))
	}

	db.Create(&Users)


	return c.Redirect(http.StatusMovedPermanently,"/login")
}
