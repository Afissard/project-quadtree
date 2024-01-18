package quadtree

/*
Fonction Seek (fichier seek.go)
Entrées: targetX / targetY -> coordonnées de la feuille (point à atteindre)
Sorties: nodeFound *node -> pointeur de la node trouvée ou emplacement juste avant celle-ci dans le cas
où elle n'existe pas
Fonction appelant l'autre fonction du fichier: SeekIn, l'ensemble permet alors de trouver le pointeur
de la node où se trouvent ces coordonnées sinon elle renvoie le pointeur de la dernière node non vide
*/
func (q *Quadtree) Seek(targetX, targetY int) (nodeFound *node) {
	return q.root.SeekIn(targetX, targetY)
}

func (q *Quadtree) GetTileValue(targetX, targetY int) (value int) {
	return q.Seek(targetX, targetY).content
}

/*
Fonction SeekIn (fichier seek.go)

Entrées: targetX / targetY -> coordonnées de la feuille (point à atteindre)
Sorties: nodeFound *node -> pointeur de la node trouvée ou emplacement avant celle-ci dans le cas
d'un arbre optimisé et où elle n'existe pas véritablement mais est représenté par une node parente.

Fonction de recherche dans l'arbre à partir de la node donnée. Retourne la node correspondant aux
coordonnée donnée, cette node peut représenter plusieurs nodes dans un arbre optimisé.
*/
func (currentNode *node) SeekIn(targetX, targetY int) (nodeFound *node) {
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
				return nextNode.SeekIn(targetX, targetY)
			}
		}

		if currentNode.topRightNode != nil {
			nextNode := currentNode.topRightNode
			if targetY >= nextNode.topLeftY &&
				targetY <= (nextNode.topLeftY+nextNode.height-1) &&
				targetX >= nextNode.topLeftX &&
				targetX <= (nextNode.topLeftX+nextNode.width-1) {
				return nextNode.SeekIn(targetX, targetY)
			}
		}

		if currentNode.bottomLeftNode != nil {
			nextNode := currentNode.bottomLeftNode
			if targetY >= nextNode.topLeftY &&
				targetY <= (nextNode.topLeftY+nextNode.height-1) &&
				targetX >= nextNode.topLeftX &&
				targetX <= (nextNode.topLeftX+nextNode.width-1) {
				return nextNode.SeekIn(targetX, targetY)
			}
		}

		if currentNode.bottomRightNode != nil {
			nextNode := currentNode.bottomRightNode
			if targetY >= nextNode.topLeftY &&
				targetY <= (nextNode.topLeftY+nextNode.height-1) &&
				targetX >= nextNode.topLeftX &&
				targetX <= (nextNode.topLeftX+nextNode.width-1) {
				return nextNode.SeekIn(targetX, targetY)
			}
		}
	}
	// else implicite -> si bloqué et ne peut aller plus loin dans l'arbre
	return currentNode
}
