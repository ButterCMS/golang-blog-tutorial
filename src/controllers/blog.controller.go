package controllers

import (
	"os"

	ButterCMS "github.com/ButterCMS/buttercms-go"
	"github.com/gin-gonic/gin"
)

type BlogController struct{}

func (x BlogController) Index(ctx *gin.Context) {
	ButterCMS.SetAuthToken(os.Getenv("BUTTERCMS_TOKEN"))

	size := "10"
	page := "1"

	if ctx.Query("page_size") != "" {
		size = ctx.Query("page_size")
	}

	if ctx.Query("page") != "" {
		page = ctx.Query("page")
	}

	params := map[string]string{
		"page":         page,
		"page_size":    size,
		"exclude_body": "false",
	}
	posts, _ := ButterCMS.GetPosts(params)

	ctx.HTML(200, "index.html", gin.H{"Posts": posts})
}

func (x BlogController) GetBlog(ctx *gin.Context) {
	ButterCMS.SetAuthToken(os.Getenv("BUTTERCMS_TOKEN"))

	slug := ctx.Param("slug")
	post, err := ButterCMS.GetPost(slug)

	if err != nil {
		ctx.HTML(200, "blog.html", gin.H{"Error": err})
		return
	}

	ctx.HTML(201, "blog.html", gin.H{"Post": post.Post})
}
