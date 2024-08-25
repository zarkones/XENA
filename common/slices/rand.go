package slices

import (
	"math/rand"
)

func Rand[T any](slice *[]T) T {
	return (*slice)[rand.Intn(len(*slice))]
}
