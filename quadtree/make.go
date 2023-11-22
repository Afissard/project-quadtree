package quadtree

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.
func MakeFromArray(floorContent [][]int) (q Quadtree) {
	totalWidth := len(floorContent[0])
	totalHeight := len(floorContent)

	rootNode := node{
		topLeftX: 0,
		topLeftY: 0,
		width:    totalWidth,
		height:   totalHeight,
		content:  -1, // contenu à -1 par défaut
	}

	q = Quadtree{
		width:  totalWidth,
		height: totalHeight,
		root:   &rootNode,
		// rajouté variable pour support coords négatives
	}
	for y := 0; y < len(floorContent); y++ {
		for x := 0; x < len(floorContent[y]); x++ {
			AddContent(&rootNode, floorContent[y][x], x, y)
		}
	}
	return q
}
