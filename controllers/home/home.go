package home

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo"
)

func Register(g *echo.Group) {
	g.GET("/", Home)
}

func getTemplate() (string, error) {
	tpl := bytes.Buffer{}
	t, err := template.ParseGlob("templates/template.html")
	if err != nil {
		return "", err
	}
	if err := t.Execute(&tpl, nil); err != nil {
		return "", err
	}
	return tpl.String(), nil
}

func Home(c echo.Context) error {
	tpl := bytes.Buffer{}
	/*master, err := getTemplate()
	if err != nil {
		return err
	}
	t, err := template.New("content").ParseGlob("templates/home.html")
	if err != nil {
		return err
	}
	render, err := t.Parse(defineDemo)
	if err != nil {
		return err
	}
	if err := render.Execute(&tpl, nil); err != nil {
		return err
	}*/
	var err error

	var defineDemo = `
	{{ define "a" }} Template A {{ end }}
	{{define "b"}} Template B {{end}}
	`
	t := template.New("defineActionDemo")
	t, err = t.Parse(defineDemo)
	if err != nil {
		fmt.Println("parsing failed: %s", err)
	}

	err = t.Execute(&tpl, nil)
	if err != nil {
		fmt.Println("execution failed: %s", err)
	}
	c.HTML(http.StatusOK, tpl.String())
	return nil
}
