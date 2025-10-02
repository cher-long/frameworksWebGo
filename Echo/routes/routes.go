package routes

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	template *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.template.ExecuteTemplate(w, name, data)
}

func SetupRoutes(e *echo.Echo) {

	renderer := &TemplateRenderer{
		template: template.Must(template.ParseGlob("template/*.html")),
	}

	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {

		return c.Render(http.StatusOK, "index.html", map[string]string{
			"Title":   "Mi aplicacion Echo",
			"Heading": "Hola Mundo",
			"Message": "Bienvenido a mi aplicacion web con Echo y plantillas HTML",
		})
	})

	e.GET("/saludo", func(c echo.Context) error {
		return c.String(http.StatusOK, "saludos desde Echo")

	})

	e.GET("/saludo/:nombre", func(c echo.Context) error {
		nombre := c.Param("nombre")
		return c.String(http.StatusOK, "¡Hola,"+nombre+"!")

	})

	e.Static("/static", "static")

}
