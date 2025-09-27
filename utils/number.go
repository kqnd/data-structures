package utils

import "math/rand"

func Generate(min, max int) int {
	return rand.Intn(int(max)-int(min) + 1) + int(min)
}