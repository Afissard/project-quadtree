package quadtree

func (q Quadtree) AppendToQuadtree(floorContent [][]int, topX, topY int) {
	for y := 0; y < len(floorContent); y++ {
		for x := 0; x < len(floorContent[y]); x++ {
			AddContent(q.root, floorContent[y][x], x+topX, y+topY)
		}
	}
}
