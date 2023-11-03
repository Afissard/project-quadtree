package bsp_dungeon

import (
	"fmt"
	"os/exec"
	"strconv"
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
	debug := true
	width, height := 256, 256
	seed42 := createLevel(42, 3, width, height)
	seed42.toMapFile("42")

	if debug {
		cmd := exec.Command("py", "convert.py", strconv.Itoa(width), strconv.Itoa(height))
		err := cmd.Run()
		if err != nil {
			t.Fatal(err)
		}
	}
}
