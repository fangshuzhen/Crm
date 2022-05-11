package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"

	uuid "github.com/satori/go.uuid"
)

// 这是一个专门用来生成UUID的方法
func GetUuid() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", -1)
}

// 这是一个专门用来生成MD5的方法
func GetMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
