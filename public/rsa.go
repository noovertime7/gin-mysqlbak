package public

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

func RsaDecode(strPlainText string) (string, error) {
	plainText, err := base64.StdEncoding.DecodeString(strPlainText)
	bytePrivateKey := []byte(PrivateKey)
	priBlock, _ := pem.Decode(bytePrivateKey)
	priKey, err := x509.ParsePKCS8PrivateKey(priBlock.Bytes)
	if err != nil {
		return "", err
	}
	decryptText, err := rsa.DecryptPKCS1v15(rand.Reader, priKey.(*rsa.PrivateKey), plainText)
	if err != nil {
		return "", err
	}
	return string(decryptText), nil
}

func RsaEncode(plain string) (string, error) {
	msg := []byte(plain)
	pubBlock, _ := pem.Decode([]byte(PublicKey))
	pubKeyValue, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		return "", err
	}
	pub := pubKeyValue.(*rsa.PublicKey)
	encryText, err := rsa.EncryptPKCS1v15(rand.Reader, pub, msg)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryText), nil
}
