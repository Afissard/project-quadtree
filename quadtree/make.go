package quadtree

import "fmt"

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.
func MakeFromArray(floorContent [][]int) (q Quadtree) {
	/*
		En entré : la map (liste 2D)
		En sortie : un quadtree qui stocke la map

		Algo :
		1. créer la couche "feuille" de l'arbre quadtree : chaque node récupère le contenu de floorContent
		2. Boucle récursive :
			- si couche 0 atteint (1 seule node à attacher) alors attacher la node restante à la racine.
			- sinon créer des nodes pour rattacher les node précédemment créés (ou aux données de la map à itération n°1).

		TODO: optimiser les quadtree (fonction à part ?):
		Si une node représente 4 fois le même contenu, alors "supprimer" les node d'en dessous,
		node.content prend la valeur du contenu des nodes supprimées, le contenu d'en dessous
		est déterminé grace avec node.width et node.height.
	*/

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
	return q
}

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

func newCNL(nodesList [][]node) (q Quadtree) {
	fmt.Println("\n---New Call---")
	// test racine de quadtree : si une seule node existe
	if len(nodesList) == 1 && len(nodesList[0]) == 1 {
		fmt.Printf("\nroot\n")
		q.width = nodesList[0][0].width
		q.height = nodesList[0][0].height
		q.root = &nodesList[0][0]
		return q
	}
	// else implicite : calcul d'une nouvelle couche
	newNodesList := [][]node{}
	for y := 0; y < len(nodesList); y++ {
		nodesLine := []node{}
		for x := 0; x < len(nodesList[y]); x++ {
			currentNode := node{}         // déclaration de currentNode ici car sinon hors porté dans if/else
			if (x%2 == 0) && (y%2 == 0) { // créer nouvelle node ?
				currentNode = node{
					topLeftX: x,
					topLeftY: y,
					width:    nodesList[y][x].width * 2, // TODO: attention aux potentiels tuiles vide avec optimisation ?
					height:   nodesList[y][x].height * 2,
				}
			} else { // sinon récupérer celle déjà créer
				currentNode = nodesList[y-(y%2)][x-(x%2)]
			}

			// fmt.Printf("\nnodeContent :%d, nodeTopCoord: %d:%d, loopCoord: %d:%d, nodeListSize: %d:%d \n", nodesList[y][x].content,
			// nodesList[y][x].topLeftX, nodesList[y][x].topLeftY, x, y, len(nodesList[y]), len(nodesList))

			// lie la node
			if (x%2 == 0) && (y%2 == 0) {
				currentNode.topLeftNode = &nodesList[y][x]
				//fmt.Printf("\nx:%d y:%d --> tl : %p", x, y, currentNode.topLeftNode)
			} else if (x%2 == 1) && (y%2 == 0) {
				currentNode.topRightNode = &nodesList[y][x]
				//fmt.Printf("\nx:%d y:%d --> tr : %p", x, y, currentNode.topRightNode)
			} else if (x%2 == 0) && (y%2 == 1) {
				currentNode.bottomLeftNode = &nodesList[y][x]
				//fmt.Printf("\nx:%d y:%d --> bl : %p", x, y, currentNode.bottomLeftNode)
			} else if (x%2 == 1) && (y%2 == 1) {
				currentNode.bottomRightNode = &nodesList[y][x]
				//fmt.Printf("\nx:%d y:%d --> br : %p", x, y, currentNode.bottomRightNode)
			}

			// ligne de nodes
			if x%2 == 0 { // append si pair
				//fmt.Println("\nNODE append :", currentNode)
				nodesLine = append(nodesLine, currentNode)
			} else { // maj de la liste
				nodesLine[x-1] = currentNode // produce error : out of range
			}
		}
		fmt.Println("\nLine")
		for x := 0; x < len(nodesLine); x++ {
			fmt.Println(nodesLine[x])
		}

		// tableau 2D de nodes
		if y%2 == 0 { // append si pair, mais contenu pas mis a jour ... -> utilisé ptr ?
			//fmt.Println("\n line : ", nodesLine)
			newNodesList = append(newNodesList, nodesLine) //BUG: lost pointer when append it to list
		} else { // maj de la liste
			newNodesList[y-1] = nodesLine
		}
	}
	fmt.Println("\nListe:")
	for y := 0; y < len(newNodesList); y++ {
		fmt.Println("line :", y)
		for x := 0; x < len(newNodesList[y]); x++ {
			fmt.Println(newNodesList[y][x])
		}
	}

	q = newCNL(newNodesList)
	return q
}
