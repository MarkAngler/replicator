package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

func Encrypt(key, plaintext string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	plaintextBytes := []byte(plaintext)
	ciphertextBytes := make([]byte, len(plaintextBytes))
	stream := cipher.NewCTR(block, []byte(key[:block.BlockSize()]))
	stream.XORKeyStream(ciphertextBytes, plaintextBytes)

	return base64.StdEncoding.EncodeToString(ciphertextBytes), nil
}

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

func GenerateSecureKey(length int) (string, error) {
	key := make([]byte, length)

	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return "", err
	}

	return hex.EncodeToString(key), nil
}
