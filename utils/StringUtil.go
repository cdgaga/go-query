package utils

import (
	"crypto/md5"
	"fmt"
)


func StringMd5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}
