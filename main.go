package main

import (
	"fmt"
	"os"

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

	routes.UrlsRoutes(r)

	r.Run(":" + os.Getenv("PORT"))
}
