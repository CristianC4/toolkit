package toolkit

import (
	"crypto/rand"
)

const randomString = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+_"

func (t *Tools) RamdonsString(length int) string {
	s, randomRune := make([]rune, length), []rune(randomString)
	for index := range s {
		prime, _ := rand.Prime(rand.Reader, len(randomRune))
		x, y := prime.Uint64(), uint64(len(randomRune))
		s[index] = randomRune[x%y]
	}
	return string(s)
}

type Tools struct{}
