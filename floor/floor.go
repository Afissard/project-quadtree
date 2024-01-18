package floor

import (
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/quadtree"
)

// Floor représente les données du terrain. Pour le moment
// aucun champs n'est exporté.
//
//   - content : partie du terrain qui doit être affichée à l'écran
//   - fullContent : totalité du terrain (utilisé seulement avec le type
//     d'affichage du terrain "fromFileFloor")
//   - quadTreeContent : totalité du terrain sous forme de quadtree (utilisé
//     avec le type d'affichage du terrain "quadtreeFloor")
type Floor struct {
	content         [][]int
	fullContent     [][]int
	QuadtreeContent quadtree.Quadtree
	// animationStep   int // scroll du floor
}

// types d'affichage du terrain disponibles
const (
	gridFloor int = iota
	fromFileFloor
	quadTreeFloor
	bspDungeonFloor
	infiniteWorld
)

var (
	fk3BlockList = []int{-1, 0, 6, 7, 8}
)
