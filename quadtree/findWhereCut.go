package quadtree

func (q Quadtree) FindWhereMod(topX, topY int) *node {
	// Start from the root of the quadtree
	return findWhereModRecursive(q.root, topX, topY)
}

func findWhereModRecursive(currentNode *node, targetX, targetY int) *node {
	if currentNode == nil {
		// If the current node is nil, no suitable node is found
		return nil
	}

	// Check if the target coordinates are within the bounds of the current node
	if targetY >= currentNode.topLeftY &&
		targetY <= (currentNode.topLeftY+currentNode.height-1) &&
		targetX >= currentNode.topLeftX &&
		targetX <= (currentNode.topLeftX+currentNode.width-1) {

		// Check if the current node is a leaf node
		if currentNode.topLeftNode == nil &&
			currentNode.topRightNode == nil &&
			currentNode.bottomLeftNode == nil &&
			currentNode.bottomRightNode == nil {
			// If it's a leaf node, return the current node for modification
			return currentNode
		}

		// If not a leaf node, recursively search in the child nodes
		// You may need to adjust this based on your specific requirements
		if modNode := findWhereModRecursive(currentNode.topLeftNode, targetX, targetY); modNode != nil {
			return modNode
		}
		if modNode := findWhereModRecursive(currentNode.topRightNode, targetX, targetY); modNode != nil {
			return modNode
		}
		if modNode := findWhereModRecursive(currentNode.bottomLeftNode, targetX, targetY); modNode != nil {
			return modNode
		}
		if modNode := findWhereModRecursive(currentNode.bottomRightNode, targetX, targetY); modNode != nil {
			return modNode
		}
	}

	// If the target coordinates are not within the bounds of the current node, return nil
	return nil
}
