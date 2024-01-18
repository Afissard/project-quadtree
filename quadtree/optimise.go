package quadtree

/*
Fonction optimise (fichier optimise.go)

Entrées: la node à optimiser
Sorties: la même node avec son contenu mis à jour si elle est optimisé

Fonction qui supprime les nodes enfants si les quatre enfant ont le
même contenu et sont des node feuille. La node parente est alors optimisé
en récupérant le contenu des nodes enfant avant de supprimé lesdits enfants.
*/
func (n *node) optimise() {
	if n.topLeftNode != nil && n.topRightNode != nil &&
		n.bottomLeftNode != nil && n.bottomRightNode != nil {
		// si 4 branches alors on essai d'optimiser
		if n.topLeftNode.content == n.topRightNode.content &&
			n.topLeftNode.content == n.bottomLeftNode.content &&
			n.bottomLeftNode.content == n.bottomRightNode.content &&
			n.topLeftNode.isLeave() && n.topRightNode.isLeave() &&
			n.bottomLeftNode.isLeave() && n.bottomRightNode.isLeave() {
			// mise à jour de node.content -> prend la valeur des nodes enfant
			n.content = n.topLeftNode.content
			// suppression des nodes désormais inutiles
			n.topLeftNode = nil
			n.topRightNode = nil
			n.bottomLeftNode = nil
			n.bottomRightNode = nil
		}
	}
}

/*
Fonction isLeave (fichier optimise.go)

Entrées: node testé
Sortie: booléen indiquant si la node est une feuille (true) ou non (false)

Fonction qui test si la node donnée à des enfant ou non.
*/
func (n node) isLeave() (answer bool) {
	answer = true
	if n.topLeftNode != nil || n.topRightNode != nil || n.bottomLeftNode != nil || n.bottomRightNode != nil {
		answer = false
	}
	return answer
}
