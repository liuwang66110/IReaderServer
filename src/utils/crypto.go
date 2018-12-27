package utils

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func MD5(str string) string {
	data := []byte(str)
	return ByteToMD5(data)
}

func ByteToMD5(data []byte) string {
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func RandMd5() (ret string) {
	nanoTime := time.Now().UnixNano()
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(nanoTime))

	var buf2 = make([]byte, 8)
	nanoTime = time.Now().UnixNano()
	binary.BigEndian.PutUint64(buf2, uint64(nanoTime))
	buf = append(buf, buf2...)
	return ByteToMD5(buf)
}

func GenPwd(pwd string) (hash string) {
	hashB, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hashB)
}

func ChkPwd(cipher string, pwd string) (ret bool) {
	err := bcrypt.CompareHashAndPassword([]byte(cipher), []byte(pwd))
	if err == nil {
		return true
	}
	return false
}
