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
			//fmt.Println(levelTest.roomList[i])
			fmt.Printf("Room %d:%d\twidth: %d, height: %d \n",
				levelTest.roomList[i].topLeftX, levelTest.roomList[i].topLeftY, levelTest.roomList[i].width, levelTest.roomList[i].height)
		}
	}
}

func Test_toMapFile(t *testing.T) {
	seed42 := createLevel(42, 4, 128, 128)
	fmt.Println("room list : ")
	for i := 0; i < len(seed42.roomList); i++ {
		fmt.Printf("Room %d:%d\twidth: %d, height: %d \n",
			seed42.roomList[i].topLeftX, seed42.roomList[i].topLeftY, seed42.roomList[i].width, seed42.roomList[i].height)
	}
	seed42.toMapFile("42")
}
