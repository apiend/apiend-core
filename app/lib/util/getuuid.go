package util

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/gogf/gf/g/util/grand"
)

// 生成一下UUID
func GenerateUUID() string {
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		return ""
	}
	uuid[8] = 0x80 // variant bits see page 5
	uuid[4] = 0x40 // version 4 Pseudo Random, see page 7
	return hex.EncodeToString(uuid)
}

// 生成一个 token key
func GenerateToken(n int) string {
	token := grand.Str(n)
	return token
}
