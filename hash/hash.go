package hash

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

// Salt for hasing
const Salt = "andpepper"

// Md5 string
func Md5(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s + Salt))

	return hex.EncodeToString(hasher.Sum(nil))
}

// GenerateMd5 base on timestamp
func GenerateMd5() string {
	hasher := md5.New()
	hasher.Write([]byte(time.Now().String() + Salt))

	return hex.EncodeToString(hasher.Sum(nil))
}
