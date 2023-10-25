package floor

import (
	"fmt"
	"testing"
)

func check(err error){
	if err != nil {
		panic(err)
	}
}

func Test_readFloorFromFile(t *testing.T) {
	floorContent := readFloorFromFile("..\\floor-files\\exemple")
	if floorContent == nil {
		t.Fatal("floorContent est vide")
	}
	fmt.Println("floorContent", floorContent)
}
