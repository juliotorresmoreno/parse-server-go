package main

import (
	"fmt"
	"log"

	"github.com/juliotorresmoreno/parse-server/bootstrap"
	"github.com/juliotorresmoreno/parse-server/db"

	"github.com/juliotorresmoreno/parse-server/config"
	"github.com/juliotorresmoreno/parse-server/controllers/auth"
	"github.com/juliotorresmoreno/parse-server/controllers/home"
	"github.com/juliotorresmoreno/parse-server/controllers/static"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config := config.NewConfig()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	db.SetDefaultConf(config.Database.Driver, config.Database.Dsn)

	bootstrap.Inicialize()
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, latency_human=${latency_human}\n",
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:X-XSRF-TOKEN",
	}))
	e.Use(middleware.Gzip())

	static.Register(e.Group(""))
	home.Register(e.Group(""))

	api := e.Group("/api/v1")

	// Restricted group
	auth.Register(api.Group("/auth"))
	api.Use(middleware.JWT([]byte(config.Secret)))

	e.Logger.Fatal(e.Start(fmt.Sprintf("%v:%v", config.Host, config.Port)))
}
