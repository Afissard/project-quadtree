package quadtree

// fonction interface, sert à rechercher une node dans un quadtree
func (q Quadtree) seek(targetX, targetY int) (nodeFound *node) {
	return seeker(q.root, targetX, targetY)
}

func seeker(currentNode *node, targetX, targetY int) (nodeFound *node) {
	if currentNode.topLeftNode == nil && currentNode.topRightNode == nil &&
		currentNode.bottomLeftNode == nil && currentNode.bottomRightNode == nil {
		// si au bout de la node (couche feuille)
		return currentNode
	} else {
		if currentNode.topLeftNode != nil { // si node existe & content à -1
			nextNode := currentNode.topLeftNode
			if targetY >= nextNode.topLeftY &&
				targetY <= (nextNode.topLeftY+nextNode.height-1) &&
				targetX >= nextNode.topLeftX &&
				targetX <= (nextNode.topLeftX+nextNode.width-1) { // si target est dans la future node
				return seeker(nextNode, targetX, targetY)
			}
		}

		if currentNode.topRightNode != nil {
			nextNode := currentNode.topRightNode
			if targetY >= nextNode.topLeftY &&
				targetY <= (nextNode.topLeftY+nextNode.height-1) &&
				targetX >= nextNode.topLeftX &&
				targetX <= (nextNode.topLeftX+nextNode.width-1) {
				return seeker(nextNode, targetX, targetY)
			}
		}

		if currentNode.bottomLeftNode != nil {
			nextNode := currentNode.bottomLeftNode
			if targetY >= nextNode.topLeftY &&
				targetY <= (nextNode.topLeftY+nextNode.height-1) &&
				targetX >= nextNode.topLeftX &&
				targetX <= (nextNode.topLeftX+nextNode.width-1) {
				return seeker(nextNode, targetX, targetY)
			}
		}

		if currentNode.bottomRightNode != nil {
			nextNode := currentNode.bottomRightNode
			if targetY >= nextNode.topLeftY &&
				targetY <= (nextNode.topLeftY+nextNode.height-1) &&
				targetX >= nextNode.topLeftX &&
				targetX <= (nextNode.topLeftX+nextNode.width-1) {
				return seeker(nextNode, targetX, targetY)
			}
		}
	}
	// else implicite -> si bloqué et ne peut aller plus loin dans l'arbre
	return currentNode
}
