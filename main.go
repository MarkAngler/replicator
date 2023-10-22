package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type sourceServer struct {
	server string `json:"server"`
	port   string `json:"port"`
}

var sourceServerList = []sourceServer{}

func postSourceServer(c *gin.Context) {
	var newServer sourceServer
	if err := c.BindJSON(&newServer); err != nil {
		return
	}

	sourceServerList = append(sourceServerList, newServer)
	c.IndentedJSON(http.StatusCreated, newServer)
}

func main() {
	r := gin.Default()

	r.POST("/newSourceServer", postSourceServer)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
