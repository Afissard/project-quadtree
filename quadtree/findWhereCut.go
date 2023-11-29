package quadtree

func (q Quadtree) renameMe(x, y int) (nodeFound *node) {
	return findWhereCut(q.root, x, y)
}

func findWhereCut(nodeQueJeRecherche *node, targetX, targetY int) *node {
	if nodeQueJeRecherche != nil {
		if nodeQueJeRecherche.topLeftX == targetX && nodeQueJeRecherche.topLeftY == targetY {
			return nodeQueJeRecherche
		} else {
			if nodeQueJeRecherche.topLeftNode != nil {
				nodeSuivante := nodeQueJeRecherche.topLeftNode
				return findWhereCut(nodeSuivante, targetX, targetY)
			}
			if nodeQueJeRecherche.topRightNode != nil {
				nodeSuivante := nodeQueJeRecherche.topRightNode
				return findWhereCut(nodeSuivante, targetX, targetY)
			}
			if nodeQueJeRecherche.bottomLeftNode != nil {
				nodeSuivante := nodeQueJeRecherche.bottomLeftNode
				return findWhereCut(nodeSuivante, targetX, targetY)
			}
			if nodeQueJeRecherche.bottomRightNode != nil {
				nodeSuivante := nodeQueJeRecherche.bottomRightNode
				return findWhereCut(nodeSuivante, targetX, targetY)
			}
		}
	}
	return nodeQueJeRecherche
}
