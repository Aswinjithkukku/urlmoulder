package main

import (
	"fmt"
	"os"

	"github.com/aswinjithkukku/url-moulder/controllers"
	"github.com/aswinjithkukku/url-moulder/initializer"
	"github.com/aswinjithkukku/url-moulder/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDatabase()
	initializer.SyncDatabase()
}

func main() {
	fmt.Println("Hello world")

	r := gin.Default()
	r.LoadHTMLGlob("templates/**")
	r.GET("/", controllers.Home)

	routes.UrlsRoutes(r)

	r.Run(":" + os.Getenv("PORT"))
}
