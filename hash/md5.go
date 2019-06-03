package hash

import (
	"crypto/md5"
	"fmt"
	"encoding/hex"
	"time"
)

// Salt for hashing
var Salt = "andpepper"

// MD5 returns a hashed string.
func MD5(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s + Salt))

	return hex.EncodeToString(hasher.Sum(nil))
}

// GenerateMD5 returns a generated hash.
func GenerateMD5() string {
	hasher := md5.New()
	t := time.Now().UnixNano()
	hasher.Write([]byte(fmt.Sprintf("%d %s", t, Salt)))

	return hex.EncodeToString(hasher.Sum(nil))
}