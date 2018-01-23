package hash

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

// Salt for hasing
const Salt = "andpepper"

// Hash string
func Hash(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s + Salt))

	return hex.EncodeToString(hasher.Sum(nil))
}

// GenerateHash base on timestamp
func GenerateHash() string {
	hasher := md5.New()
	hasher.Write([]byte(time.Now().String() + Salt))

	return hex.EncodeToString(hasher.Sum(nil))
}
