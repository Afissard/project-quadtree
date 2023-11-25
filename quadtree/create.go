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

	// la node qui contiendra le chunk
	newChunkTree := node{
		topLeftX: topX,
		topLeftY: topY,
		width:    CHUNK_SIZE,
		height:   CHUNK_SIZE,
		content:  -1,
	}

	for y := 0; y < len(floorContent); y++ {
		for x := 0; x < len(floorContent[y]); x++ {
			AddContent(&newChunkTree, floorContent[y][x], topX+x, topY+y)
		}
	}

	// Trouve où ajouter le chunk généré dans le quadtree général
	nodeToMod := q.FindWhereMod(topX, topY)
	// Ajout du chunk au quadtree général
	q.AppendToQuadtree(nodeToMod, &newChunkTree)
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
