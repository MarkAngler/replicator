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
	Type     string `json:"type"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Config is the top-level configuration structure.
// Servers is a nested structure containing a map of server names to details.
// ServerName is a map where the key is the server name and the value is the server details.
// Config.Servers.ServerName["testservername"]

var Config struct {
	Servers struct {
		ServerName map[string]ServerDetails `yaml:"Servers"`
	}
}

// ServerDetails holds all the configuration details for a single server.
type ServerDetails struct {
	Host string `yaml:"host"`

	Port string `yaml:"port"`

	Type string `yaml:"type"`

	Database string `yaml:"database"`

	Username string `yaml:"username"`

	Password string `yaml:"password"`
}

func initializeYaml() {
	allServerYaml, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(allServerYaml, &Config)
	if err != nil {
		fmt.Println(err)
	}
}

func addSourceServerToYaml(y sourceServer) ([]byte, error) {

	data, err := yaml.Marshal(y)

	if err != nil {
		fmt.Println(err)
		return nil, err

	}
	os.WriteFile("config.yaml", data, 0644)
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

	addSourceServerToYaml(newServer)
	c.IndentedJSON(http.StatusCreated, newServer)
}

// func getSourceServers(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, servers)
// }
