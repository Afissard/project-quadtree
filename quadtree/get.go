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
			fmt.Printf("\nsearch for x:%d y:%d", targetX, targetY)
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
	if currentNode.topLeftNode == nil { // si au bout de la node (couche feuille)
		if currentNode.topLeftX == targetX && currentNode.topLeftY == targetY {
			fmt.Printf("\nfound : %d, at x:%d y:%d", currentNode.content, targetX, targetY)
			return currentNode.content
		}
	} else if targetY >= currentNode.topLeftY &&
		targetY <= (currentNode.topLeftY+currentNode.height) &&
		targetX >= currentNode.topLeftX &&
		targetX <= (currentNode.topLeftX+currentNode.width) { // si target dans node
		fmt.Printf("\ntarget is in x:[%d <= %d <= %d], y:[%d <= %d <= %d]",
			currentNode.topLeftX, targetX, (currentNode.topLeftX + currentNode.width),
			currentNode.topLeftY, targetY, (currentNode.topLeftY + currentNode.height))

		// détermine potentiel conteneur
		if currentNode.topLeftNode != nil { // FIXME: erreur ici : forcement vrai puisque pas racine
			nextNode := currentNode.topLeftNode
			if targetY >= nextNode.topLeftY &&
				targetY <= (nextNode.topLeftY+nextNode.height) &&
				targetX >= nextNode.topLeftX &&
				targetX <= (nextNode.topLeftX+nextNode.width) {

				fmt.Println("\n---> top left")
				return getNodeContent(nextNode, targetX, targetY)
			}

		} else if currentNode.topRightNode != nil {
			nextNode := currentNode.topRightNode
			if targetY >= nextNode.topLeftY &&
				targetY <= (nextNode.topLeftY+nextNode.height) &&
				targetX >= nextNode.topLeftX &&
				targetX <= (nextNode.topLeftX+nextNode.width) {

				fmt.Println("\n---> top right")
				return getNodeContent(nextNode, targetX, targetY)
			}

		} else if currentNode.bottomLeftNode != nil {
			nextNode := currentNode.bottomLeftNode
			if targetY >= nextNode.topLeftY &&
				targetY <= (nextNode.topLeftY+nextNode.height) &&
				targetX >= nextNode.topLeftX &&
				targetX <= (nextNode.topLeftX+nextNode.width) {

				fmt.Println("\n---> bottom left")
				return getNodeContent(nextNode, targetX, targetY)
			}

		} else if currentNode.bottomRightNode != nil {
			nextNode := currentNode.bottomRightNode
			if targetY >= nextNode.topLeftY &&
				targetY <= (nextNode.topLeftY+nextNode.height) &&
				targetX >= nextNode.topLeftX &&
				targetX <= (nextNode.topLeftX+nextNode.width) {

				fmt.Println("\n---> bottom right")
				return getNodeContent(nextNode, targetX, targetY)
			}
		}
	}
	fmt.Printf("\nfound nothing for x:%d y:%d \n", targetX, targetY)
	return -1 // vide
}
