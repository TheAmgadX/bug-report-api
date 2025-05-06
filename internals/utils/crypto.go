package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"os"

	"github.com/joho/godotenv"
)

var encryptionKey []byte

// this function called once when package is loaded to load the key from .env file.

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("invalid reading for .env file")
	}

	keyHex := os.Getenv("ENCRYPTION_KEY")

	key, err := hex.DecodeString(keyHex)

	if err != nil || len(key) != 32 {
		panic("Invalid ENCRYPTION_KEY in .env file")
	}

	encryptionKey = []byte(key)
}

func Encrypt(str string) (string, error) {
	plainText := []byte(str)

	block, err := aes.NewCipher(encryptionKey)

	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return "", err
	}

	// number used once.
	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := gcm.Seal(nonce, nonce, plainText, nil)

	return hex.EncodeToString(cipherText), nil
}

func Decrypt(cipherHex string) (s string, err error) {

	cipherData, err := hex.DecodeString(cipherHex)

	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(encryptionKey)

	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return "", err
	}
	
	nonceSize := gcm.NonceSize()

	if nonceSize > len(cipherData) {
		return "", errors.New("cipher text is too short")
	}

	nonce, cipherText := cipherData[:nonceSize], cipherData[nonceSize:]


	plainText, err := gcm.Open(nil, nonce, cipherText, nil)

	if err != nil {
		return "", err
	}

	
	return string(plainText), nil	
}
