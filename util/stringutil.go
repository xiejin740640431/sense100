package util

import (
	"time"
	"math/rand"
	"strconv"
	"crypto/md5"
	"fmt"
	"encoding/hex"
)

func IsEmpty(str string) bool {
	if str == "" {
		return true
	}
	return false
}

func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetImageFileName(suffix string) string {
	var timeStr = strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Println(timeStr)
	var h = md5.New()
	h.Write([]byte(timeStr))
	var resultByte = h.Sum(nil)
	return hex.EncodeToString(resultByte) + suffix
}
