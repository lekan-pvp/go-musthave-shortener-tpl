package app

import "math/rand"

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	prefix = "http://localhost:8080/"
)

func Shorting() string {
	b := make([]byte, 5)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return prefix+string(b)
}
