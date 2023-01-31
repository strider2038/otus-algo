package datatesting

import (
	"math/rand"
	"time"
)

var defaultChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ 0123456789")

func GenerateRandomString(n int, chars ...rune) string {
	if len(chars) == 0 {
		chars = defaultChars
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func GenerateRandomStrings(n int, chars ...rune) []string {
	if len(chars) == 0 {
		chars = defaultChars
	}

	rand.Seed(time.Now().UnixNano())

	ss := make([]string, n)
	for i := 0; i < n; i++ {
		ss[i] = GenerateRandomString(15, chars...)
	}

	return ss
}
