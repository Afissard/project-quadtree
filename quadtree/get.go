package quadtree

// GetContent remplit le tableau contentHolder (qui représente
// un terrain dont la case le plus en haut à gauche a pour coordonnées
// (topLeftX, topLeftY)) à partir du quadtree q.
func (q Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int) {
	/*
		En entré : contentHolder, les coordonnées de départ (pour le remplir)
		En sortie : contentHolder remplis au maximum avec le contenu de q (Quadtree)
		en commençant au coordonnées données (contentHolder est une copie, donc pas
		besoin de le retourné, il est directement modifié à l’exécution)

		Algo :
		1. Remonte une branche de l'arbre jusqu'a ce que node.content != nil
			- si coordonnées dans contentHolder alors append
			- sinon échec
		2. Si toutes les branches ont été explorées et contentHolder pas remplit,
			alors le remplir de tuiles vide (tuile par default)
	*/
	for y := 0; y < len(contentHolder); y++ {
		for x := 0; x < len(contentHolder[y]); x++ {
			targetY, targetX := y+topLeftY, x+topLeftX
			//fmt.Printf("\n----\nsearch for x:%d y:%d", targetX, targetY)
			//contentHolder[y][x] = getNodeContent(q.root, targetX, targetY) // remove for seeker
			contentHolder[y][x] = q.seek(targetX, targetY).content
		}
	}
}

// func getNodeContent(currentNode *node, targetX, targetY int) (content int) {
// 	/*
// 		Algo :
// 		Si currentNode est une feuille, retourne currentNode.content
// 		Sinon si currentNode peut contenir target : se rappel dans la node potentiel
// 		Sinon retourne nil
// 	*/
// 	// fmt.Println("\ncurrent node is :", currentNode)
// 	// fmt.Printf("\ntarget is in x:[%d <= %d <= %d], y:[%d <= %d <= %d]",
// 	// currentNode.topLeftX, targetX, (currentNode.topLeftX + currentNode.width - 1),
// 	// currentNode.topLeftY, targetY, (currentNode.topLeftY + currentNode.height - 1))

// 	content = -1 // tuile vide par default

// 	if currentNode.topLeftNode == nil && currentNode.topRightNode == nil &&
// 		currentNode.bottomLeftNode == nil && currentNode.bottomRightNode == nil { // si au bout de la node (couche feuille)
// 		return currentNode.content
// 	} else if targetY >= currentNode.topLeftY &&
// 		targetY <= (currentNode.topLeftY+currentNode.height-1) &&
// 		targetX >= currentNode.topLeftX &&
// 		targetX <= (currentNode.topLeftX+currentNode.width-1) { // si target dans node (redondance)

// 		// détermine potentiel conteneur
// 		if content == -1 && currentNode.topLeftNode != nil { // si node existe & content à -1
// 			nextNode := currentNode.topLeftNode
// 			if targetY >= nextNode.topLeftY &&
// 				targetY <= (nextNode.topLeftY+nextNode.height-1) &&
// 				targetX >= nextNode.topLeftX &&
// 				targetX <= (nextNode.topLeftX+nextNode.width-1) { // si target est dans la future node

// 				//fmt.Printf("\n---> top left")
// 				content = getNodeContent(nextNode, targetX, targetY)
// 				return content
// 			}
// 		}

// 		if content == -1 && currentNode.topRightNode != nil {
// 			nextNode := currentNode.topRightNode
// 			if targetY >= nextNode.topLeftY &&
// 				targetY <= (nextNode.topLeftY+nextNode.height-1) &&
// 				targetX >= nextNode.topLeftX &&
// 				targetX <= (nextNode.topLeftX+nextNode.width-1) {

// 				//fmt.Printf("\n---> top right")
// 				content = getNodeContent(nextNode, targetX, targetY)
// 				return content
// 			}
// 		}

// 		if content == -1 && currentNode.bottomLeftNode != nil {
// 			nextNode := currentNode.bottomLeftNode
// 			if targetY >= nextNode.topLeftY &&
// 				targetY <= (nextNode.topLeftY+nextNode.height-1) &&
// 				targetX >= nextNode.topLeftX &&
// 				targetX <= (nextNode.topLeftX+nextNode.width-1) {

// 				//fmt.Printf("\n---> bottom left")
// 				content = getNodeContent(nextNode, targetX, targetY)
// 				return content
// 			}
// 		}

// 		if content == -1 && currentNode.bottomRightNode != nil {
// 			nextNode := currentNode.bottomRightNode
// 			if targetY >= nextNode.topLeftY &&
// 				targetY <= (nextNode.topLeftY+nextNode.height-1) &&
// 				targetX >= nextNode.topLeftX &&
// 				targetX <= (nextNode.topLeftX+nextNode.width-1) {

// 				//fmt.Printf("\n---> bottom right")
// 				content = getNodeContent(nextNode, targetX, targetY)
// 				return content
// 			}
// 		}
// 	}
// 	//fmt.Printf("\nfound : %d, at x:%d y:%d \n", content, targetX, targetY)
// 	return content
// }
