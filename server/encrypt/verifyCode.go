package encrypt

import (
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
var n = 4

func GenerateVcode() string {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Nanosecond * time.Duration(rand.Intn(100)))
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
