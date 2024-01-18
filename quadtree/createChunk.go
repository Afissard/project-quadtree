package quadtree

import (
	"fmt"
	"log"

	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/configuration"
)
/*
Créer un chunk avec l'algo choisi dans le fichier config.json aux coordonnées fournis
*/
func (q *Quadtree) CreateChunk(topX, topY int) {
	var err error
	switch configuration.Global.AlgoGeneration {
	case 0: // chunk simple
		q.CreateSimpleChunk(topX, topY, 4)
	default:
		err = fmt.Errorf("algo génération n°%d non supporté", configuration.Global.AlgoGeneration)
	}
	if err != nil {
		log.Panic(err)
	}
}

/*
Fonction CreateChunkFrom (fichier createChunk)

Entrées: topX, topY, floorContent -> Coordonnées du point le plus en haut à gauche du chunk à créer ainsi que le contenu du chunk
Sorties: le quadtree mis à jour avec le nouveau chunk

Le but de cette fonction est de créer un chunk rempli du contenu défini par les paramètre
de génération infini et de l'ajouté au quadtree existant.
*/
func (q *Quadtree) CreateChunkFrom(topX, topY int, floorContent [][]int) {
	for y := 0; y < len(floorContent); y++ {
		for x := 0; x < len(floorContent[y]); x++ {
			q.AppendToQuadtree(x+topX, y+topY, floorContent[y][x])

		}
	}
}

/*
Fonction createSimpleChunkContent (fichier createChunk)

Entrées: topX, topY, content -> coordonnée en haut à gauche du chunk et valeur du contenu du chunk
Sorties: le quadtree mis à jour avec le nouveau chunk

La fonction créer un chunk uniforme selon la valeur du
contenu donnée en entrée. Ces chunk sont utilisé pour
la génération de monde infini procéduraux "simple" (sans
algorithme de génération aléatoire, exemple : WFC)
*/
func (q *Quadtree) CreateSimpleChunk(topX, topY, content int) {
	floorContent := [][]int{}
	for y := 0; y < CHUNK_SIZE; y++ {
		line := []int{}
		for x := 0; x < CHUNK_SIZE; x++ {
			line = append(line, content)
		}
		floorContent = append(floorContent, line)
	}
	q.CreateChunkFrom(topX, topY, floorContent)
}
