package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type sourceServer struct {
	Server   string `json:"server"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var servers = []sourceServer{}

// encrypt a string with a key
func encrypt(key []byte, text string) string {
	return text
}

func decrypt(key []byte, text string) string {
	return text
}

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
