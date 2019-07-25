package random

import (
	"math/rand"
	"time"
)

func GetRandom(n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += string(byte(97 + getIndex()))
	}
	return result
}

func getIndex() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(26)
}
