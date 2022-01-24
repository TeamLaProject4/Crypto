package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetHash(valueToHash string) string {
	sha256Hasher := sha256.New()
	sha256Hasher.Write([]byte(valueToHash))
	hashBytes := sha256Hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
