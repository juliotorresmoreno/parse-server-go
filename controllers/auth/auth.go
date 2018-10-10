package auth

import (
	"github.com/juliotorresmoreno/parse-server/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Register(g *echo.Group) {
	config := config.NewConfig()
	g.POST("/login", Login(config.Secret))
	g.POST("/register", SignUp(config.Secret))

	r := g.Group("")
	r.Use(middleware.JWT([]byte(config.Secret)))
	r.GET("/session", Restricted(config.Secret))
}
