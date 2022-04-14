package hasher

import (
	"crypto/sha256"
	"encoding/hex"
)

func CreateHash(sid string) string {
	return hex.EncodeToString(CreateHashBytes([]byte(sid)))
}

func CreateHashBytes(data []byte) []byte {
	hasher := sha256.New()
	hasher.Write(data)

	return hasher.Sum(nil)
}
