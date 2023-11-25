package quadtree

func (q Quadtree) AppendToQuadtree(nodeToMod, newChunkTree *node) {
	/*
		Algo :
		1. détermine où couper par rapport à l'input
			- cas où il faut juste rajouté une branche à la node
			- cas où il faut créer une node supérieur
		2. reconstruit l'arbre (attention aux dimensions des node au dessus de la découpe)
	*/
	if newChunkTree.topLeftX < nodeToMod.topLeftX+nodeToMod.width/2 {
		if newChunkTree.topLeftY < nodeToMod.topLeftY+nodeToMod.height/2 {
			// tl
			nodeToMod.topLeftNode = newChunkTree
		} else {
			// bl
			nodeToMod.bottomLeftNode = newChunkTree
		}
	} else if newChunkTree.topLeftX <= nodeToMod.topLeftX+nodeToMod.width/2 {
		if newChunkTree.topLeftY < nodeToMod.topLeftY+nodeToMod.height/2 {
			// tr
			nodeToMod.topRightNode = newChunkTree
		} else {
			// br
			nodeToMod.bottomRightNode = newChunkTree
		}
	} else {
		// créer la nouvelle node "root"
		oldRootNode := q.root
		rootNode := node{
			topLeftX: 0,
			topLeftY: 0,
			width:    oldRootNode.width + newChunkTree.width,
			height:   oldRootNode.height + newChunkTree.height,
			content:  -1,
		}
		q.AppendToQuadtree(&rootNode, oldRootNode)
		q.root = &rootNode
		// réessaye de rajouté le chunk
		q.AppendToQuadtree(q.root, newChunkTree)
	}
}
