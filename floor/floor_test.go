package floor

import (
	"testing"
)

func Test_readFloorFromFile_n1(t *testing.T) {
	/*
		test the "exemple" file map
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
				t.Fatalf("Pour la map : %q \nAttend : %d \nReçoit : %d", filepath, expectedContent[i], floorContent[i])
			}
		}
	}
}

func Test_readFloorFromFile_n2(t *testing.T) {
	/*
		test the "validation" file map
	*/
	filepath := "..\\floor-files\\validation"
	floorContent := readFloorFromFile(filepath)
	expectedContent := [][]int{
		{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
		{1, 2, 2, 2, 1, 2, 2, 2, 1, 2, 2, 2, 1, 2, 1, 2, 1, 2, 2, 2, 1},
		{0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0},
		{1, 2, 2, 2, 1, 2, 2, 2, 1, 2, 2, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1},
		{0, 2, 0, 2, 0, 2, 2, 1, 0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0},
		{1, 2, 2, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 0, 2, 0, 1, 2, 2, 2, 1},
		{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
	}
	if floorContent == nil {
		t.Fatal("floorContent est vide")
	}
	for i := 0; i < len(expectedContent); i++ {
		for j := 0; j < len(expectedContent[i]); j++ {
			if floorContent[i][j] != expectedContent[i][j] {
				t.Fatalf("Pour la map : %q \nAttend : %d \nReçoit : %d", filepath, expectedContent[i], floorContent[i])
			}
		}
	}
}
