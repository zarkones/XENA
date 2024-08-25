package slices

import (
	"fmt"
	"testing"
)

func TestChunk(t *testing.T) {
	x := []string{"x", "y", "z"}
	y := Chunk(x, 3)
	fmt.Println(y)
	if len(y) != 3 {
		t.Log("unexpected len:", len(y))
		t.FailNow()
	}
}

func TestChunk2(t *testing.T) {
	x := []string{"x", "y", "z"}
	y := Chunk(x, 2)
	fmt.Println(y)
	if len(y) != 2 {
		t.Log("unexpected len:", len(y))
		t.FailNow()
	}
}
