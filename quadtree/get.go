package quadtree

import "fmt"

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
			alors le remplir de tuiles vide ?

		TODO: quand MakeFromArray optimisera les quadtree :
		Adapté l'algo pour prendre en compte les modifications réalisé sur la
		structure des quadtree.
	*/

	for y := 0; y < len(contentHolder); y++ {
		for x := 0; x < len(contentHolder[y]); x++ {
			targetY, targetX := y+topLeftY, x+topLeftX
			fmt.Printf("\n----\nsearch for x:%d y:%d", targetX, targetY)
			contentHolder[y][x] = getNodeContent(q.root, targetX, targetY)
		}
	}
}

func getNodeContent(currentNode *node, targetX, targetY int) (content int) {
	/*
		Algo :
		Si currentNode est une feuille, retourne currentNode.content
		Sinon si currentNode peut contenir target : se rappel dans la node potentiel
		Sinon retourne nil
	*/

	fmt.Printf("\ntarget is in x:[%d <= %d <= %d], y:[%d <= %d <= %d]",
		currentNode.topLeftX, targetX, (currentNode.topLeftX + currentNode.width - 1),
		currentNode.topLeftY, targetY, (currentNode.topLeftY + currentNode.height - 1))

	if currentNode.topLeftNode == nil { // si au bout de la node (couche feuille)
		fmt.Printf("\nnode feuille, x: %d == %d = %t y: %d == %d = %t",
			currentNode.topLeftX, targetX, (currentNode.topLeftX == targetX),
			currentNode.topLeftY, targetY, (currentNode.topLeftY == targetY))

		if (currentNode.topLeftX == targetX) && (currentNode.topLeftY == targetY) {
			fmt.Printf("\nfound : %d, at x:%d y:%d \n", currentNode.content, targetX, targetY)
			return currentNode.content
		}
	} else if targetY >= currentNode.topLeftY &&
		targetY <= (currentNode.topLeftY+currentNode.height-1) &&
		targetX >= currentNode.topLeftX &&
		targetX <= (currentNode.topLeftX+currentNode.width-1) { // si target dans node

		fmt.Printf("\ntarget is contain in this node")

		content := -1 // tuile vide par default

		// détermine potentiel conteneur
		fmt.Printf("\ntl exist : %t", currentNode.topLeftNode != nil)
		if content == -1 && currentNode.topLeftNode != nil {
			fmt.Printf("\ntry tl")
			nextNode := currentNode.topLeftNode
			if targetY >= nextNode.topLeftY &&
				targetY <= (nextNode.topLeftY+nextNode.height-1) &&
				targetX >= nextNode.topLeftX &&
				targetX <= (nextNode.topLeftX+nextNode.width-1) {

				fmt.Printf("\n---> top left")

				content = getNodeContent(nextNode, targetX, targetY)
			}
		}

		fmt.Printf("\ntr exist : %t", currentNode.topRightNode != nil)
		if content == -1 && currentNode.topRightNode != nil {
			fmt.Printf("\ntry tr")
			nextNode := currentNode.topRightNode
			if targetY >= nextNode.topLeftY &&
				targetY <= (nextNode.topLeftY+nextNode.height-1) &&
				targetX >= nextNode.topLeftX &&
				targetX < (nextNode.topLeftX+nextNode.width-1) {

				fmt.Printf("\n---> top right")
				content = getNodeContent(nextNode, targetX, targetY)
			}
		}

		fmt.Printf("\nbl exist : %t", currentNode.bottomLeftNode != nil)
		if content == -1 && currentNode.bottomLeftNode != nil {
			fmt.Printf("\ntry br")
			nextNode := currentNode.bottomLeftNode
			if targetY >= nextNode.topLeftY &&
				targetY <= (nextNode.topLeftY+nextNode.height-1) &&
				targetX >= nextNode.topLeftX &&
				targetX <= (nextNode.topLeftX+nextNode.width-1) {

				fmt.Printf("\n---> bottom left")
				content = getNodeContent(nextNode, targetX, targetY)
			}
		}

		fmt.Printf("\nbr exist : %t", currentNode.bottomRightNode != nil)
		if content == -1 && currentNode.bottomRightNode != nil {
			fmt.Printf("\ntry bl")
			nextNode := currentNode.bottomRightNode
			if targetY >= nextNode.topLeftY &&
				targetY <= (nextNode.topLeftY+nextNode.height-1) &&
				targetX >= nextNode.topLeftX &&
				targetX <= (nextNode.topLeftX+nextNode.width-1) {

				fmt.Printf("\n---> bottom right")
				content = getNodeContent(nextNode, targetX, targetY)
			}
		}
		fmt.Printf("\nfound : %d, at x:%d y:%d \n", content, targetX, targetY)
		return content
	}
	fmt.Printf("\nWhy am'I here ? found nothing for x:%d y:%d \n", targetX, targetY)
	return -1 // tuile vide
}
