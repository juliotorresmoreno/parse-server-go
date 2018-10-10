package auth

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/juliotorresmoreno/parse-server/models"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/labstack/echo"
)

func Login(secret string) func(c echo.Context) error {
	return func(c echo.Context) error {
		data := helper.GetPostParams(c.Request())
		username := data.Get("username")
		password := data.Get("password")
		users := models.Users{}
		user, err := users.FindByUserName(username)
		if err != nil {
			log.Println(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if helper.IsValid(user.Passwd, password) {
			token := jwt.New(jwt.SigningMethodHS256)

			// Set claims
			claims := token.Claims.(jwt.MapClaims)
			claims["id"] = user.ID
			claims["nombres"] = user.Nombres
			claims["apellidos"] = user.Apellidos
			claims["fullname"] = user.FullName
			claims["email"] = user.Email
			claims["usuario"] = user.Usuario
			claims["role"] = user.Role
			claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

			// Generate encoded token and send it as response.
			t, err := token.SignedString([]byte(secret))
			if err != nil {
				log.Println(err)
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return c.JSON(http.StatusOK, map[string]string{
				"token": t,
			})
		}

		return echo.ErrUnauthorized
	}
}

func Restricted(secret string) func(c echo.Context) error {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		id, _ := strconv.Atoi(fmt.Sprintf("%v", claims["id"]))
		_user := map[string]interface{}{
			"id":        uint(id),
			"nombres":   fmt.Sprintf("%v", claims["nombres"]),
			"apellidos": fmt.Sprintf("%v", claims["apellidos"]),
			"usuario":   fmt.Sprintf("%v", claims["usuario"]),
			"fullname":  fmt.Sprintf("%v", claims["fullname"]),
			"email":     fmt.Sprintf("%v", claims["email"]),
			"role":      fmt.Sprintf("%v", claims["role"]),
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
			"result":  _user,
		})
	}
}
