package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
)

func Encrypt(key, filePath, fileType string) (string, error) {
	plaintext, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	// Save back to file
	cyName := GetFilePath(filePath) + "." + fileType
	err = ioutil.WriteFile(cyName, ciphertext, 0777)
	if err != nil {
		return "", err
	}
	return cyName, nil
}
