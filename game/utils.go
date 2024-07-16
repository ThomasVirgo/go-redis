package game

import "math/rand"

func GenerateID(n int) string {
	var letters = []rune("ABCDEFGHJKLMNPQRSTUVWXYZ123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
