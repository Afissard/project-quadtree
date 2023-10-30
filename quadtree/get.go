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
			// TODO : ajouter cases vide ?
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
		if currentNode.topLeftNode != nil {
			if targetY >= currentNode.topLeftNode.topLeftY &&
				targetY <= (currentNode.topLeftNode.topLeftY+currentNode.topLeftNode.height) &&
				targetX >= currentNode.topLeftNode.topLeftX &&
				targetX <= (currentNode.topLeftNode.topLeftX+currentNode.topLeftNode.width) {

				fmt.Println("\n---> top left")
				return getNodeContent(currentNode.topLeftNode, targetX, targetY)
			}

		} else if currentNode.topRightNode != nil {
			if targetY >= currentNode.topRightNode.topLeftY &&
				targetY <= (currentNode.topRightNode.topLeftY+currentNode.topRightNode.height) &&
				targetX >= currentNode.topRightNode.topLeftX &&
				targetX <= (currentNode.topRightNode.topLeftX+currentNode.topRightNode.width) {

				fmt.Println("\n---> top right")
				return getNodeContent(currentNode.topRightNode, targetX, targetY)
			}

		} else if currentNode.bottomLeftNode != nil {
			if targetY >= currentNode.bottomLeftNode.topLeftY &&
				targetY <= (currentNode.bottomLeftNode.topLeftY+currentNode.bottomLeftNode.height) &&
				targetX >= currentNode.bottomLeftNode.topLeftX &&
				targetX <= (currentNode.bottomLeftNode.topLeftX+currentNode.bottomLeftNode.width) {

				fmt.Println("\n---> bottom left")
				return getNodeContent(currentNode.bottomLeftNode, targetX, targetY)
			}

		} else if currentNode.bottomRightNode != nil {
			if targetY >= currentNode.bottomRightNode.topLeftY &&
				targetY <= (currentNode.bottomRightNode.topLeftY+currentNode.bottomRightNode.height) &&
				targetX >= currentNode.bottomRightNode.topLeftX &&
				targetX <= (currentNode.bottomRightNode.topLeftX+currentNode.bottomRightNode.width) {

				fmt.Println("\n---> bottom right")
				return getNodeContent(currentNode.bottomRightNode, targetX, targetY)
			}
		}
	}
	fmt.Printf("\nfound nothing for x:%d y:%d \n", targetX, targetY)
	return -1 // vide
}

/*
func getNodeContent(currentNode *node, topLeftX, topLeftY int, contentHolder [][]int) {
	// Algo :
	// Remonte une branche de l'arbre jusqu'a ce que node.content != nil
	// - si coordonnées dans contentHolder alors append
	// - sinon échec
	fmt.Println("wrong fonction")

	if currentNode.topLeftNode == nil { // si il n'y a pas de node attaché au bout
		if currentNode.topLeftY >= topLeftY &&
			(currentNode.topLeftY-topLeftY) < len(contentHolder) &&
			currentNode.topLeftX >= topLeftX &&
			(currentNode.topLeftX-topLeftX) < len(contentHolder[0]) { // si node dans coordonnée de contentHolder (supposé rectangulaire)

			// alors ajout du contenu de la node au bon endroit dans contentHolder
			contentHolder[currentNode.topLeftY-topLeftY][currentNode.topLeftX-topLeftX] = currentNode.content
			// HERE: retour en arrière vers les autres branches
		}
	} else { // si il y a des nodes attaché, alors, continué de remonter
			// FIXME: pas de retour en arrière dans les branches ...
			// Solutions :
			// - programmation "dynamique" : array "Petit Poucet" ?
			// - algo backward ? (on doit trouver tout les résultats, pas qu'un seul)
		if currentNode.topLeftNode != nil {
			getNodeContent(currentNode.topLeftNode, topLeftX, topLeftY, contentHolder)
		}
		if currentNode.topRightNode != nil {
			getNodeContent(currentNode.topRightNode, topLeftX, topLeftY, contentHolder)
		}
		if currentNode.bottomLeftNode != nil {
			getNodeContent(currentNode.bottomLeftNode, topLeftX, topLeftY, contentHolder)
		}
		if currentNode.bottomRightNode != nil {
			getNodeContent(currentNode.bottomRightNode, topLeftX, topLeftY, contentHolder)
		}
	}
}
*/
