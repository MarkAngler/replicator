package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type sourceServer struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Type     string `json:"type"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Example
// Config.Servers['server1'].Host
var Config struct {
	Servers map[string]ServerDetails `yaml:"Servers"`
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

	Config.Servers[y.Host] = ServerDetails{
		Host:     y.Host,
		Port:     y.Port,
		Type:     y.Type,
		Database: y.Database,
		Username: y.Username,
		Password: y.Password,
	}

	data, err := yaml.Marshal(&Config)

	if err != nil {
		fmt.Println(err)
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

func getSourceServers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Config.Servers)
}
