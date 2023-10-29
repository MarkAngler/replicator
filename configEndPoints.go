package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type sourceServer struct {
	Server   string `json:"server"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var servers = []sourceServer{}

func createYaml(y sourceServer, fileName string) ([]byte, error) {
	data, err := yaml.Marshal(y)
	if err != nil {
		fmt.Println(err)
		return nil, err

	}
	os.WriteFile(fileName, data, 0644)
	return data, nil
}

func postSourceServers(c *gin.Context) {
	var newServer sourceServer
	if err := c.ShouldBind(&newServer); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Hash the password using the separate HashPassword function
	hashedPassword, err := Encrypt("key12349590295afdafds02489013741", newServer.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not hash password"})
		return
	}
	newServer.Password = hashedPassword
	//servers = append(servers, newServer)
	createYaml(newServer, "config.yaml")
	c.IndentedJSON(http.StatusCreated, newServer)
}

func getSourceServers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, servers)
}
