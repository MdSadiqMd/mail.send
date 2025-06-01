package utils

import (
	"crypto/rand"
	"encoding/base64"

	logger "github.com/MdSadiqMd/mail.send/pkg/log"
)

var helperLogger = logger.New("RandomStrings")

func GenerateRandomBytes(length int) ([]byte, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		helperLogger.Error("Failed to generate random bytes: %v", err)
		return nil, err
	}
	return bytes, nil
}

func GenerateRandomString(s int) (string, error) {
	bytes, err := GenerateRandomBytes(s)
	if err != nil {
		helperLogger.Error("Failed to generate random string: %v", err)
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}
