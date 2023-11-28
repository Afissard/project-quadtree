package quadtree

import (
	"fmt"
	"testing"
)

func Test_seek_n1(t *testing.T) {
	fmt.Println("\nTest newChunk n°1:")
	// initialisation
	mapContent := [][]int{
		{1, 1, 3, 4},
		{1, 1, 4, 3},
		{0, 0, 2, 2},
		{0, 0, 2, 2},
	}
	quad := MakeFromArray(mapContent)
	getError := false

	for y := 0; y < len(mapContent); y++ {
		for x := 0; x < len(mapContent[y]); x++ {
			//fonction
			resultNode := quad.seek(x, y)
			// test
			if resultNode.content != mapContent[y][x] {
				t.Fatalf("\n Contenu de la node différent de la map : %d != %d", resultNode.content, mapContent[y][x])
				getError = true
			}
		}
	}
	if !getError {
		fmt.Println("ok")
	}
}
