package dungeon

/*
searchRoom retourne la salle, présente dans la branche, la plus proche des coordonnées données en paramètres
*/
func (n *bspNode) searchRoom(targetX, targetY int) (roomFound *bspNode) {
	if n.nodeLeft != nil && n.nodeRight != nil {
		/*
			comparaison des deux branche -> en choisie une pour continué la recherche
			1. détermine l'axe de comparaison (=> sur quel axe a été opéré la coupe lors de la génération de la node)
			2. détermine la node la plus proche de la cible et l'utilise pour la prochaine recherche
		*/
		if distanceBetween(n.nodeLeft.topX, n.nodeRight.topX) > distanceBetween(n.nodeLeft.topY, n.nodeRight.topY) {
			if distanceBetween(n.nodeLeft.topX, targetX) < distanceBetween(n.nodeRight.topX, targetX) {
				return n.nodeLeft.searchRoom(targetX, targetY)
			} else {
				return n.nodeRight.searchRoom(targetX, targetY)
			}
		} else {
			if distanceBetween(n.nodeLeft.topY, targetY) < distanceBetween(n.nodeRight.topY, targetY) {
				return n.nodeLeft.searchRoom(targetX, targetY)
			} else {
				return n.nodeRight.searchRoom(targetX, targetY)
			}
		}
	}
	// else implicite
	return n
}

/*
distanceBetween calcule la distance entre deux entier représentant deux coordonnées sur un même axe
*/
func distanceBetween(n1, n2 int) (result uint) {
	tempResult := n1 - n2
	if tempResult < 0 {
		tempResult = -1 * tempResult
	}
	return uint(tempResult)
}
