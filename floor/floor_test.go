package floor

import (
	"fmt"
	"testing"
)

func Test_readFloorFromFile(t *testing.T) {
	floorContent := readFloorFromFile("../floor-file/exemple")
	if floorContent == nil {
		t.Fatal("floorContent est vide")
	}
	fmt.Println("floorContent", floorContent)
}
