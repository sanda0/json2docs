package json2docs

import (
	"math/rand"
	"strconv"
	"time"
)

func IsDigit(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}
func RandStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
