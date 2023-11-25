package quadtree

func (q Quadtree) AppendToQuadtree(nodeToMod, newNode *node) {
	/*
		Algo :
		1. détermine où couper par rapport à l'input
			- cas où il faut juste rajouté une branche à la node
			- cas où il faut créer une node supérieur
		2. reconstruit l'arbre (attention aux dimensions des node au dessus de la découpe)
	*/
	if newNode.topLeftX < nodeToMod.topLeftX+nodeToMod.width/2 {
		if newNode.topLeftY < nodeToMod.topLeftY+nodeToMod.height/2 {
			// tl
			nodeToMod.topLeftNode = newNode
		} else {
			// bl
			nodeToMod.bottomLeftNode = newNode
		}
	} else if newNode.topLeftX <= nodeToMod.topLeftX+nodeToMod.width/2 {
		if newNode.topLeftY < nodeToMod.topLeftY+nodeToMod.height/2 {
			// tr
			nodeToMod.topRightNode = newNode
		} else {
			// br
			nodeToMod.bottomRightNode = newNode
		}
	} else {
		// créer la nouvelle node "root"
		oldRootNode := q.root
		rootNode := node{
			topLeftX: 0,
			topLeftY: 0,
			width:    oldRootNode.width + newNode.width,
			height:   oldRootNode.height + newNode.height,
			content:  -1,
		}
		q.AppendToQuadtree(&rootNode, oldRootNode)
		q.root = &rootNode
		// réessaye de rajouté le chunk
		q.AppendToQuadtree(q.root, newNode)
	}
}
