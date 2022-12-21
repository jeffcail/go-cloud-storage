package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// Sha256
func Sha256(str string) string {
	srcBytes := []byte(str)
	hash := sha256.New()
	hash.Write(srcBytes)
	hashBytes := hash.Sum(nil)
	sha256String := hex.EncodeToString(hashBytes)
	return sha256String
}
