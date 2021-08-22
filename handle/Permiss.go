package handle

import (
	"echo_gorm/modle"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"net/http"
	"strings"
)

func Permiss_Get(c echo.Context)error  {
	db := Connect()
	db.LogMode(true)

	id := cast.ToInt64(c.QueryParam("id"))

	role := &modle.Role{
		Id: id,

	}
	db.Find(role).Select("Role_name")
	//a := fmt.Sprint(role)
	fmt.Println(role.Role_name)

	a := []string{"admin","add","edit","read","delete"}


	for _,b := range a{
		s:= strings.Contains(role.Role_name,b)
		d := fmt.Sprint(b ,":",s)
		fmt.Println("opject ",d)
	}


	return c.Render(http.StatusMovedPermanently,"Edit",role)
}

func Permiss_Post(c echo.Context)error  {
	db := Connect()

	db.LogMode(true)

	id := cast.ToInt64(c.QueryParam("id"))
	//id := int64(2)
	//a := c.Request().FormValue("role")
	b ,_ := c.FormParams()
	d:= (b["role"])

	str := strings.Join(d,", ")

	role := modle.Role{
		Id: id,
		Role_name: str,
		Method: "",

	}
	fmt.Println(str)
	//for _,value := range d{
	//	log.Println(value)
	//
	//}

	db.Model(&role).Update(role)

	return c.String(http.StatusOK,"update role success")
}
