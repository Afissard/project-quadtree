package quadtree

/*
Initialise un quadtree avec les dimension données.
*/
func (q *Quadtree) Init(width, height int) {
	q.Width = width
	q.Height = height

	q.root = &node{
		topLeftX: 0,
		topLeftY: 0,
		width:    q.Width,
		height:   q.Height,
		content:  -1,
	}
}
