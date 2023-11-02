package bsp_dungeon

import (
	"fmt"
	"testing"
)

func Test_createLevel(t *testing.T) {
	debug := false
	levelTest := createLevel(42, 4, 64, 64)
	if levelTest.root == nil {
		t.Fatalf("\nL'étage est vide !\n")
	} else if debug {
		fmt.Println("room list : ")
		for i := 0; i < len(levelTest.roomList); i++ {
			fmt.Println(levelTest.roomList[i])
		}
	}
}

func Test_toMapFile(t *testing.T) {
	seed42 := createLevel(42, 4, 128, 128)
	seed42.toMapFile("42")
}
