package bsp_dungeon

import "math/rand"

/*
Structure détenant l'arbre BSP d'un niveau
width, height : dimension du niveau
root : départ de l'arbre BSP
alea : seed du niveau
roomList : liste de salle du niveau
*/
type BSP_tree struct {
	width, height int
	root          *node
	alea          *rand.Rand // seed du niveau
	roomList      []*room    // salles et couloirs du niveau
}

type node struct {
	// node data
	parent   *node
	children []*node // uniquement 2 enfant (même si plus "possible")

	// bsp data
	topLeftX int
	topLeftY int
	width    int
	height   int

	// salles / couloir
	content *room
}

type room struct {
	// position
	topLeftX int
	topLeftY int
	width    int
	height   int

	// autres information
	isRoom       bool // salle ou couloir
	floorContent int  // TODO: à remplacé par une liste 2D, représentant chaque tuile
}
