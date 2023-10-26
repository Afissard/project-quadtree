package floor

import (
	"testing"
)

func Test_readFloorFromFile(t *testing.T) {
	/*
		only test for now the "exemple" file map
	*/
	filepath := "..\\floor-files\\exemple"
	floorContent := readFloorFromFile(filepath)
	expectedContent := [][]int{
		{1, 1, 3, 4},
		{1, 1, 4, 3},
		{0, 0, 2, 2},
		{0, 0, 2, 2},
	}
	if floorContent == nil {
		t.Fatal("floorContent est vide")
	}
	for i := 0; i < len(expectedContent); i++ {
		for j := 0; j < len(expectedContent[i]); j++ {
			if floorContent[i][j] != expectedContent[i][j] {
				t.Fatalf("For p = %q, expected %d. Got %d.", filepath, expectedContent[i], floorContent[i])
			}
		}
	}
}
