package router

import (
	"fmt"
	"html/template"
	"os"

	"github.com/gin-gonic/gin"
	"github.otrex/blog/src/routes"
	"github.otrex/blog/src/utils"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Reads templates
	templatePath := os.Getenv("TEMPLATE_PATH")
	files, err := utils.ReadFiles(templatePath)

	if err != nil {
		fmt.Print(err)
		os.Exit(0)
		return r
	}

	// Set max file size to 8 << 20
	r.MaxMultipartMemory = 8 << 20
	// loads up template
	html := template.Must(template.ParseFiles(files...))

	r.SetHTMLTemplate(html)
	// r.LoadHTMLFiles(files...)

	r.Static("/p", os.Getenv("ASSETS_PATH"))

	routes.BlogRoutes(r)

	// NOT FOUND HANDLER
	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	return r
}
