package dungeon

import (
	"math/rand"

	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/quadtree"
)

/*
Structure contenant le donjon :
- Arbre BSP
- Quadtree du donjon
*/
type Dungeon struct {
	bspTree bspTree
	quad    quadtree.Quadtree
}

/*
Racine de l'arbre BSP : point de départ de la génération BSP
*/
type bspTree struct {
	width, height int
	topX, topY    int // par défaut (0,0) ?
	root          *bspNode
}

/*
Structure d'une node pour l'arbre BSP du donjon
*/
type bspNode struct {
	width, height int
	topX, topY    int
	nodeLeft      *bspNode
	nodeRight     *bspNode
}

/*
Structure d'un portail, chaque portail est lié à un autre portail qui sert point de sortie.
*/
type Portal struct {
	PosX, PosY int 
	ExitX, ExitY int
	LinkedPortail *Portal
}

var (
	alea *rand.Rand // seed du donjon
)
