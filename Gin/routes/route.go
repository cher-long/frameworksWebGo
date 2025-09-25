package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {

		c.String(http.StatusOK, "Hoal Mundo!")
	})

	r.GET("/saludo/:nombre", func(c *gin.Context) {
		nombre := c.Param("nombre")
		c.String(http.StatusOK, "Hola, %s!", nombre)
	})

}
