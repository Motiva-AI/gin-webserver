package main

import (
	"github.com/gin-gonic/gin"
	"github.com/glibs/gin-webserver"
)

func main() {
	server := InitializeServer(":80")
	server.Start()
	// After the server is not useful
	server.Stop()
}

func InitializeServer(host string) (server *network.WebServer) {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.LoadHTMLGlob("./templates/*")
	engine.Static("/public/css/", "./public/css")
	engine.Static("/public/js/", "./public/js/")
	engine.Static("/public/fonts/", "./public/fonts/")
	engine.Static("/public/img/", "./public/img/")

	server.GET("/", GetIndex)
	server.GET("/about", GetAbout)
	server.GET("/contact", GetContact)

	return network.InitializeWebServer(engine, host)
}

func GetIndex(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"variableName": "value",
	})
}

func GetAbout(c *gin.Context) {
	c.HTML(200, "about.html", gin.H{
		"variableName": "value",
	})
}

func GetContract(c *gin.Context) {
	c.HTML(200, "contact.html", gin.H{
		"variableName": "value",
	})
}
