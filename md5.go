package tools

import (
	"crypto/md5"
	"fmt"
)

func Md5(str string) string {
	sum := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", sum)
}
