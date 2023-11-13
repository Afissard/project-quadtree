package quadtree

import "fmt"

/*
Fonction pour optimisé les quadtree :
Parcours l'arbre et retire les branche inutile
*/
func optimize(n *node) (content int) {
	if n.topLeftNode == nil {
		fmt.Printf("feuille : x%d y%d content: %d\n", n.topLeftX, n.topLeftY, n.content)
		return n.content
	} else {
		topLeft := optimize(n.topLeftNode)
		topRight := optimize(n.topRightNode)
		bottomLeft := optimize(n.bottomLeftNode)
		bottomRight := optimize(n.bottomRightNode)
		fmt.Printf("%d %d\n%d %d\n", topLeft, topRight, bottomLeft, bottomRight)
		if (topLeft == topRight) && (bottomLeft == bottomRight) && (topLeft == bottomLeft) {
			n.content = n.topLeftNode.content
			fmt.Printf("optimisation, content: %d\n", n.content)

			// suppression des nodes enfant en supprimant les ptr (attention mémoire ?)
			// bug : ne met pas à jour les pointeurs après la fonction
			n.topLeftNode = nil
			n.topRightNode = nil
			n.bottomLeftNode = nil
			n.bottomRightNode = nil
		}
		return n.content
	}
}
