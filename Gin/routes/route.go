package routes

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

var usuarios []Usuario

func SetupRoutes(r *gin.Engine) {
	r.Static("/static", "./static")

	r.LoadHTMLGlob("template/*")

	r.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/:page", func(ctx *gin.Context) {
		page := ctx.Param("page")

		if !strings.HasSuffix(page, ".html") {
			page += ".html"
		}

		if _, err := os.Stat("template/" + page); err == nil {
			ctx.HTML(http.StatusOK, page, nil)
		} else {
			ctx.HTML(http.StatusNotFound, "404.html", nil)
		}
	})
}
