package quadtree

import (
	"fmt"

	"github.com/Afissard/project-quadtree/configuration"
)

func (q Quadtree) CreateChunk(topX, topY int) {
	//TODO: ajouté test configuration -> infiniteMap
	// Création du contenu du chunk (TODO: à modifié avec implémentation de la génération avec WFC)
	floorContent := [][]int{}
	if configuration.Global.InfiniteMap == 2 {
		// TODO: génération WFC
		fmt.Println("not yet implemented")
	} else {
		floorContent = createSimpleChunkContent(2)
	}
	newChunkTree := MakeFromArray(floorContent) // quadtree temporaire, todo: modif coords

	// Trouve où ajouter le chunk généré dans le quadtree général
	nodeToMod := q.seek(topX, topY)
	// Ajout du chunk au quadtree général
	q.AppendToQuadtree(nodeToMod, newChunkTree.root)
}

func createSimpleChunkContent(content int) (floorContent [][]int) {
	// Création du contenu d'un chunk simple et uniforme
	floorContent = make([][]int, CHUNK_SIZE)
	for y := 0; y < len(floorContent); y++ {
		for x := 0; x < len(floorContent[y]); x++ {
			floorContent[y][x] = content
		}
	}
	return floorContent
}
