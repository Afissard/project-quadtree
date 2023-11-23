package quadtree

import (
	"fmt"

	"github.com/Afissard/project-quadtree/configuration"
)

func (q Quadtree) CreateChunk(topX, topY, content int) {
	// Création du contenu du chunk
	floorContent := make([][]int, CHUNK_SIZE)

	inTest := true // TODO: retiré à la fin des test ou trouver alternative

	if configuration.Global.InfiniteMap == 1 || inTest {
		for y := 0; y < len(floorContent); y++ {
			for x := 0; x < len(floorContent[y]); x++ {
				floorContent[y][x] = content
			}
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
