package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type sourceServer struct {
	Server   string `json:"server"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var servers = []sourceServer{}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func postSourceServers(c *gin.Context) {
	var newServer sourceServer
	if err := c.ShouldBind(&newServer); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Hash the password using the separate HashPassword function
	hashedPassword, err := HashPassword(newServer.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not hash password"})
		return
	}
	newServer.Password = hashedPassword

	servers = append(servers, newServer)
	c.IndentedJSON(http.StatusCreated, newServer)
}

func getSourceServers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, servers)
}
