package service

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateSessionId() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}
