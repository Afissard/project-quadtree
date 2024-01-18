package quadtree

/*
Fonction AppendToQuadtree (fichier appendToQuadtree)

Entrées: targetX, targetY, content -> coordonnée et contenu de la node à ajouter
Sorties: la node créer est ajouter au quadtree, sa taille peut avoir été modifié pour contenir la nouvelle node

Cette fonction à pour but de modifié un quadtree et de l'agrandir
sans avoir à le reconstruire de zero à chaque fois.
*/
func (q *Quadtree) AppendToQuadtree(targetX, targetY, content int) {
	/*
		Algo :
		1. détermine où modifier/ajouter le contenu dans l'arbre :
			- cas n°1 : modification -> modification de la node feuille (ajout de branche si nécessaire)
			- cas n°2 : ajout d'une nouvelle branche -> calcul des nouvelles dimensions et restructuration de l'arbre
	*/
	if q.checkIsIn(targetX, targetY) { // cas n°1
		q.AddNode(targetX, targetY, content)
	} else { // cas n°2
		/*
			Avec le nouvel algo de création du quadtree, le seul cas ou la node n'est pas une feuille
			est le cas où la cible n'est pas dans l'arbre. Cela nécessite donc la modification de la
			racine du quadtree.
			Algo :
			1. calcule des nouvelle dimension (q.width*2 et q.height*2)
			2. création de q.newRoot :
				- initialisation
				- ajout à l'une des branche q.currentRoot
				- ajout du nouveau contenu
		*/
		// Création de la nouvelle racine
		newRoot := createNode(q.root.topLeftX, q.root.topLeftX, q.Width*2, q.Height*2, -1)

		// choix de la branche pour la racine actuelle
		needMoveCurrentRootX := false
		if targetX < q.root.topLeftX {
			newRoot.topLeftX = q.root.topLeftX - q.Width
			needMoveCurrentRootX = true
		}
		needMoveCurrentRootY := false
		if targetY < q.root.topLeftY {
			newRoot.topLeftY = q.root.topLeftY - q.Height
			needMoveCurrentRootY = true
		}
		if needMoveCurrentRootX && needMoveCurrentRootY {
			newRoot.bottomRightNode = q.root
		} else if needMoveCurrentRootX {
			newRoot.topRightNode = q.root
		} else if needMoveCurrentRootY {
			newRoot.bottomLeftNode = q.root
		} else {
			newRoot.topLeftNode = q.root
		}

		// assignation de la nouvelle racine
		q.root = newRoot
		q.Width = q.root.width
		q.Height = q.root.height
		q.AppendToQuadtree(targetX, targetY, content) // appel récursif pour ajouter la nouvelle node ou encore modifié le quadtree
	}
}
