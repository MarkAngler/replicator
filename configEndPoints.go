package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
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

func Encrypt(key, plaintext string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	plaintextBytes := []byte(plaintext)
	ciphertextBytes := make([]byte, len(plaintextBytes))
	stream := cipher.NewCTR(block, []byte(key[:block.BlockSize()]))
	stream.XORKeyStream(ciphertextBytes, plaintextBytes)

	return base64.StdEncoding.EncodeToString(ciphertextBytes), nil
}

// Decrypt decrypts the ciphertext with the given key using AES.
func Decrypt(key, ciphertext string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	if len(ciphertextBytes) < block.BlockSize() {
		return "", errors.New("ciphertext too short")
	}

	plaintextBytes := make([]byte, len(ciphertextBytes))
	stream := cipher.NewCTR(block, []byte(key[:block.BlockSize()]))
	stream.XORKeyStream(plaintextBytes, ciphertextBytes)

	return string(plaintextBytes), nil
}

func postSourceServers(c *gin.Context) {
	var newServer sourceServer
	if err := c.ShouldBind(&newServer); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Hash the password using the separate HashPassword function
	hashedPassword, err := Encrypt("1", newServer.Password)
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
