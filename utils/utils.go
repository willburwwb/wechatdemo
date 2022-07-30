package utils

import (
	"math/rand"
	"time"
)

func CreateVerifyCode() string {
	var letters = []byte("123456789")
	verifyCode := make([]byte, 6)
	rand.Seed(time.Now().Unix())
	for i := range verifyCode {
		verifyCode[i] = letters[rand.Intn(len(letters))] //随机数
	}
	return string(verifyCode)
}
