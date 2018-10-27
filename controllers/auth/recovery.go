package auth

import (
	"log"
	"net/http"

	"github.com/juliotorresmoreno/parse-server/mail"

	"gopkg.in/mgo.v2/bson"

	"github.com/juliotorresmoreno/parse-server/models"
	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/labstack/echo"
)

type recoveryData struct {
	Token string
}

func Recovery(secret string) func(c echo.Context) error {
	return func(c echo.Context) error {
		token := bson.NewObjectId().Hex()
		data := helper.GetPostParams(c.Request())
		username := data.Get("username")
		users := models.Users{}
		user, err := users.FindByUserName(username)
		if err != nil {
			log.Println(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		user.Recovery = token
		err = user.Save()
		if err != nil {
			log.Println(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		mail := mail.NewMail()
		mail.To = []string{user.Email}
		mail.Template = "recovery"
		mail.Data = recoveryData{Token: token}
		err = mail.SendMail()
		if err != nil {
			log.Println(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"success": true,
		})
	}
}
