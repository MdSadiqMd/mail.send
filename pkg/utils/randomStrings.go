package randomstrings

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomBytes(length int) ([]byte, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func GenerateRandomString(s int) (string, error) {
	bytes, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(bytes), err
}
