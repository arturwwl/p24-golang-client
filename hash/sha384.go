package hash

import (
	"crypto"
	_ "crypto/sha512"
	"encoding/hex"
)

func CalcSha384(baseString string) string {
	sha384 := crypto.SHA384.New()
	sha384.Write([]byte(baseString))
	signBytes := sha384.Sum(nil)
	return hex.EncodeToString(signBytes)
}
