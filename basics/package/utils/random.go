package utils

import "math/rand"

func RandomInt(max int) int {
	return rand.Intn(max)
}
