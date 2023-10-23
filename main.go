package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/form", func(c *gin.Context) {
		c.HTML(200, "form.html", nil)
	})

	r.POST("/sourceServers", postSourceServers)
	r.GET("/sourceServers", getSourceServers)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
