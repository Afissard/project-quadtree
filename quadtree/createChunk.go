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
	for y := 0; y < CHUNK_SIZE; y++ {
		line := []int{}
		for x := 0; x < CHUNK_SIZE; x++ {
			line = append(line, content)
		}
		floorContent = append(floorContent, line)
	}
	fmt.Println("floorContent: ", floorContent)
	return floorContent
}
