package quadtree

import "log"

/*
Fonction AddNode (fichier addNode.go)

Entrées: targetX, targetY, content -> coordonnée et contenu de la node à ajouter
Sorties: la node créer est ajouter au quadtree

Fonction qui créer une node dans le quadtree aux coordonnée voulue avec le contenu
désiré. Elle sert d'interface pour addNodeFrom, commence la recherche du point
d'ancrage en haut de l'arbre.
*/
func (q *Quadtree) AddNode(targetX, targetY, content int) {
	q.Seek(targetX, targetY).addNodeFrom(targetX, targetY, content)
}

/*
Fonction addNodeFrom (fichier addNode.go)

Entrées: targetX, targetY, content -> coordonnée et contenu de la node à ajouter
Sorties: la node créer est ajouter au quadtree

Fonction qui, à partir de la node donnée va créer si nécessaire une nouvelle couche de nodes
jusqu'à atteindre la node ciblé, une fois atteinte le contenu de la node est mis à jour.
Après cela, l'arbre est optimisé en fin de fonction (remonte dans les appels récursifs)
*/
func (previousNode *node) addNodeFrom(targetX, targetY, content int) {
	currentNode := previousNode.SeekIn(targetX, targetY) // Se déplace dans l'arbre existant
	if currentNode.width > 1 || currentNode.height > 1 {
		// Création de nodes
		// dimension des nodes
		halfWidth, halfHeight := currentNode.width, currentNode.height
		if halfWidth%2 != 0 {
			halfWidth++
		}
		if halfHeight%2 != 0 {
			halfHeight++
		}
		halfWidth = halfWidth / 2
		halfHeight = halfHeight / 2
		// création des 4 node
		currentNode.topLeftNode = createNode(currentNode.topLeftX, currentNode.topLeftY, halfWidth, halfHeight, -1)
		currentNode.topRightNode = createNode(currentNode.topLeftX+halfWidth, currentNode.topLeftY, halfWidth, halfHeight, -1)
		currentNode.bottomLeftNode = createNode(currentNode.topLeftX, currentNode.topLeftY+halfHeight, halfWidth, halfHeight, -1)
		currentNode.bottomRightNode = createNode(currentNode.topLeftX+halfWidth, currentNode.topLeftY+halfHeight, halfWidth, halfHeight, -1)
		// Appel récursif
		currentNode.addNodeFrom(targetX, targetY, content)
	} else if currentNode.topLeftX == targetX && currentNode.topLeftY == targetY {
		// Attention : réécriture possible mais nécessaire pour la génération infini !
		currentNode.content = content
	} else {
		log.Panicf("Quadtree malformé ! currentNode : %v", currentNode)
	}
	currentNode.optimise()
}

/*
Fonction createNode (fichier addNode.go)

Entrées: topLeftX, topLeftY, width, height, content -> variables de la nouvelle node
Sorties: newNode -> node créer par la fonction

Fonction utilisé pour créer une node et la rattaché à un pointeur d'une node parente
lors de l'appel de ladite fonction.
*/
func createNode(topLeftX, topLeftY, width, height, content int) (newNode *node) {
	newNode = &node{
		topLeftX: topLeftX,
		topLeftY: topLeftY,
		width:    width,
		height:   height,
		content:  content,
	}
	return newNode
}
