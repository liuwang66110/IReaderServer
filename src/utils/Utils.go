package utils

import (
	"math/rand"
	"time"
)

//生成随机字符串
func GetRandomString(strSize int) string{
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < strSize; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
