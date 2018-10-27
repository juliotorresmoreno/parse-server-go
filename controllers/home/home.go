package home

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/juliotorresmoreno/parse-server/render"

	"github.com/labstack/echo"
)

var _renderer render.Renderer

func Register(g *echo.Group, renderer render.Renderer) {
	_renderer = renderer
	g.GET("/", Home)
	g.GET("/sign-in", SignIn)
	g.GET("/sign-up", SignUp)
}

type Data struct {
	Nombre string
}

func Home(c echo.Context) error {
	tpl := bytes.Buffer{}
	t, err := template.ParseGlob("templates/home/home.html")
	if err != nil {
		return err
	}
	if err := t.ExecuteTemplate(&tpl, "content", Data{"Julio"}); err != nil {
		return err
	}
	_renderer.Content = tpl.String()
	c.HTML(http.StatusOK, _renderer.Render())
	return nil
}

func SignUp(c echo.Context) error {
	tpl := bytes.Buffer{}
	t, err := template.ParseGlob("templates/home/register.html")
	if err != nil {
		return err
	}
	if err := t.ExecuteTemplate(&tpl, "content", Data{"Julio"}); err != nil {
		return err
	}
	_renderer.Content = tpl.String()
	c.HTML(http.StatusOK, _renderer.Render())
	return nil
}

func SignIn(c echo.Context) error {
	tpl := bytes.Buffer{}
	t, err := template.ParseGlob("templates/home/login.html")
	if err != nil {
		return err
	}
	if err := t.ExecuteTemplate(&tpl, "content", Data{"Julio"}); err != nil {
		return err
	}
	_renderer.Content = tpl.String()
	c.HTML(http.StatusOK, _renderer.Render())
	return nil
}
