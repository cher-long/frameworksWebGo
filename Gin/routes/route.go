package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Usuario struct {
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

var usuarios []Usuario

func SetupRoutes(r *gin.Engine) {
	r.LoadHTMLGlob("template/*")

	r.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title":   "Mi Aplicación",
			"Heading": "¡Hola Mundo!",
			"Message": "Bienvenido a mi aplicación web con Gin y platilla HTML",
		})
	})

	r.Static("/static", "./static")

}
