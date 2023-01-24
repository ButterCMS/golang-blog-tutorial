package routes

import (
	"github.com/gin-gonic/gin"
	"github.otrex/blog/src/controllers"
)

var BlogController = controllers.BlogController{}

func BlogRoutes(r *gin.Engine) {
	r.GET("/", BlogController.Index)
	r.GET("/blog/:slug", BlogController.GetBlog)
}
