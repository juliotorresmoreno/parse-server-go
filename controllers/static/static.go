package static

import (
	"github.com/labstack/echo"
)

func Register(g *echo.Group) {
	g.Static("/", "public")
	g.Static("/node_modules", "node_modules")
}
