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
	getNodeContent(q.root, topLeftX, topLeftY, contentHolder)
	// TODO : ajouter cases vide ?
	fmt.Println("exit ?")
}

func getNodeContent(currentNode *node, topLeftX, topLeftY int, contentHolder [][]int) {
	/*
		Algo :
			Remonte une branche de l'arbre jusqu'a ce que node.content != nil
			- si coordonnées dans contentHolder alors append
			- sinon échec
	*/
	fmt.Println("there ?")
	if currentNode.topLeftNode == nil { // si il n'y a pas de node attaché au bout
		fmt.Println("try append")
		if currentNode.topLeftY >= topLeftY &&
			(currentNode.topLeftY-topLeftY) < len(contentHolder) &&
			currentNode.topLeftX >= topLeftX &&
			(currentNode.topLeftX-topLeftX) < len(contentHolder[0]) { // si node dans coordonnée de contentHolder (supposé rectangulaire)

			fmt.Println("current contentHolder\t", contentHolder)
			// alors append au bon endroit dans contentHolder
			//contentHolder[currentNode.topLeftY-topLeftY] = append(contentHolder[currentNode.topLeftY-topLeftY], currentNode.content)
			contentHolder[currentNode.topLeftY-topLeftY][currentNode.topLeftX-topLeftX] = currentNode.content
			fmt.Println("new contentHolder\t", contentHolder)
		}
	} else { // si il y a des nodes attaché, alors, continué de remonter
		fmt.Println("remonte")
		getNodeContent(currentNode.topLeftNode, topLeftX, topLeftY, contentHolder)
		getNodeContent(currentNode.topRightNode, topLeftX, topLeftY, contentHolder)
		getNodeContent(currentNode.bottomLeftNode, topLeftX, topLeftY, contentHolder)
		getNodeContent(currentNode.bottomRightNode, topLeftX, topLeftY, contentHolder)
	}
}
