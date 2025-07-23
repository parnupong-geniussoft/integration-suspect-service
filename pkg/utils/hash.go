package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func StringHash(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashed := hasher.Sum(nil)
	return base64.StdEncoding.EncodeToString(hashed)
}
