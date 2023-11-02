package bsp_dungeon

import "math/rand"

type BSP_tree struct {
	width, height int
	root          *node
	alea          *rand.Rand
}

type node struct {
	// node data
	parent   *node
	children []*node // only 2 children even if more are "possible"

	// bsp data
	topLeftX int
	topLeftY int
	width    int
	height   int

	// room or corridor data
	content *room
}

type room struct {
	// room position
	topLeftX int
	topLeftY int
	width    int
	height   int

	// other data
	isRoom       bool // room or corridor
	floorContent int
}
