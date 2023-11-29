package bspDungeon

import (
	"fmt"
	"os/exec"
	"strconv"
	"testing"
)

func Test_createLevel(t *testing.T) {
	debug := false
	levelTest := CreateLevel("42", 4, 64, 64)
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

func Test_toMapFile_n1(t *testing.T) {
	seed, width, height := "42", 128, 128

	level, floorContent := GetLevel(seed, width, height)
	level.SaveToMapFile(floorContent, seed)

	err := convertToImage(seed, width, height)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_toMapFile_n2(t *testing.T) {
	seed, width, height := "Léa", 128, 128

	level, floorContent := GetLevel(seed, width, height)
	level.SaveToMapFile(floorContent, seed)

	err := convertToImage(seed, width, height)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_toMapFile_n3(t *testing.T) {
	seed, width, height := "Sacha", 128, 128

	level, floorContent := GetLevel(seed, width, height)
	level.SaveToMapFile(floorContent, seed)

	err := convertToImage(seed, width, height)
	if err != nil {
		t.Fatal(err)
	}
}

func convertToImage(seed string, width, height int) (err error) {
	cmd := exec.Command("py", "convert.py", seed, strconv.Itoa(width), strconv.Itoa(height))
	err = cmd.Run()
	return err
}