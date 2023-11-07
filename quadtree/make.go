package quadtree

import "fmt"

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.
func MakeFromArray(floorContent [][]int) (q Quadtree) {
	/*
		En entré : la map (liste 2D)
		En sortie : un quadtree qui stocke la map

		Algo 1:
		1. créer la couche "feuille" de l'arbre quadtree : chaque node récupère le contenu de floorContent
		2. Boucle récursive :
			- si couche 0 atteint (1 seule node à attacher) alors attacher la node restante à la racine.
			- sinon créer des nodes pour rattacher les node précédemment créés (ou aux données de la map à itération n°1)

		(Nouvel) Algo 2 :
		créer la racine, taille de floorContent
		Boucle parcours floorContent :
			- Pour chaque élément appel addContent
		addContent :
			- regarde les coordonnée et cherche la branche qui correspond :
				- si au bout d'une branche et taille<1 -> créer une branche -> récursion dans la branche
				- si au bout d'une branche et taille == 1 -> assigne la valeur -> fin
				- sinon erreur

		TODO: optimiser les quadtree (fonction à part ?):
		Si une node représente 4 fois le même contenu, alors "supprimer" les node d'en dessous,
		node.content prend la valeur du contenu des nodes supprimées, le contenu d'en dessous
		est déterminé grace avec node.width et node.height.
	*/
	/*
		// Algo 1
			if len(floorContent) <= 0 || len(floorContent[0]) <= 0 {
				panic("floorContent vide !")
			} else {
				// création des node pour la couche "feuille"
				nodesList := [][]node{}
				for y := 0; y < len(floorContent); y++ {
					nodesLine := []node{}
					for x := 0; x < len(floorContent[y]); x++ {
						currentNode := node{
							topLeftX: x,
							topLeftY: y,
							width:    1,
							height:   1,
							content:  floorContent[y][x],
						}
						// ligne de nodes
						nodesLine = append(nodesLine, currentNode)
					}
					// tableau 2D de nodes
					nodesList = append(nodesList, nodesLine)
				}
				//q = createNodesLayer(nodesList)
				q = createNodesLayer(nodesList)
			}
	*/

	// Algo 2
	totalWidth := len(floorContent[0])
	totalHeight := len(floorContent)
	if totalWidth%2 != 0 { // on s'assure que la taille est pair
		totalWidth++
	} else if totalHeight%2 != 0 {
		totalHeight++
	}
	rootNode := node{
		topLeftX: 0,
		topLeftY: 0,
		width:    totalWidth,
		height:   totalHeight,
		content:  -1, // contenu à -1 par défaut
	}

	q = Quadtree{
		width:  totalWidth,
		height: totalHeight,
		root:   &rootNode,
	}
	for y := 0; y < len(floorContent); y++ {
		for x := 0; x < len(floorContent[y]); x++ {
			addContent(&rootNode, floorContent[y][x], x, y)
		}
	}

	return q
}

/*
(Nouvel) Algo addContent :
regarde les coordonnée et cherche la branche qui correspond :
- si au bout d'une branche et taille<1 -> cherche/créé une branche -> récursion dans la branche
- si au bout d'une branche et taille == 1 -> assigne la valeur -> fin (retourne pointeur de la node ?)
- sinon erreur
*/
func addContent(currentNode *node, content, targetX, targetY int) {
	if currentNode.width == 1 && currentNode.height == 1 { // assignation de content
		currentNode.content = content
		fmt.Println("add", currentNode)
	} else if currentNode.width > 1 { // recherche/création d'une branche
		// recherche topLeft
		if targetX >= currentNode.topLeftX && targetX <= currentNode.topLeftX+currentNode.width/2 &&
			targetY >= currentNode.topLeftY && targetY <= currentNode.topLeftY+currentNode.height/2 {
			fmt.Println("top left")
			if currentNode.topLeftNode == nil { // si la node n'existe pas : création d'une node
				fmt.Println("new")
				newNode := node{
					topLeftX: currentNode.topLeftX,
					topLeftY: currentNode.topLeftY,
					width:    currentNode.width / 2,
					height:   currentNode.height / 2,
					content:  -1, // contenu à -1 par défaut
				}
				currentNode.topLeftNode = &newNode
			}
			addContent(currentNode.topLeftNode, content, targetX, targetY)

		} else if targetX >= currentNode.topLeftX+currentNode.width/2 && targetX <= currentNode.topLeftX+currentNode.width &&
			targetY >= currentNode.topLeftY && targetY <= currentNode.topLeftY+currentNode.height/2 {
			fmt.Println("top right")
			if currentNode.topRightNode == nil { // si la node n'existe pas : création d'une node
				fmt.Println("new")
				newNode := node{
					topLeftX: currentNode.topLeftX + currentNode.width/2,
					topLeftY: currentNode.topLeftY,
					width:    currentNode.width / 2,
					height:   currentNode.height / 2,
					content:  -1,
				}
				currentNode.topRightNode = &newNode
			}
			addContent(currentNode.topLeftNode, content, targetX, targetY)

		} else if targetX >= currentNode.topLeftX && targetX <= currentNode.topLeftX+currentNode.width/2 &&
			targetY >= currentNode.topLeftY+currentNode.height/2 && targetY <= currentNode.topLeftY+currentNode.height {
			fmt.Println("bottom left")
			if currentNode.bottomRightNode == nil { // si la node n'existe pas : création d'une node
				fmt.Println("new")
				newNode := node{
					topLeftX: currentNode.topLeftX,
					topLeftY: currentNode.topLeftY + currentNode.height/2,
					width:    currentNode.width / 2,
					height:   currentNode.height / 2,
					content:  -1,
				}
				currentNode.bottomRightNode = &newNode
			}
			addContent(currentNode.topLeftNode, content, targetX, targetY)

		} else if targetX >= currentNode.topLeftX+currentNode.width/2 && targetX <= currentNode.topLeftX+currentNode.width &&
			targetY >= currentNode.topLeftY+currentNode.height/2 && targetY <= currentNode.topLeftY+currentNode.height {
			fmt.Println("bottom right")
			if currentNode.bottomRightNode == nil { // si la node n'existe pas : création d'une node
				fmt.Println("new")
				newNode := node{
					topLeftX: currentNode.topLeftX + currentNode.width/2,
					topLeftY: currentNode.topLeftY + currentNode.height/2,
					width:    currentNode.width / 2,
					height:   currentNode.height / 2,
					content:  -1,
				}
				currentNode.bottomRightNode = &newNode
			}
			addContent(currentNode.topLeftNode, content, targetX, targetY)
		}
	} else {
		panic("quadtree malformé !")
	}
	fmt.Println("exit")
}

/*
Algo 1, défectueux pour map supérieur à 6*6:
1. créer la couche "feuille" de l'arbre quadtree : chaque node récupère le contenu de floorContent
2. Boucle récursive :
	- si couche 0 atteint (1 seule node à attacher) alors attacher la node restante à la racine.
	- sinon créer des nodes pour rattacher les node précédemment créés (ou aux données de la map à itération n°1)
*/
func createNodesLayer(nodesList [][]node) (q Quadtree) {
	/*
		Sert à makeFromArray, associe les node entre elles dans un arbre récursif

		BUG: ne marche pas avec grande map...
	*/
	fmt.Println("\n---New Call---")
	// test racine de quadtree : si une seule node existe
	if len(nodesList) == 1 && len(nodesList[0]) == 1 {
		q.width = nodesList[0][0].width
		q.height = nodesList[0][0].height
		q.root = &nodesList[0][0]
		return q
	}
	// else implicite : calcul d'une nouvelle couche
	newNodesList := [][]node{}
	for y := 0; y < len(nodesList); y += 2 { // +2 au lieu de +1 car une node couvre (2*2=4) nodes
		nodesLine := []node{}
		for x := 0; x < len(nodesList[y]); x += 2 {
			currentNode := node{
				topLeftX: x,
				topLeftY: y,
				width:    nodesList[y][x].width * 2, // TODO: attention aux potentiels tuiles vide ?
				height:   nodesList[y][x].height * 2,
			}

			// lie les nodes, en faisant attention aux potentiels nodes inexistante
			//FIXME: ne vas pas au delà de x6 y6

			currentNode.topLeftNode = &nodesList[y][x]
			if x+1 < len(nodesList[y]) {
				currentNode.topRightNode = &nodesList[y][x+1]
			}
			if y+1 < len(nodesList) {
				currentNode.bottomLeftNode = &nodesList[y+1][x]
				if x+1 < len(nodesList[y]) {
					currentNode.bottomRightNode = &nodesList[y+1][x+1]
				}
			}
			// ligne de nodes
			nodesLine = append(nodesLine, currentNode)
		}
		// tableau 2D de nodes
		newNodesList = append(newNodesList, nodesLine)
	}
	fmt.Println("\nListe:")
	for y := 0; y < len(newNodesList); y++ {
		fmt.Println("line :", y)
		for x := 0; x < len(newNodesList[y]); x++ {
			fmt.Println(newNodesList[y][x])
		}
	}

	q = createNodesLayer(newNodesList)
	return q
}
