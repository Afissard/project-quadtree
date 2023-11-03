package bsp_dungeon

import (
	"fmt"
	"os/exec"
	"strconv"
	"testing"
)

func Test_createLevel(t *testing.T) {
	debug := false
	levelTest := createLevel("42", 4, 64, 64)
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
	seed, width, height := "42", 512, 512
	seed42 := createLevel(seed, 3, width, height)
	seed42.toMapFile(seed)

	if debug {
		err := convertToImage(seed, width, height)
		if err != nil {
			t.Fatal(err)
		}
	}

	seed = "Léa"
	lea := createLevel(seed, 3, width, height)
	lea.toMapFile(seed)
	if debug {
		err := convertToImage(seed, width, height)
		if err != nil {
			t.Fatal(err)
		}
	}

	seed = "Sacha"
	sacha := createLevel(seed, 3, width, height)
	sacha.toMapFile(seed)
	if debug {
		err := convertToImage(seed, width, height)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func convertToImage(seed string, width, height int) (err error) {
	cmd := exec.Command("py", "convert.py", seed, strconv.Itoa(width), strconv.Itoa(height))
	err = cmd.Run()
	return err
}
