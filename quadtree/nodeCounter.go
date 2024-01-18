package quadtree

/*
Fonction NodeCounter (fichier nodeCounter.go)

Sortie : retourne le nombre de node dans le quadtree
*/
func (q *Quadtree) NodeCounter() (nbrNode int) {
	return q.root.nodeCount()
}

/*
Fonction nodeCount (fichier nodeCounter.go)

Sortie : retourne le nombre de node enfant de cette node, elle même inclue dans le compte.
*/
func (n *node) nodeCount() (nbrNode int) {
	nbrNode = 1
	if n.topLeftNode != nil {
		nbrNode += n.topLeftNode.nodeCount()
	}
	if n.topRightNode != nil {
		nbrNode += n.topRightNode.nodeCount()
	}
	if n.bottomLeftNode != nil {
		nbrNode += n.bottomLeftNode.nodeCount()
	}
	if n.bottomRightNode != nil {
		nbrNode += n.bottomRightNode.nodeCount()
	}
	return nbrNode
}
