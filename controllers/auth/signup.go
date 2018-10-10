package auth

import (
	"log"
	"net/http"

	"github.com/juliotorresmoreno/parse-server/models"
	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/labstack/echo"
)

func SignUp(secret string) func(c echo.Context) error {
	return func(c echo.Context) error {
		data := helper.GetPostParams(c.Request())
		user := models.User{
			Nombres:   data.Get("nombres"),
			Apellidos: data.Get("apellidos"),
			FullName:  data.Get("nombres") + " " + data.Get("apellidos"),
			Email:     data.Get("email"),
			Usuario:   data.Get("username"),
			Passwd:    data.Get("password"),
			Role:      "user",
		}
		err := user.Save()
		if err != nil {
			log.Println(err)
			return echo.NewHTTPError(http.StatusNotAcceptable, err.Error())
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"success": true,
			"result":  user,
		})
	}
}
