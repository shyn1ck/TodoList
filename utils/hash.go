package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateHash(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	hashedBytes := hash.Sum(nil)
	return hex.EncodeToString(hashedBytes)
}
