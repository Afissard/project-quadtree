package quadtree

/*
Fonction pour optimisé les quadtree :
Parcours l'arbre et retire les branche inutile
*/
func optimize(n *node) (content int) {
	if n.topLeftNode == nil {
		//fmt.Printf("feuille : x%d y%d content: %d\n", n.topLeftX, n.topLeftY, n.content)
		return n.content
	} else {
		//BUG: semble trop supprimé des branche
		topLeft := optimize(n.topLeftNode)
		topRight := optimize(n.topRightNode)
		bottomLeft := optimize(n.bottomLeftNode)
		bottomRight := optimize(n.bottomRightNode)
		//fmt.Printf("\n%d %d\n%d %d\n", topLeft, topRight, bottomLeft, bottomRight)
		if topLeft == topRight && bottomLeft == bottomRight && topLeft == bottomLeft && n.topLeftNode.topLeftNode == nil {
			n.content = n.topLeftNode.content
			//fmt.Printf("optimisation x%d y%d, w%d h%d, content: %d\n", n.topLeftX, n.topLeftY, n.width, n.height, n.content)

			// suppression des nodes désormais inutiles
			n.topLeftNode = nil
			n.topRightNode = nil
			n.bottomLeftNode = nil
			n.bottomRightNode = nil
		}
		return n.content
	}
}
