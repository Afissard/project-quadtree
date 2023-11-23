package quadtree

import (
	"fmt"

	"github.com/Afissard/project-quadtree/configuration"
)

func (q Quadtree) CreateChunk(topX, topY, content int) {
	// Création du contenu du chunk
	floorContent := [][]int{}

	inTest := true
	if configuration.Global.InfiniteMap == 1 || inTest {
		for y := 0; y < CHUNK_SIZE; y++ {
			line := []int{}
			for x := 0; x < CHUNK_SIZE; x++ {
				line = append(line, content)
			}
			floorContent = append(floorContent, line)
		}
	} else if configuration.Global.InfiniteMap == 2 {
		err := fmt.Errorf(fmt.Sprintf("\nWave Fonction Collapse not yet Implemented but config.InfiniteMap = %d", configuration.Global.InfiniteMap))
		fmt.Println(err.Error())
	} else {
		panic("\nShouldn't ask for chunk generation")
	}

	// Ajout du chunk au quadtree général
	// TODO: mettre à jour la taille du quadtree
	q.AppendToQuadtree(floorContent, topX, topY)
}
