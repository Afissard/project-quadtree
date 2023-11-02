package bsp_dungeon

import (
	"testing"
)

func Test_createLevel(t *testing.T) {
	levelTest := createLevel(42, 4, 64, 64)
	if levelTest.root == nil {
		t.Fatalf("\nL'étage est vide !\n")
	}
}
