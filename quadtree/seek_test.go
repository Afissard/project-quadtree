package quadtree

import (
	"fmt"
	"testing"
)

/*
Test de la validité de la fonction Seek

Le test reste valide, bien que désormais MakeFromArray
utilise aussi Seek via AddNode pour la création du quadtree
*/
func Test_Seek_n1(t *testing.T) {
	fmt.Println("\nTest newChunk n°1:")
	// initialisation
	mapContent := [][]int{
		{1, 1, 3, 4},
		{1, 1, 4, 3},
		{0, 0, 2, 2},
		{0, 0, 2, 2},
	}
	quad := MakeFromArray(mapContent) // Utilise désormais Seek pour la création du quadtree ...
	getError := false

	for y := 0; y < len(mapContent); y++ {
		for x := 0; x < len(mapContent[y]); x++ {
			//fonction
			resultNode := quad.Seek(x, y)
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
