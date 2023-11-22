package chunk

import (
	"github.com/Afissard/project-quadtree/configuration"
	"github.com/Afissard/project-quadtree/quadtree"
)

func createChunk(topX, topY, content int) (q quadtree.Quadtree) {
	// Création du contenu du chunk
	floorContent := [][]int{}

	if configuration.Global.InfiniteMap == 1 {
		for y := 0; y < CHUNK_SIZE; y++ {
			line := []int{}
			for x := 0; x < CHUNK_SIZE; x++ {
				line = append(line, content)
			}
			floorContent = append(floorContent, line)
		}
	} else if configuration.Global.InfiniteMap == 2 {
		panic("Not yet Implemented")
	} else {
		panic("Shouldn't ask for chunk generation")
	}

	// Ajout du chunk au quadtree général

	return quadtree.MakeFromArray(floorContent)
}
