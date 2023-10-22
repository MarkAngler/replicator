package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type sourceServer struct {
	Server string `json:"server"`
	Port   string `json:"port"`
}

var servers = []sourceServer{}

func postSourceServers(c *gin.Context) {
	var newServer sourceServer
	if err := c.BindJSON(&newServer); err != nil {
		return
	}

	servers = append(servers, newServer)
	c.IndentedJSON(http.StatusCreated, newServer)
}

func getSourceServers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, servers)
}

func main() {
	r := gin.Default()

	r.POST("/sourceServers", postSourceServers)
	r.GET("/sourceServers", getSourceServers)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
