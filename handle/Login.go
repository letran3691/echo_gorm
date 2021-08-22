package handle

import (
	"echo_gorm/mdw"
	"echo_gorm/modle"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
	jwk "github.com/lestrrat-go/jwx/jwk"
	"context"
	"errors"

)

func Login_Get(c echo.Context)error  {
	return c.Render(http.StatusOK,"Login","")
}

func Login_Post(c echo.Context) error  {
	db := Connect()
	db.LogMode(true)
	Username := c.FormValue("username")
	Password := c.FormValue("password")

	//log.Printf("user %v, pass %v",Username, Password)

	user := modle.User{}
	db.First(&user)

	ID_user := user.Id
	fmt.Println("aaaaaaaaaaaaaaaaaa",ID_user)
	//log.Println(user.Password)
	//checkpass := mdw.CheckPass(Password,user.Password)
	//log.Printf("check %v",checkpass)
	if mdw.CheckPass(Password,user.Password){
		c.Set("username",Username)
		c.Set("password",Password)


		//create token
		token := jwt.New(jwt.SigningMethodHS256)

		//set claim
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = Username
		claims["id"] = ID_user
		claims["exp"] = time.Now().Add(4*time.Minute).Unix()

		t,err := token.SignedString(modle.GetSecretKey())
		if err != nil {
			return err
		}
		return c.String(http.StatusOK,t)
		//return c.Redirect(http.StatusMovedPermanently,"/listsong")
	}

	return c.Redirect(http.StatusMovedPermanently,"/login")
}



func GetKey(token *jwt.Token) (interface{}, error) {

	// For a demonstration purpose, Google Sign-in is used.
	// https://developers.google.com/identity/sign-in/web/backend-auth
	//
	// This user-defined KeyFunc verifies tokens issued by Google Sign-In.
	//
	// Note: In this example, it downloads the keyset every time the restricted route is accessed.
	keySet, err := jwk.Fetch(context.Background(), "https://www.googleapis.com/oauth2/v3/certs")
	if err != nil {
		return nil, err
	}

	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("expecting JWT header to have a key ID in the kid field")
	}

	key, found := keySet.LookupKeyID(keyID)

	if !found {
		return nil, fmt.Errorf("unable to find key %q", keyID)
	}

	var pubkey interface{}
	if err := key.Raw(&pubkey); err != nil {
		return nil, fmt.Errorf("Unable to get the public key. Error: %s", err.Error())
	}

	return pubkey, nil
}

func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	Username := claims["username"].(string)
	return c.String(http.StatusOK, "Welcome "+Username+"!")
}