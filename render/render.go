package render

import (
	"bytes"
	"text/template"
)

type Renderer struct {
	template  *template.Template
	Bussisnes string
	Content   string
	Contact   string
	Team      string
	Menu      string
	Footer    string
}

func (t Renderer) render(template string, data interface{}) string {
	buff := bytes.Buffer{}
	t.template.ExecuteTemplate(&buff, template, data)
	return buff.String()
}

func (t Renderer) Render() string {
	var err error
	t.template, err = template.ParseGlob("templates/*.html")
	if err != nil {
		return ""
	}
	if t.Menu == "" {
		t.Menu = t.render("menu", t)
	}
	if t.Footer == "" {
		t.Footer = t.render("footer", t)
	}
	if t.Contact == "" {
		t.Contact = t.render("contact", t)
	}
	if t.Bussisnes == "" {
		t.Bussisnes = t.render("bussisnes", t)
	}
	if t.Team == "" {
		t.Team = t.render("team", t)
	}
	return t.render("html", t)
}

func NewRenderer() Renderer {
	return Renderer{}
}
