package codes

import (
	"crypto/rand"
	"encoding/hex"
)

func generateCode() (string, error) {
	token := make([]byte, 8)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(token), nil
}
