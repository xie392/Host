package utils

import "math/rand"

// GenerateId 随机生成id的方法
func GenerateId(length int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
