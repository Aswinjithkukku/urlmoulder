package routes

import (
	"github.com/aswinjithkukku/url-moulder/controllers"
	"github.com/gin-gonic/gin"
)

func UrlsRoutes(router *gin.Engine) {
	urls := router.Group("/api/url")
	router.GET("/:slug", controllers.GiveRedirectionOutput)
	{
		urls.GET("/all", controllers.GetAllUrls)
		urls.POST("/add", controllers.AddUrl)
	}
}
